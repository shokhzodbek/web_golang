package main

import (
	"html/template"
	"net/http"
)


type User struct {
	Name string
	Lang string
	Member bool
}
var tmpl *template.Template
var user User

func main() {
   user = User{
	Name: "Shokhz",
	Lang: "English",
	Member: false,
   }
   http.HandleFunc("/",userHandler)

   tmpl,_= template.ParseGlob("template/*.html")
	http.ListenAndServe("localhost:8000",nil)

}


func userHandler(res http.ResponseWriter,req *http.Request)  {
tmpl.ExecuteTemplate(res,"user.html",user)
}