package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.naver33.com")
	println(0)
	if err != nil {
		fmt.Println(1)
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(2)
		log.Fatal(err)
	}
	println(3)
	fmt.Printf("%s\n", page)
	println(4)
}
