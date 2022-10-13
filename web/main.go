package main

import (
	"fmt"
	"net/http"
)


func main() {
   http.HandleFunc("/",indexHandler)
   http.HandleFunc("/about",aboutHandler)
   http.ListenAndServe("localhost:8000",nil)
}

func indexHandler(res http.ResponseWriter,req *http.Request) {

	fmt.Fprintf(res,"path %s",req.URL.Path)

}

func aboutHandler(res http.ResponseWriter,req *http.Request)  {

	fmt.Fprintf(res,"HI it is about page")
	
}