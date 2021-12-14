package main

import (
	"fmt"
	"test-app/greet"
	"test-app/weather"
)

func main() {
	fmt.Println("main: ", greet.Words)
	fmt.Println("weather: ", weather.Greetname)
}
