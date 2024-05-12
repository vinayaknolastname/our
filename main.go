package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
)

func main() {
	data := make(map[string]*string)
	v := "Alice"

	data["20"] = &v
	s, u := data["20"]
	fmt.Println(s)
	fmt.Println(u)

}
