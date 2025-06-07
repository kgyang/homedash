package handler

import (
	"net/http"
	"log"
)

const (
	GetEnvPath = "/api/v1/env"
)

type Handler struct {
	Server *http.Server
}

func NewHandler() *Handler {
	mux := http.NewServeMux()
	h := &Handler{Server:&http.Server{Addr:":8080", Handler:mux}}
	mux.HandleFunc(GetEnvPath, h.GetEnvHandler)
	return h
}

func (h *Handler) GetEnvHandler(w http.ResponseWriter, r *http.Request) {
	resp := &GetEnvResponse{Temperature:"27.0", Humidity:"88.2"}
	jresp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s", err)
		return
	}
	w.Write(jresp)
}
