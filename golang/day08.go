//style one : assigning function to a variable

// func test(x int) {
// 	println("Hello", x)
// }
// func main() {
// 	x := test
// 	x(5)
// 	test(5)
// }

//style two : Anonnamous (must be defined inside a function only)
//inferred variable or deeclaration cant be declared gobally
// func main() {
// 	test := func(x int) {
// 		println(x)
// 	}
// 	test(5)
// }

// var test = func(x int) { //inferred like this is possiable but ':=' do not work
// 	println(x)
// }

// func main() {
// 	test(5)
// }

//style three : return type of anomnmous with diffent calling strategy

// func main() {
// 	test := func(x int) int {
// 		return x * x
// 	}(5) //default argument
// 	fmt.Println(test)
// }

// func calculate(x , y int) (int, int) {//last arg ka type mention kiya tho bhi it works
// 	return x + y, x - y
// }
// func main() {
// 	sum, difference := calculate(10, 20)
// 	println("sum is ", sum, "differnce is ", difference)
// }

//******************Date and Time Function **************************
// func main() {
// 	Y, M, D := time.Now().Date()
// 	println(D, "/", M, "/", Y)
// 	println("Date is :", D)
// 	println("Month is :", M)
// 	println("Year is :", Y)
//

// func main() {
// 	currentdatetime := time.Now()
// 	day := currentdatetime.Day()
// 	month := currentdatetime.Month()
// 	year := currentdatetime.Year()
// 	Hour := currentdatetime.Hour()
// 	Min := currentdatetime.Minute()
// 	Sec := currentdatetime.Second()

// 	fmt.Printf("The report is printed on %d/%d/%d at %d:%d:%d", day, month, year, Hour, Min, Sec)

// }

// func main() {
// 	println("HELLO!!")
// 	time.Sleep(4* time.Second)
// 	println("My Name is Ketan :)")
// 	time.Sleep(4* time.Second)
// 	println("I'am Budding Data Scientiest")
// }

// func main() {
// 	num := 1
// 	// for num <= 5 { //works as a while
// 	// 	println(num)
// 	// 	num++
// 	// }

// 	// for i:= 0; i < 7; i++ { //works as normal for i must be initalized
// 	// 	println(i)
// 	// }
// 	for {
// 		//code
// 		println("this code runs default once ,then only after the condition is checked")
// 		num++
// 		if num > 5 {
// 			break //condition (terninates the loop)
// 		}
// 	}

// }

// func main() {
// 	num := 1
// 	var a int
// 	print("Enter the number")
// 	fmt.Scanln(&a)
// 	for num <= 10 {
// 		product := 5 * num
// 		fmt.Println(product)
// 		num++
// 	}

// }

//slicce in GO langauge

// func main() {
// 	var colors = []string{"Red", "Blue", "Green"} //slice wthout the size
// 	for index, val := range colors {              //range always returns 2 values index and value as showm
// 		fmt.Println(index, ":-", val)
// 	}

// 	for _, val := range colors { //range always returns 2 values index and value as showm
// 		fmt.Println(val)
// 	}

// 	fmt.Printf("%T", colors)

// }

// func main() {
// 	var arrayOfinteger = [5]int{1, 2, 3, 4, 5} //array is specified with the size else it becomes the slicew
// 	fmt.Println(arrayOfinteger)
// 	fmt.Printf("%T", arrayOfinteger)
// }

// func main() {
// 	weather := [3]string{"r", "t", "k"}
// 	weather[0] = "y"
// }

//for dynamic array use [...] in size
// func main() {
// 	var arr = [...]string{"Ketan", "Polawar"}
// 	fmt.Println(len(arr))

// }