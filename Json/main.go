package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int	`json:"age"`
}

func main() {
	fmt.Println("now we are learnign json")
	person := Person{Name: "Pratghamesh",Age: 22}
	fmt.Println(person.Age)
	fmt.Println(person.Name)

	// convert person into json encoding // (marshalling)

	json_data,err:=json.Marshal(person)
	if err != nil{
		fmt.Println("Error Marshalling",err)
		return
	}
	fmt.Println("Person data is: ",string(json_data))

	// Demarshalling

	var personData Person
	json.Unmarshal(json_data, &personData )

	//FIXME - 

}