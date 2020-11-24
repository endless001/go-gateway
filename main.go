package main

import (
	"go-gateway/conf"
	"go-gateway/server/http"
)

func main() {

	c := &conf.Config{}
	http.Init(c)

}
