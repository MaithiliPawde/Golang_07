/*package main //recurssion sum
import "fmt"
func sum(number int)int{ 
    if number==0{
        return 0
    }else{
        return number+sum(number-1)
    }
}
func main(){
    var number int
    fmt.Println("enter the number ")
    fmt.Scanln(&number)
    if number<0{
        fmt.Println("invalid input please try again !")
    }else{
    var result=sum(number)
    fmt.Println("sum",result)
    }
}
package main // recurssion factorial
import "fmt"
func factorial(number int)int{ 
    if number==0{
        return 1
    }else{
        return number*factorial(number-1)
    }
}
func main(){
    var number int
    fmt.Println("enter the number ")
    fmt.Scanln(&number)
    var result=factorial(number)
    fmt.Println("factorial",result)
}

package main // annonomous function
import "fmt"
func main(){
    var greet= func(){
        fmt.Println("hello")
    }
greet()
}

package main // annonomous function
import "fmt"
func sum(){
    greet()
}
func main(){
    var greet= func(){
        fmt.Println("hello")
    }
sum()
}
package main // annonomous function just like var a int=10 so i can do var b:=a here a =greet nad b =welcome
import "fmt"
func main(){
    var greet= func(){

        fmt.Println("hello")
    }
    welcome:=greet
    welcome()
}
package main // annoomous function sum 
import "fmt"
func main(){// n1 ko int nhi diya oth works bt if n1 ko diya hai aur n2 ko nhi diya hai toh error dega 
    var sum=func(n1,n2 int){
        result:=n1+n2
        fmt.Println("sum is",result)
    }
    sum(5,3)
    fmt.Printf("%T",sum)
}
//return value froom ananomous function
c
package main // aas argument to othwer function
import "fmt"
var sum=0
func findsquare(num int) int {
    square := num*num
    return square
}
func main(){
    sum:=func(num1 int ,num2 int)int{
        return num1+num2
    }
    result:=findsquare(sum(5,9))
    fmt.Println("result is ",result)
}
// package main

// import "fmt"

// func calculate() func() int {
// 	num := 1
// 	return func() int {
// 		num = num + 2
// 		return num
// 	}
// }

// func main() {
// 	odd := calculate()
// 	fmt.Println(odd())
// 	fmt.Println(odd())
// 	fmt.Println(odd())

// 	odd2 := calculate()
// 	fmt.Println(odd2())
// }

// package main

// import "fmt"

// func greet() func() string {
// 	name := "john"
// 	return func() string {
// 		name = "Hi " + name
// 		return name
// 	}
// }

// func main() {
// 	messages := greet()
// 	//ketan
// 	fmt.Println(messages())

// }
// package main

// import "fmt"

// func displayNumbers() func() int {
// 	number := 0
// 	return func() int {
// 		number++
// 		return number
// 	}
// }
// func main() {
// 	num1 := displayNumbers()
// 	fmt.Println(num1())
// 	num2 := displayNumbers()
// 	fmt.Println(num2())
// 	num3 := displayNumbers()
// 	fmt.Println(num3())

// package main

// import "fmt"

// func displayNumbers() func() int {
// 	number := 0
// 	return func() int {
// 		number++
// 		return number
// 	}
// }
// func main() {
// 	num1 := displayNumbers()
// 	fmt.Println(num1())
// 	num2 := displayNumbers()
// 	fmt.Println(num2())
// 	fmt.Println(displayNumbers()())
// 	num3 := displayNumbers()
// 	fmt.Println(displayNumbers())
// 	fmt.Println(num3())

// }

//*********febo using recursion and ananymous***********

// package main

// import "fmt"

// func main() {

// 	var fib func(n int) int
// 	fib = func(n int) int {
// 		if n < 2 {
// 			return n
// 		}
// 		return fib(n-1) + fib(n-2)
// 	}
// 	fmt.Println(fib(7))
// }
