package main

import (
	"go-gateway/conf"
	"go-gateway/http"
)

func main() {

	c := &conf.Config{}
	http.Init(c)

}
