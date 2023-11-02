package server

import "net/http"
type Country struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Language string `json:"language"`
}

var countries []*Country = []*Country{}

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}

}