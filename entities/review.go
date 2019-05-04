package entities

// Review is an entity.
type Review struct {
	Os      string
	Author  string
	Title   string
	Comment string
	Rating  string
	Updated string
}

// NewReview is a function.
func NewReview(os string, author string, title string, comment string, rating string, updated string) Review {
	return Review{
		Os:      os,
		Author:  author,
		Title:   title,
		Comment: comment,
		Rating:  rating,
		Updated: updated,
	}
}
