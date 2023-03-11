package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kvertinum01/views-counter/internal/app/images"
)

const (
	NAME_REQUIRED_CODE = 1
	TOO_BIG_NAME_CODE  = 2
	NAME_REQUIRED_ERR  = "'name' field is required"
	TOO_BIG_NAME_ERR   = "'name' cannot be longer than 16 characters"
)

func makeBadReqResp(cont string, code int, w http.ResponseWriter) {
	resp := ErrResponse{
		Code:  code,
		Error: cont,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResp)
}

func (s *Server) handleRandomImage(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("name") {
		makeBadReqResp(NAME_REQUIRED_ERR, NAME_REQUIRED_CODE, w)
		return
	}

	userName := r.URL.Query().Get("name")

	if len(userName) > 16 {
		makeBadReqResp(TOO_BIG_NAME_ERR, TOO_BIG_NAME_CODE, w)
		return
	}

	newUserValue, err := s.rst.AddByName(userName)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := images.NumToImage(s.imgDir, int(newUserValue))

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-Control", "max-age=20")
	w.Header().Set("cf-cache-status", "EXPIRED")

	w.Write(buf)
}
