package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


func main1() {
	fmt.Println("Testing Web-Request")
	response,err:=http.Get("https://jsonplaceholder.typicode.com/todos/")
	if err != nil{
		fmt.Println("Error Which Raise Through URL Is ",err)
		return
	}
	// fmt.Println("RESPONSE IS",response)

	defer response.Body.Close()

	// res := response.Body
	// fmt.Println(res)


	// Converting array of bytes into string

	data,err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("RESPONSE WHICH WE GET IS",string(data))


}

func main(){
	fmt.Println("LEARNING URLS")
	my_url := "https://jsonplaceholder.typicode.com/todos/"  // this is string
	fmt.Printf("Type of urls is %T \n",my_url)

	parsed_url ,err :=  url.Parse(my_url)
	if err != nil{
		fmt.Println(err)

	}

	fmt.Printf("Now Type of URL is %T \n",parsed_url)
	fmt.Println(parsed_url.RawQuery)
	fmt.Println(parsed_url.Host)
	fmt.Println(parsed_url.Hostname())
	fmt.Println(parsed_url.Path)
	fmt.Println(parsed_url.Scheme)


	// YOU CAN ALSO MODIFIES IT.

	// WIDELY USED IN PRODUCTION LEVEL
	


}