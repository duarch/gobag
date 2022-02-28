package routes

import (
	t "duarch/gobag/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", t.Index)
}
