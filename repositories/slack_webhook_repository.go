package repositories

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/shirobrak/app-review-slackbot/entities"
)

// Message is data struct for creating the slack message.
type Message struct {
	Name        string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment is data struct for creating the slack message.
type Attachment struct {
	AuthorName string   `json:"author_name"`
	Text       string   `json:"text"`
	Fields     []Field  `json:"fields"`
	Color      string   `json:"color"`
	MrkdwnIn   []string `json:"mrkdwn_in"`
}

// Field is data struct for creating the slack message.
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// SlackWebhookRepository is a repository for handling the slack webhook integration.
type SlackWebhookRepository struct {
	client *http.Client
}

// NewSlackWebhookRepository is a method to create an instance of SlackWebhookRepository.
func NewSlackWebhookRepository() *SlackWebhookRepository {
	client := &http.Client{}
	return &SlackWebhookRepository{client: client}
}

// Send is a method to send reviews to the specified slack channel.
func (r *SlackWebhookRepository) Send(reviews []entities.Review) error {
	var webhookURL = os.Getenv("SLACK_WEBHOOK_URL")

	msg := convReviewToMessage(reviews)

	// メッセージをJSON形式に変換
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// リクエストの作成
	req, err := http.NewRequest(
		"POST",
		webhookURL,
		bytes.NewBuffer(jsonBytes),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// SlackにPOST
	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func convReviewToMessage(reviews []entities.Review) Message {
	msg := Message{}
	msg.Name = os.Getenv("SLACK_BOT_NAME")
	msg.IconEmoji = os.Getenv("SLACK_BOT_ICON")
	msg.Channel = os.Getenv("SLACK_TARGET_CHANNEL")
	msg.Text = os.Getenv("SLACK_MSG_PREFIX_TEXT")
	for _, review := range reviews {
		attachment := Attachment{}
		attachment.MrkdwnIn = append(attachment.MrkdwnIn, "text")
		// 投稿者
		authorField := Field{Title: "", Value: "*投稿者：" + review.Author + "*"}
		attachment.Fields = append(attachment.Fields, authorField)
		// 評価
		ratingText, _ := makeRatingText(review.Rating)
		ratingFiled := Field{Title: "", Value: ratingText}
		attachment.Fields = append(attachment.Fields, ratingFiled)
		// タイトル, コメント
		commentField := Field{Title: "タイトル：" + review.Title, Value: "```" + review.Comment + "```"}
		attachment.Fields = append(attachment.Fields, commentField)
		// OS情報
		osField := Field{Title: "", Value: "*OS：" + review.Os + "*"}
		attachment.Fields = append(attachment.Fields, osField)
		if review.Os == "iOS" {
			attachment.Color = "#ff0000"
		} else if review.Os == "Android" {
			attachment.Color = "#0000ff"
		}
		// 日時
		updatedField := Field{Title: "", Value: "*投稿時刻：" + review.Updated + "*"}
		attachment.Fields = append(attachment.Fields, updatedField)
		// アタッチメントの追加
		msg.Attachments = append(msg.Attachments, attachment)
	}
	return msg
}

func makeRatingText(strRate string) (string, error) {
	iRate, err := strconv.Atoi(strRate)
	if err != nil {
		return "", err
	}
	var ratingText = "*評価：*"
	for i := 0; i < iRate; i++ {
		ratingText += ":star:"
		i++
	}
	return ratingText, nil
}
