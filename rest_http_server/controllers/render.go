package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rewardStyle/campaign-service/lib/errutil"
	"github.com/rewardStyle/campaign-service/lib/log"
)

func RenderError(ctx context.Context, w http.ResponseWriter, err error) {
	errutil.TrackError(ctx, err)
	if pErr, ok := err.(*errutil.PublicError); ok {
		Render(ctx, w, pErr.Code, pErr)
	} else if pErr, ok := err.(*errutil.BadRequestError); ok {
		Render(ctx, w, http.StatusBadRequest, errutil.NewPublicErrorMessage(http.StatusBadRequest, pErr.Error()))
	} else {
		Render(ctx, w, http.StatusInternalServerError, errutil.NewPublicErrorMessage(http.StatusInternalServerError, err.Error()))
	}
}

func Render(ctx context.Context, w http.ResponseWriter, code int, value interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Note all the headers have to be written before WriteHeader's call
	if code >= 0 {
		w.WriteHeader(code)
	}

	if code == http.StatusNoContent || value == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		// last chance
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"errors\":[{\"message\":\"error serializing object\"}]}")
		log.ErrorMessage("could not marshal structure %v:%s", value, err)
		return
	}
}
