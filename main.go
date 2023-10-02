package main

import (
	"Peko5/config"
	"Peko5/functions"
	"fmt"
)

func main() {
	fmt.Println("Thank you for using Peko5!")

	fmt.Println("First, you need to have a config file, checking...")
	c := config.ReadConfigFile("config.json")

	fmt.Println("Now, let's login to your subscription.")
	client := functions.InitProject(c)

	fmt.Println("Now, let's do something about it.")
	functions.ActionList(client, c)
}
