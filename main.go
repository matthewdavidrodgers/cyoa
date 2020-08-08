package main

import (
	"fmt"

	"github.com/matthewdavidrodgers/cyoa/adventure"
	"github.com/matthewdavidrodgers/cyoa/server"
)

func main() {
	story, err := adventure.Load("./example.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = server.SetupServer(story)
	if err != nil {
		fmt.Println("Could not start server")
		fmt.Println(err)
		return
	}
}
