package main

import (
	"example.com/backend"
)

// func main() {
// 	backend.Run(":8080")
// }
func main (){
	a := backend.App{}
	a.Port = ":8080"
	a.Initialize()
	a.Run()
}