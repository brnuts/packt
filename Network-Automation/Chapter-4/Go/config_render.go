package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	yfile, err := ioutil.ReadFile("items.yaml")

	if err != nil {

		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	for k, v := range data {

		fmt.Printf("%s -> %d\n", k, v)
	}

	sweat := Inventory{"wool", 17}

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweat)
	if err != nil {
		panic(err)
	}
}
