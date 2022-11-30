package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Info().Msg("hello world are triggered")

	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Hello World",
	})
}
