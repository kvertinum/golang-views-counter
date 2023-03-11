package server

type ErrResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
