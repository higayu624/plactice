package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name string
	Age int
}

var tpl *template.template

func init() {
	tpl = template.Must(template.ParseFiles("tpls.html"))
}

func main() {
	p1 := Person{
		Name := "hogefuga",
		Age := 28,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}