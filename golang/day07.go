// Nested if else block
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	var year int
// 	println("enter the value for year:")
// 	fmt.Scanln(&year)
// 	if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%100 == 0 && year%400 == 0) {
// 		if year == 0 {
// 			println("invalid year")
// 		} else {
// 			println(year, " is a leap year")
// 		}
// 	} else {
// 		println(year, " is not a leap year")
// 	}
// }

// func main() {
// 	x := 8
// 	y := 10

// 	if x != y {
// 		if x < y {
// 			println(x, "is less than", y)
// 		} else if x > y {
// 			println(x, "is more than", y)
// 		} else {
// 			println(x, " is equal to ", y)
// 		}

// 	}
// }

// func main() {
// 	var time int
// 	println("Enter the time:")
// 	fmt.Scanln(&time)
// 	if 0 < time && time < 24 && time != 0 {
// 		if time < 10 {
// 			println("good moring")
// 		} else if time > 10 && time <= 20 {
// 			println("good day")
// 		} else {
// 			println("good evening")
// 		}

// 	}

// }

//******************Switch Statements****************

// func main() {
// 	M := time.Now()    //get system date/month/year
// 	Month := M.Month() //corp the Month from the date returned
// 	fmt.Printf("%T\n", M)
// 	switch Month {
// 	case 1:
// 		println("January")
// 	case 2:
// 		println("Febrary")
// 	case 3, 4:
// 		println("March", "April")
// 	case 5:
// 		println("May")
// 	}

// 	switch Month {
// 	case 1, 3, 5, 7, 8, 10:
// 		fmt.Println("31 days")
// 		fallthrough //goes to next case ,and direcly runs the next case without even verifing the nect case
// 	case 4, 6, 9, 11:
// 		fmt.Println("30 days")
// 	default:
// 		fmt.Println("this")

// 	}

// 	//var x interface{} = "RKNEC"
// 	var x interface{}
// 	switch i := x.(type) {
// 	case nil:
// 		fmt.Printf("type of x is %T", i)
// 	case int:
// 		fmt.Printf("type of x is int")
// 	case float32:
// 		fmt.Printf("type of x is float32")
// 	case bool, string:
// 		fmt.Print("Type of x is string or bool")
// 	default:
// 		fmt.Print("dont know the type")
// 	}

// }
//
//	func main() {
//		i := 0
//
// loop:
//
//		fmt.Println(i)
//		i++
//		if i < 5 {
//			goto loop
//		}
//		println("ends the loop")
//	}
// func main() {
// 	var str string = "Hello World"
// 	var substr string = "Wor"
// 	if strings.Contains(str, substr) { //checks the contains of substr in str
// 		fmt.Printf("String (%s) contains substring (%s)\n", str, substr)
// 	} else {
// 		fmt.Printf("String (%s) do not contain substring (%s)\n", str, substr)
// 	}
// 	println(strings.ToUpper(str)) //to upper
// 	println(strings.ToLower(str)) //to lower
// 	if (strings.Index(str, "W")) < 0 {
// 		println("does not exists")
// 	} else {
// 		println("exists st the position", (strings.Index(str, "W")))
// 	} //position of the "W" in str
// }