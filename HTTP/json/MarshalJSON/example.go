package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Artifact struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func main() {
	scroll := Artifact{
		Id:    "001-a",
		Name:  "sample",
		Level: 25,
	}
	fmt.Println("Unmarshalled struct: ", scroll)
	data, err := json.Marshal(scroll)
	if err != nil {
		log.Fatal(err)
	}
	var indentedData bytes.Buffer
	json.Indent(&indentedData, data, "", "	")
	fmt.Println(indentedData.String())
}
