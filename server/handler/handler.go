package handler

import (
	"net/http"
	"log"
	"encoding/json"
)

const (
	GetEnvPath = "/api/v1/env"
)

type Handler struct {
	server *http.Server
}

func NewHandler() *Handler {
	mux := http.NewServeMux()
	addr := ":8080"
	h := &Handler{server:&http.Server{Addr:addr, Handler:mux}}
	mux.HandleFunc(GetEnvPath, h.GetEnvHandler)
	mux.Handle("/", http.FileServer(http.Dir("./dist")))
	return h
}

func (h *Handler) Start() {
	log.Printf("http server started")
	h.server.ListenAndServe()
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
