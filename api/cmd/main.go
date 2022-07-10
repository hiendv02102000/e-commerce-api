package main

import "api/api/router"

func main() {
	r := router.NewRouter()
	r.Engine.Run()

}
