package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	l, _ := resp.Body.Read(bs)
	fmt.Println(l)
	fmt.Println(string(bs))

}
