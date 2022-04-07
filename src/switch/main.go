package main

import "fmt"

func main() {
	//Uma observação importante do switch em Go é que ele para quando um caso é verdadeiro. Se você quiser fazer com que vários casos possam ser verdadeiros, basta usar a palavra-chave "fallthrough"
	numDaysInMonth := 30
	switch {
	case numDaysInMonth >= 28 && numDaysInMonth <= 29:
		fmt.Println("February")
	case numDaysInMonth == 30:
		fmt.Println("April, June, September, November")
	case numDaysInMonth == 31:
		fmt.Println("January, March, May, July, August, October,December")
	}
	operatingSystem := "Windows"
	switch operatingSystem {
	case "Windows":
		fmt.Println("Made by Microsoft")
	case "OS X":
		fmt.Println("Made by Apple")
	case "Linux":
		fmt.Println("Made by a lot of clever volunteers")
	default:
		fmt.Println("I don't recognize that operating system. GEEK!")
	}
}
