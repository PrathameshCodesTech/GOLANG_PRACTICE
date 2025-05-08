Go program are organized into packages
Go standard library, provides different core packages for us to use
packages is collection of source file.

"fmt" is format, is one of these, which we we can use by importing it

Scopes:
local variable , global variable and packages level variable.

variable shadowing :- uses nearest variable if variables names are same.

_________________________________________________________________________

IOTA:

iota is a predeclared identifier in Go used inside const blocks. It automatically increments by 1, starting from 0, each time the const block is compiled.

It's most commonly used for defining enumerated constants (like enums in other languages).

package main

import "fmt"

const (
	Customer = iota // 0
	Product         // 1
	Order           // 2
)

func main() {
	fmt.Println(Customer, Product, Order)
}

How iota Works:
It resets to 0 in every const block.

If you skip a value, you can use _ to ignore it.

You can also use arithmetic operations with it.

____________________________________________________________________

case sensitive language

_____________________________________________________________________

## ✅ Go Comparison Operators

Go supports standard comparison operators:

| Operator | Description           | Example (`a`, `b`) |
| -------- | --------------------- | ------------------ |
| `==`     | Equal to              | `a == b`           |
| `!=`     | Not equal to          | `a != b`           |
| `<`      | Less than             | `a < b`            |
| `<=`     | Less than or equal    | `a <= b`           |
| `>`      | Greater than          | `a > b`            |
| `>=`     | Greater than or equal | `a >= b`           |

> ⚠️ **Note:** Only **comparable types** can be compared using `==` and `!=`.
> Go **does not allow** comparison of slices, maps, or functions **directly**, except against `nil`.

_____________________________________________________________________
Arithmetic Operators in Go

| Operator | Description         | Example (`a = 10`, `b = 3`) | Result |
| -------- | ------------------- | --------------------------- | ------ |
| `+`      | Addition            | `a + b`                     | `13`   |
| `-`      | Subtraction         | `a - b`                     | `7`    |
| `*`      | Multiplication      | `a * b`                     | `30`   |
| `/`      | Division (int)      | `a / b`                     | `3`    |
| `%`      | Modulus (remainder) | `a % b`                     | `1`    |

_____________________________________________________________________
Arithmetic Assignment Operators in Go

| Operator | Equivalent To | Description         | Example (`a := 10`) | Result |
| -------- | ------------- | ------------------- | ------------------- | ------ |
| `+=`     | `a = a + b`   | Add and assign      | `a += 5`            | `15`   |
| `-=`     | `a = a - b`   | Subtract and assign | `a -= 3`            | `7`    |
| `*=`     | `a = a * b`   | Multiply and assign | `a *= 2`            | `20`   |
| `/=`     | `a = a / b`   | Divide and assign   | `a /= 4`            | `2`    |
| `%=`     | `a = a % b`   | Modulus and assign  | `a %= 3`            | `1`    |

_____________________________________________________________________
✅ Logical Operators in Go
Logical operators are used to combine or invert boolean expressions. These only work with bool values.

AND (&&)
| A     | B     | A && B |
| ----- | ----- | ------ |
| true  | true  | true   |
| true  | false | false  |
| false | true  | false  |
| false | false | false  |

OR (||)
| A     | B     | A || B |
| ----- | ----- | -------- |
| true  | true  | true     |
| true  | false | true     |
| false | true  | true     |
| false | false | false    |

NOT (!)
| A     | !A    |
| ----- | ----- |
| true  | false |
| false | true  |

_______________________________________________________________________










