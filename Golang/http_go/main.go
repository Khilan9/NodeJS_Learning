package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// it represents which name should appear in json keys
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"-"`
	Address string   `json:"address,omitempty"`
	Hobbies []string `json:"hobbies"`
}

func httpGetOperation() {
	const myurl = "https://dummyjson.com/users/search?q=John"
	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var responsestring strings.Builder
	responsestring.WriteString(string(body))
	fmt.Println(responsestring.String())
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.ContentLength)

	result, _ := url.Parse(myurl)
	fmt.Println("Scheme ", result.Scheme)
	fmt.Println("Host ", result.Host)
	fmt.Println("Port ", result.Port())
	fmt.Println("Path ", result.Path)
	fmt.Println("RawQuery ", result.RawQuery)

	qparams := result.Query()
	for _, val := range qparams {
		fmt.Println(val)
	}
}

func httpPostOperation() {
	// const myurl = "https://jsonplaceholder.typicode.com/posts"
	url := "https://api.restful-api.dev/objects"

	payload := strings.NewReader(`{
   "name": "Apple MacBook Pro 16",
   "data": {
      "year": 2019,
      "price": 1849.99,
      "CPU model": "Intel Core i9",
      "Hard disk size": "1 TB"
   }
}`)

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func httpFormPostOperation() {
	data := url.Values{}
	data.Add("foo1", "bar1")
	data.Add("foo2", "bar2")
	resp, err := http.PostForm("https://postman-echo.com/post", data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func jsonmarshal() {
	// converts Go data structures into JSON format
	num := 42
	numJSON, _ := json.Marshal(num)
	fmt.Println(string(numJSON))

	text := "hello"
	textJSON, _ := json.Marshal(text)
	fmt.Println(string(textJSON))

	// Slices
	names := []string{"Alice", "Bob", "Charlie"}
	namesJSON, _ := json.Marshal(names)
	fmt.Println(string(namesJSON))

	// Maps
	user := map[string]interface{}{
		"name":   "John",
		"age":    30,
		"active": true,
	}
	userJSON, _ := json.Marshal(user)
	fmt.Println(string(userJSON))

	p := Person{
		Name:    "Jane",
		Age:     28,
		Hobbies: []string{"reading", "hiking"},
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))

	jsondatafromweb := []byte(`
	{
		"name" : "val1",
		"address": "addval",
		"hobbies":["reading","jjkk"]
	}
	`)
	checkvalid := json.Valid(jsondatafromweb)
	var p2 Person
	if checkvalid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsondatafromweb, &p2)
		fmt.Println(p2)
	}

	// convert json data into ley value pair only
	// here we have used interface as we are unaware that value is of which type
	// whether value is string or int or array
	var onlinedata map[string]interface{}
	json.Unmarshal(jsondatafromweb, &onlinedata)
	fmt.Printf("Type is %T and data is %v\n", onlinedata, onlinedata)
}
func main() {
	// httpGetOperation()
	// httpPostOperation()
	// httpFormPostOperation()
	jsonmarshal()
}
