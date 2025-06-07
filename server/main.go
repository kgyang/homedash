package main

import (
	"github.com/kgyang/homedash/server/handler"
)

func main() {
	h := handler.NewHandler()
	h.Start()
}
