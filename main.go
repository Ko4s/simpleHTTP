package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

var mainView *template.Template

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html") // do tego wrócumi jak będziemy korzystać z template html

	user := User{
		Name: "Piotrek",
	}

	mainView.Execute(w, user)
}

func main() {
	var err error
	mainView, err = template.ParseFiles("./index.html")
	checkErr(err)

	http.HandleFunc("/", mainPage)

	fmt.Println("Server starting on port 3000....")
	http.ListenAndServe("localhost:3000", nil)
}

//Requesty http -> body <- cała nasza strona html a header <- informacje o statusie requesta czyli np kod 200 że ok i informacje o typie dokumentu

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
