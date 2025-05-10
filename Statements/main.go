package main

import "fmt"



func main() {
	var firstname string = "David"
	lastname := "laid"

	if firstname == lastname {
		fmt.Println("firstname is equals to last name")
	}else if firstname == "cbum"{
		fmt.Println("firstname is not equals to last name")
	}else if firstname == "David" && lastname =="laid"{
		fmt.Println("firstname == last name")
	}else{
		fmt.Println("")
	}

}


func main1() {
    num := 10

    if num > 5 {
        fmt.Println("Number is greater than 5")
    }
}



func main2() {
    age := 16

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }
}

// short declaration inside IF-ELSE
func main3() {
    if n := 7; n%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // fmt.Println(n) //  Error: n is not accessible here
}

// switch

func main4() {
    day := 3

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Another day")
    }
}


func main5() {
    char := 'a'

    switch char {
    case 'a', 'e', 'i', 'o', 'u':
        fmt.Println("Vowel")
    default:
        fmt.Println("Consonant")
    }
}


// acts like if-else if

func main6() {
    num := 25

    switch {
    case num%2 == 0:
        fmt.Println("Even")
    case num%2 != 0:
        fmt.Println("Odd")
    }
}


// short declaration inside switch
func main7() {
    switch length := len("golang"); {
    case length < 5:
        fmt.Println("Short word")
    case length == 6:
        fmt.Println("Six letters")
    default:
        fmt.Println("Long word")
    }
}

// Go — it’s a special keyword used in switch to force the next case to run, even if it doesn’t match.

func main8() {
    num := 1

    switch num {
    case 1:
        fmt.Println("Case 1")
        fallthrough
    case 2:
        fmt.Println("Case 2")
        fallthrough
    case 3:
        fmt.Println("Case 3")
    default:
        fmt.Println("Default case")
    }
}

