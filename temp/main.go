package main

import (
	"html/template"
	"net/http"
)


var tpl *template.Template

func main() {

	tpl,_= template.ParseGlob("templates/*.html")

	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/about",aboutHandler)
	http.HandleFunc("/login",loginHandler)
	http.ListenAndServe("localhost:8000",nil)



}

func indexHandler(res http.ResponseWriter,req *http.Request)  {
	tpl.ExecuteTemplate(res,"index.html",nil)
}
func aboutHandler(res http.ResponseWriter,req *http.Request)  {
	tpl.ExecuteTemplate(res,"about.html",nil)
	
}
func loginHandler(res http.ResponseWriter,req *http.Request)  {
	tpl.ExecuteTemplate(res,"login.html",nil)
}