package main

import "fmt"

func main() {
	//var color map[string]string
	//colors := make(map[string]string)

	color := map[string]string{
		"red": "0xff0000",
	}
	color["white"] = "0xffffff"

	fmt.Println(color)
}
