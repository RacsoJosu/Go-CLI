package main

import (
	"github.com/RacsoJosu/Api-Rest/server"
)

func main() {

	srv := server.New(":3000")

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
