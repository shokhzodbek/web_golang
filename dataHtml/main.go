package main

import (
	"html/template"
	"net/http"
)

type productSpe struct {
	Size string
	Weidth int
	Desc string
}

type product struct {
	ProId int
	Name string 
	Cost float32
	Spec productSpe
}

var tpl *template.Template
var prod1 product
func main() {
   prod1 = product{
	Name: "Product",
	Cost: 45.54,
	Spec: productSpe{
		Size: "34x5",
		Weidth: 43,
		Desc: "Sdfefgdsad",
	},
   }

	tpl,_ = template.ParseGlob("template/*.html")
	http.HandleFunc("/product",productHandler)
	http.ListenAndServe("localhost:5000",nil)
}

func productHandler(res http.ResponseWriter,req *http.Request) {
	tpl.ExecuteTemplate(res,"product.html",prod1)

}