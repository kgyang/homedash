package main

import (
	"github.com/kgyang/homedash/handler"
)

func main() {
	h := handler.NewHandler()
	h.Server.ListenAndServ()
}
