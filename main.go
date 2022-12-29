package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	fmt.Println("Hello OpenLibrary")
	var book interface{}
	url := "https://openlibrary.org/works/OL17590212W.json"
	getJson(url, book)
	//resp,err := http.Get()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer resp.Body.Close()
	//bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//dec := json.NewDecoder(bytes.NewReader(bodyBytes))
	//dec.UseNumber()
	fmt.Println(book)
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body) // response body is []byte
	fmt.Println(string(body))
	return json.NewDecoder(r.Body).Decode(target)
}
