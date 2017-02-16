package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type App struct {
	Id string `json:"id"`
	Title string `json:"title"`
}

func main() {
	fmt.Println("json")
	data := []byte(`
		{
			"id": "ABCD123",
			"title": "Awesome json read"
		}
	`)

	var app App
	err := json.Unmarshal(data, &app)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(app.Id)
	fmt.Println(app.Title)

}
