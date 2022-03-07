package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string
}

func main() {
	p1 := person{Name: "ayush"}
	p2 := person{Name: "shetty"}

	TP := []person{p1, p2}

	bs, err := json.Marshal(TP)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))
}
