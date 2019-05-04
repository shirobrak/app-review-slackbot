package repositories

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
)

// feed is data struct for dealing with iTunes Store Rss.
type feed struct {
	Updated string  `xml:"updated"`
	Entry   []entry `xml:"entry"`
}

// entry is data struct for dealing with iTunes Store Rss.
type entry struct {
	Author  author    `xml:"author"`
	ID      string    `xml:"id"`
	Title   string    `xml:"title"`
	Content []content `xml:"content"`
	Rating  rating    `xml:"rating"`
	Updated string    `xml:"updated"`
}

// author is data struct for dealing with iTunes Store Rss.
type author struct {
	Name string `xml:"name"`
}

// content is data struct for dealing with iTunes Store Rss.
type content struct {
	TypeAttr string `xml:"type,attr"`
	Val      string `xml:",chardata"`
}

// rating is data struct for dealing with iTunes Store Rss.
type rating struct {
	Score string `xml:",chardata"`
}

// AppStoreRssRepository is a repository.
type AppStoreRssRepository struct {
	client *http.Client
}

// NewAppStoreRssRepository is a method to create an instance of AppStoreRssRepository.
func NewAppStoreRssRepository() *AppStoreRssRepository {
	client := &http.Client{}
	return &AppStoreRssRepository{client: client}
}

// Get is a method to get review data from iTunes Store Rss.
func (r *AppStoreRssRepository) Get() ([]map[string]string, error) {
	var reviews []map[string]string
	baseURL := "https://itunes.apple.com/jp/rss/customerreviews/id="
	targetURL := baseURL + os.Getenv("TARGET_IOS_APP_ID") + "/sortBy=mostRecent/xml"
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		feed := feed{}
		err = xml.Unmarshal(body, &feed)
		if err != nil {
			return nil, err
		}

		// レスポンスの整形
		for _, entry := range feed.Entry {
			review := map[string]string{}
			review["id"] = entry.ID
			review["author"] = entry.Author.Name
			review["title"] = entry.Title
			review["rating"] = entry.Rating.Score
			review["comment"] = parseComment(entry.Content)
			review["updated"] = entry.Updated
			// 作成したレビューの追加
			reviews = append(reviews, review)
		}
	}
	return reviews, nil
}

func parseComment(contents []content) string {
	for _, content := range contents {
		if content.TypeAttr == "text" {
			return content.Val
		}
	}
	return ""
}
