package main

import (
	"net/http"

	r "duarch/gobag/routes"

	_ "github.com/lib/pq"
)

func main() {
	r.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
