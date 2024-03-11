// 	//***************OPERATORS***********************
// 	// var a float64
// 	// var b string = "maith"
// 	// var area float64
// 	// //var c float32 = 30.67
// 	// //var c int = 10
// 	// println("/nenter the number")
// 	// fmt.Scanf("%d", &a)
// 	// // print(a % 2)
// 	// // println(c)
// 	// // println(c & a)
// 	// // println(c | a)
// 	// // println(c ^ a)
// 	// // println(a > 5)
// 	// // println(a == 10)
// 	// //x>>=3

// 	// //sizeof is defined in 'unsafe' libraray
// 	// println(unsafe.Sizeof(a))
// 	// println(unsafe.Sizeof(b))
// 	// fmt.Println(reflect.TypeOf(a))
// 	// defer fmt.Println(reflect.TypeOf(b))
// 	// defer fmt.Println(reflect.ValueOf(a).Kind())
// 	// fmt.Println(reflect.ValueOf(b).Kind())
// 	// area = PI * a
// 	// println(area)

// 	// var ftemp float64
// 	// var ctemp float64
// 	// fmt.Print("Enter the temperatur in celcius:")
// 	// fmt.Scanf("%f", &ftemp)
// 	// ctemp = (ftemp * 1.8) + 32
// 	// fmt.Printf("The temperature in farenheit : %.2f", ctemp)
// 	// fmt.Print(" its maith)
// 	// var dd int = 20
// 	// var mm int = 02
// 	// var yy int = 2024
// 	// var str string
// 	// var flt float32 = 90.78
// 	// str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
// 	// io.WriteString(os.Stdout, str)

// 	// fmt.Printf("\nenter the string:")
// 	// fmt.Scan(&str, &dd, &flt)
// 	// fmt.Printf("\nResult:%s,%d\n", str, dd)

// 	//Why type casting is important??
// 	// var a int = 123
// 	// var b uint = 0
// 	// b = uint(a) //type cast nahi kiya tho error
// 	// c := float64(-10.00)
// 	// fmt.Printf("\t%g", math.Abs(c))
// 	// fmt.Printf("\t%g", math.Sqrt(math.Abs(c)))
// 	// //assign value to b of a
// 	// fmt.Printf("\na=%d,b=%d\n", a, b)

// 	// var x int = 225
// 	// var r float32
// 	// r = float32(math.Sqrt(float64(x)))
// 	// print(r)

// 	// var x int = 42
// 	// var y float64 = float64(x)
// 	// var z uint = uint(y)
// 	// fmt.Printf("\nValue of x is %d and type is %T\n", x, x)
// 	// fmt.Printf("\nValue of y is %2f and type is %T\n", y, y)
// 	// fmt.Printf("\nvalue of z is %d and type is %T\n", z, z)
// 	// var num1 int
// 	// var num2 int
// 	// var avg float64
// 	// num1 = 10
// 	// num2 = 5
// 	// avg = (float64(num1) + float64(num2)) / 2
// 	// //avg = (num1+num2) / 2 will give error as int cant be stored in float
// 	// fmt.Printf("Average of %d and %d is %.2f\n", num1, num2, avg)

// 	//****Selection Iterations Sequencing ****
// 	/*if condition/expression {
// 		code to be executed
// 	}*/

// 	//a := 20
// 	// b := 10

// 	// if !(a > b) {
// 	// 	fmt.Println("Hi its ketan 1")
// 	// }
// 	// if 20 > 10 {
// 	// 	fmt.Println("Hi its ketan 2")
// 	// }
// 	// if true {
// 	// 	fmt.Println("Hi its ketan 3")
// 	// }

// 	// if b := 10; a > b { //b would be block variable ,only one ststement kar sakte a,b=10,20 works
// 	// 	fmt.Println("a is less than b")
// 	// }
// 	//var age int
// 	//var name string
// 	// print("enter the values")
// 	// if _, err := fmt.Scan(&name, &age); err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }
// 	// println("Name:", name)
// 	// println("Age:", age)

// 	// print("enter the values")
// 	// if a, err := fmt.Scan(&name, &age); err == nil {//fmt.scan jo value return karta hai voh return karata hai hum determine nahi kar sakta hai
// 	// 	fmt.Println("Hey", a)
// 	// }
// 	// println("Name:", name)
// 	// println("Age:", age)

// 	// var a, b, c int
// 	// println("Enter the values of a,b,c")
// 	// fmt.Scanln(&a, &b, &c) //eak line mai hi inputs likho
// 	// // fmt.Scanln(&a)//harr statement enter k badd input fmt.Scanln(&a,&b) dono eak line mai likh sakte
// 	// // fmt.Scanln(&b)
// 	// // fmt.Scanln(&c)
// 	// if (a > b) && (a > c) {
// 	// 	println("The Largerst value is :", a)
// 	// }
// 	// if (b > a) && (b > c) {
// 	// 	println("The Largerst value is :", b)
// 	// }
// 	// if (c > b) && (c > a) {
// 	// 	println("The Largerst value is :", c)
// 	// }
// 	/*if condn {

// 	}else{ //else must be written in same line of ending } of if

// 	}*/

// 	// a := 20
// 	// b := 10

// 	// if a > b {
// 	// 	println("a is greater than b")
// 	// } else {
// 	// 	println("b is greater the a")
// 	// }

// 	// var a int
// 	// println("enter the number")
// 	// fmt.Scan(&a)
// 	// // if a%2 == 0 {
// 	// // 	println("It is even number")
// 	// // } else {
// 	// // 	println("it is odd number")
// 	// // }

// 	// if a < 0 {
// 	// 	a = -1 * a
// 	// 	fmt.Printf("Abs value of a is: %d", a)
// 	// } else {
// 	// 	println("the value of a is ", a)
// 	// }