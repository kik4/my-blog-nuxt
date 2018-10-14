package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kik4/my-portfolio-nuxt/gae/domain"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
)

// RequestQiitaItems is handler request items from qiita
func RequestQiitaItems(w http.ResponseWriter, r *http.Request) {
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

	e := domain.QiitaItemsEntity{}
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
