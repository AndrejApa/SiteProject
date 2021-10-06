package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Car struct {
	Name, Model string
	Price, Year uint
	Conf        []string
}

func (c *Car) getInfo() string {
	return fmt.Sprintf("Car brand is: %s, model is : %s", c.Name, c.Model)
}
func home(w http.ResponseWriter, r *http.Request) {
	bmw := Car{Name: "BMW", Model: "320", Price: 20000, Year: 2020, Conf: []string{"Color: Black", "Mileage: 38000", "Location: Minsk"}}
	//fmt.Fprintf(w, bmw.getInfo())

	tmpl, err := template.ParseFiles("SiteProject/templates/home.html")
	err = tmpl.Execute(w, bmw)
	if err != nil {
		return
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from about page")
}
func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about/", about)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}

func main() {
	handleRequest()
}
