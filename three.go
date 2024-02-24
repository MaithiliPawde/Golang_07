/*
package main
import "fmt"

func main() { inferred declaration

	 a ,b ,c,d :=1,2.9,3,"hello"
	 var e,f=6,"hell"
	  fmt.Println(a)
	  fmt.Println(b)
	  fmt.Println(c)
	  fmt.Println(d)
	}

package main
import "fmt"

	func main() {
	    var(
	        a int
	        b int=1
	        c string="hello"
	        )
	  fmt.Println(a)
	  fmt.Println(b)
	  fmt.Println(c)
	}

package main
import "fmt"
a=10; var b int =20

	func main() {
	    var(
	        a int
	        b int=1
	        c string="hello"
	        )
	  //fmt.Println(a,b,c)
	  fmt.Printf("a b and c are %d %d %s",a,b,c)
	  fmt.scanf("%d")
	  fmt.Println(a)
	}

package main
import "fmt"

	func main() {
	  var name,last string
	  fmt.Printf("enter ur name and lastname : ")
	  fmt.Scanf("%s%s",&name,&last)
	  fmt.Println("hello",name,last)
	}

package main
import "fmt"

	func main() {
	  var first string
	  fmt.Printf("enter string name and lname: ")
	  fmt.Scanln(&first&second)
	  var second string
	  fmt.Printf("enter string: ")
	  fmt.Scanln(&second)
	  fmt.Println(first+second)
	}

package main
import "fmt"

	func main() {
	  var first string
	  fmt.Printf("enter string name and lname: ")
	  fmt.Scanln(&first&second)
	  var second string
	  fmt.Printf("enter string: ")
	  fmt.Scanln(&second)
	  fmt.Println(first+second)
	}

package main
import "fmt"
func main() {
var a,b int =10,20
fmt.Println(a,b)
fmt.Println(a+b)
}
package main
import "fmt"

	func main() {
	  var first, second string
	  fmt.Printf("enter string name and lname: ")
	  fmt.Scanln("%s%s",&first,&second)
	  fmt.Println(first+second)
	}

package main
import "fmt"

	func main() {
	  var i,j string="hello","maith"
	  fmt.Print(i)
	  fmt.Print(j)
	}

package main
import "fmt"

	func main() {
	  var i,j,k string="maith","RCOEM","DS"
	  var l int =20
	  fmt.Println("name:",i)
	  fmt.Println("age:",l)
	  fmt.Println("college:",j)
	  fmt.Println("branch:",k)
	}

package main
import "fmt"

	func main() {
	  var i,j,k string="maith","RCOEM","DS"
	  var l int =20
	  fmt.Printf("Name: %s\tAge: %d\nCollege: %s\tBranch: %s\t", i, l, j, k)
	}

package main
import "fmt"
func main() {
var a string ="maith"

	  fmt.Printf("the value of a is '%s'",a)
	}

package main

	import {
	    "fmt"
	}

	func main() {
	    const name,dept string ="gfg","cs"
	    fmt.Printf("%s is a %s",name, dept)
	}

package main

import (

	"fmt"

)
const (

	C = "pawde"

)
//f:=10  global variable cannot be declared using':='

	func main() {
	        var str = "maith_pawde"
	    fmt.Printf("The string is %s \n", str) //string
	    var num1 int = 20
	    fmt.Printf("The string is %d \n", num1) //int
	    fmt.Printf("The string is %b \n", num1) //prints binary value
	    var num3 float32 = 1.22
	    fmt.Printf("The string is %g \n", num3) //float num %f can also be employeed
	    var num2 = 4 + 1i
	    fmt.Printf("The string is %e \n", num2) //complex number

	    fmt.Printf("The string is %v,%v,%v,%v \n", str, num3, num2, num1) //it uses the original type and prints the values

ox - hexadecimal
0 or 0ois octal number
a=b assignmemt
a==b comparison if same then true retuen karega
}
*/
package main

import "fmt"

func main() {

	fmt.Println(15 == 017) //true
	fmt.Println(15 == 0xf) //true
	var a int = 2
	a++
	println(a)
}
