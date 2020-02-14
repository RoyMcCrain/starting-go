package main

import (
	"fmt"

	"github.com/RoyMcCrain/starting-go/zoo/animals"
)

func main() {
	fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
