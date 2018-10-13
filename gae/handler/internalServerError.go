package handler

import (
	"context"
	"net/http"

	"google.golang.org/appengine/log"
)

// InternalServerError response 500
func InternalServerError(ctx context.Context, w http.ResponseWriter, err error) {
	log.Errorf(ctx, "Error: %v", err)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
