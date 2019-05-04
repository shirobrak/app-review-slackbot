package usecases

import (
	"time"

	"github.com/shirobrak/app-review-slackbot/entities"
)

// ReviewGetterIF is an interface to get App's Review.
type ReviewGetterIF interface {
	Get(fromDateStr string) ([]entities.Review, error)
}

// ReviewsSenderIF is an interface to send App's Review.
type ReviewsSenderIF interface {
	Send(reviews []entities.Review) error
}

// BatchManagerIF is an interface to get configuration of batch.
type BatchManagerIF interface {
	GetLastUpdated() (string, error)
	SetLastUpdated(lastUpdated string) error
}

// SendReviewUseCase is an usecase that send a review to slack channel.
type SendReviewUseCase struct {
	reviewGetter ReviewGetterIF
	reviewSender ReviewsSenderIF
	batchManger  BatchManagerIF
}

// NewSendReviewUseCase is a method to create an instance of SendReviewUseCase.
func NewSendReviewUseCase(reviewGetter ReviewGetterIF, reviewSender ReviewsSenderIF, batchManager BatchManagerIF) *SendReviewUseCase {
	return &SendReviewUseCase{
		reviewGetter: reviewGetter,
		reviewSender: reviewSender,
		batchManger:  batchManager,
	}
}

// Run is a method to execute the SendReviewUseCase.
func (u *SendReviewUseCase) Run() (int, error) {

	// 最新更新日を取得する
	lastUpdated, err := u.batchManger.GetLastUpdated()
	if err != nil {
		return 0, err
	}
	// ログファイルが空の場合はデフォルトの日時を設定
	if lastUpdated == "" {
		lastUpdated = time.Date(2016, 1, 2, 34, 45, 56, 0, time.Local).Format(time.RFC3339)
	}

	// レビューを集める
	reviews, err := u.reviewGetter.Get(lastUpdated)
	if err != nil {
		return 0, err
	}

	// レビューがある場合, レビューを投稿する
	reviewNum := len(reviews)
	if reviewNum > 0 {
		err = u.reviewSender.Send(reviews)
		if err != nil {
			return 0, err
		}
		// 投稿に成功した場合は最新更新日を更新する
		now := time.Now().Format(time.RFC3339)
		err = u.batchManger.SetLastUpdated(now)
		if err != nil {
			return 0, err
		}
	}
	return reviewNum, nil
}
