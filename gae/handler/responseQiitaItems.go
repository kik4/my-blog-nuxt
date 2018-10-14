package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kik4/my-portfolio-nuxt/gae/domain"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// ResponseQiitaItems is handler response qiita items
func ResponseQiitaItems(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/qiita/items" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)
	e := domain.QiitaItemsEntity{}
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
