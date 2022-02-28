package routes

import (
	t "duarch/gobag/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", t.Index)
	http.HandleFunc("/new", t.New)
	http.HandleFunc("/insert", t.Insert)

}
