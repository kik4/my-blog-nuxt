package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func main() {
	http.HandleFunc("/qiita/items/request", handle)
	http.HandleFunc("/qiita/items", handle2)
	appengine.Main()
}

// QiitaItemsData is from /user/items
type QiitaItemsData []struct {
	RenderedBody  string    `json:"rendered_body" datastore:"rendered_body,noindex"`
	Body          string    `json:"body" datastore:"body,noindex"`
	Coediting     bool      `json:"coediting"`
	CommentsCount int       `json:"comments_count"`
	CreatedAt     time.Time `json:"created_at"`
	// Group          interface{} `json:"group"`
	ID             string `json:"id"`
	LikesCount     int    `json:"likes_count"`
	Private        bool   `json:"private"`
	ReactionsCount int    `json:"reactions_count"`
	// Tags           []struct {
	// 	Name     string        `json:"name"`
	// 	Versions []interface{} `json:"versions"`
	// } `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       string `json:"description"`
		FacebookID        string `json:"facebook_id"`
		FolloweesCount    int    `json:"followees_count"`
		FollowersCount    int    `json:"followers_count"`
		GithubLoginName   string `json:"github_login_name"`
		ID                string `json:"id"`
		ItemsCount        int    `json:"items_count"`
		LinkedinID        string `json:"linkedin_id"`
		Location          string `json:"location"`
		Name              string `json:"name"`
		Organization      string `json:"organization"`
		PermanentID       int    `json:"permanent_id"`
		ProfileImageURL   string `json:"profile_image_url"`
		TwitterScreenName string `json:"twitter_screen_name"`
		WebsiteURL        string `json:"website_url"`
	} `json:"user"`
	// PageViewsCount interface{} `json:"page_views_count"`
}

// Entity wraps Qiita data
type Entity struct {
	Value     QiitaItemsData
	UpdatedAt time.Time
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/qiita/items/request" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("https://qiita.com/api/v2/users/kik4/items")
	if err != nil || resp.StatusCode != http.StatusOK {
		InternalServerError(ctx, w, err)
		return
	}

	e := Entity{}
	err = json.NewDecoder(resp.Body).Decode(&e.Value)
	if err != nil {
		InternalServerError(ctx, w, err)
		return
	}
	e.UpdatedAt = time.Now()

	fmt.Fprintf(w, "%#v", e)

	key := datastore.NewKey(ctx, "qiita", "kik4", 0, nil)
	key, err = datastore.Put(ctx, key, &e)
	if err != nil {
		InternalServerError(ctx, w, err)
		return
	}
	fmt.Fprintln(w, key)
}

// InternalServerError response 500
func InternalServerError(ctx context.Context, w http.ResponseWriter, err error) {
	log.Errorf(ctx, "Error: %v", err)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

func handle2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/qiita/items" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)
	e := Entity{}
	key := datastore.NewKey(ctx, "qiita", "kik4", 0, nil)
	err := datastore.Get(ctx, key, &e)
	if err != nil {
		InternalServerError(ctx, w, err)
		return
	}

	res, err := json.Marshal(e)
	if err != nil {
		InternalServerError(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}
