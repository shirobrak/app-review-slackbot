package adapters

import (
	"time"

	"github.com/shirobrak/app-review-slackbot/entities"
)

// ReviewRepositoryIF is an interface to access to the review repository.
type ReviewRepositoryIF interface {
	Get() ([]map[string]string, error)
}

// ReviewGetter is an adapter to get reviews from repositories.
type ReviewGetter struct {
	iosReviewRepository ReviewRepositoryIF
}

// Get is a method to get reviews from external repository.
func (rg *ReviewGetter) Get(fromDateStr string) ([]entities.Review, error) {

	// レビューリストの初期化
	var reviews []entities.Review

	// リポジトリからiOSアプリのレビューを取得
	iosReviews, err := rg.iosReviewRepository.Get()
	if err != nil {
		return nil, err
	}

	// 対象レビューの抽出（iOS）
	fromDate, err := time.Parse(time.RFC3339, fromDateStr)
	if err != nil {
		return nil, err
	}

	for _, iosReviewData := range iosReviews {
		updatedDate, err := time.Parse(time.RFC3339, iosReviewData["updated"])
		if err != nil {
			return nil, err
		}
		// レビュー日が指定日より後のレビューのみ取得
		if updatedDate.After(fromDate) {
			if iosReviewData["updated"] > fromDateStr {
				iosReview := makeReview("iOS", iosReviewData)
				reviews = append(reviews, iosReview)
			}
		}
	}

	return reviews, nil
}

// NewReviewGetter is a method to create an instance of ReviewGetter.
func NewReviewGetter(iosReviewRepository ReviewRepositoryIF) *ReviewGetter {
	return &ReviewGetter{iosReviewRepository: iosReviewRepository}
}

func makeReview(os string, reviewData map[string]string) entities.Review {
	review := entities.NewReview(
		os,
		reviewData["author"],
		reviewData["title"],
		reviewData["comment"],
		reviewData["rating"],
		reviewData["updated"],
	)
	return review
}
