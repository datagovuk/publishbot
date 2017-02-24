package main

import "fmt"
import "github.com/datagovuk/publishbot/adapters"

func main() {

	fmt.Println("Publishbot")
	adapters := adapters.List()
	fmt.Printf("%v\n", adapters)
}
