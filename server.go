package main

import (
	"log"
	"net/http"
	"fmt"
	//"io/ioutil"
)
func hello(w http.ResponseWriter, r *http.Request) {
    // if r.URL.Path != "/" {
    //     http.Error(w, "404 not found.", http.StatusNotFound)
    //     return
    // }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "./static/form.html")
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        name := r.FormValue("name")
        emailaddress := r.FormValue("emailaddress")
        fmt.Fprintf(w, "Name = %s\n", name)
        fmt.Fprintf(w, "Address = %s\n", emailaddress)
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}
 
func main() {
    http.HandleFunc("/", hello)
 
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}