package adapters

import "github.com/shirobrak/app-review-slackbot/entities"

// ReviewSenderIF is an interface for accessing a repository for sending reviews.
type ReviewSenderIF interface {
	Send(review []entities.Review) error
}

// ReviewSender is an adapter to send reviews.
type ReviewSender struct {
	senderRepository ReviewSenderIF
}

// Send is a method to send to reviews.
func (rs *ReviewSender) Send(reviews []entities.Review) error {
	err := rs.senderRepository.Send(reviews)
	if err != nil {
		return err
	}
	return nil
}

// NewReviewSender is a method to create an instance of ReviewSender.
func NewReviewSender(senderRepository ReviewSenderIF) *ReviewSender {
	return &ReviewSender{senderRepository: senderRepository}
}
