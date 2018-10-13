package domain

import "time"

// QiitaItemsEntity wraps Qiita data
type QiitaItemsEntity struct {
	Value     QiitaItemsValue
	UpdatedAt time.Time
}

// QiitaItemsValue is from /user/items
type QiitaItemsValue []struct {
	RenderedBody string `json:"rendered_body" datastore:"rendered_body,noindex"`
	Body         string `json:"body" datastore:"body,noindex"`
	// Coediting     bool      `json:"coediting"`
	// CommentsCount int       `json:"comments_count"`
	CreatedAt time.Time `json:"created_at"`
	// Group          interface{} `json:"group"`
	ID string `json:"id"`
	// LikesCount     int    `json:"likes_count"`
	// Private        bool   `json:"private"`
	// ReactionsCount int    `json:"reactions_count"`
	// Tags           []struct {
	// 	Name     string        `json:"name"`
	// 	Versions []interface{} `json:"versions"`
	// } `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	// User      struct {
	// 	Description       string `json:"description"`
	// 	FacebookID        string `json:"facebook_id"`
	// 	FolloweesCount    int    `json:"followees_count"`
	// 	FollowersCount    int    `json:"followers_count"`
	// 	GithubLoginName   string `json:"github_login_name"`
	// 	ID                string `json:"id"`
	// 	ItemsCount        int    `json:"items_count"`
	// 	LinkedinID        string `json:"linkedin_id"`
	// 	Location          string `json:"location"`
	// 	Name              string `json:"name"`
	// 	Organization      string `json:"organization"`
	// 	PermanentID       int    `json:"permanent_id"`
	// 	ProfileImageURL   string `json:"profile_image_url"`
	// 	TwitterScreenName string `json:"twitter_screen_name"`
	// 	WebsiteURL        string `json:"website_url"`
	// } `json:"user"`
	// PageViewsCount interface{} `json:"page_views_count"`
}
