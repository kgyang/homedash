package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kgyang/homedash/server/env"
)

const (
	GetEnvPath = "/api/v1/env"
)

type Handler struct {
	env    *env.Env
	server *http.Server
}

func NewHandler() *Handler {
	mux := http.NewServeMux()
	addr := ":8080"
	h := &Handler{server: &http.Server{Addr: addr, Handler: mux}, env: env.NewEnv()}
	mux.HandleFunc(GetEnvPath, h.GetEnvHandler)
	mux.Handle("/", http.FileServer(http.Dir("./dist")))
	return h
}

func (h *Handler) Start() {
	log.Printf("http server started")
	h.server.ListenAndServe()
}

func (h *Handler) GetEnvHandler(w http.ResponseWriter, r *http.Request) {
	resp := h.env.GetEnvData()
	jresp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s", err)
		return
	}
	w.Write(jresp)
}
