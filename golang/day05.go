// //******************** Recursion *********************

// /*func recurse(){
// 	---
// 	---
// 	recurse()
// }*/

// // func countdown(n1 int) {
// // 	fmt.Println(n1)
// // 	if n1 > 1 {
// // 		countdown(n1 - 1)
// // 	}

// // }

// // func factorial(num int) int {
// // 	if num == 1 || num == 0 {
// // 		return num
// // 	}
// // 	return num * factorial(num-1)
// // }

// func sum(number int) int {
// 	if number == 0 {
// 		return 0
// 	} else {
// 		return number + sum(number-1)
// 	}
// }

// /*func(){

// }*/

// var greet = func() {
// 	fmt.Println("\nHello,how are you")
// }
// var sum1 = func(n1, n2 int) {
// 	result := n1 + n2
// 	fmt.Println("Sum is:", result)
// }

// var sum3 = func(n1, n2 int) int {
// 	result := n1 + n2
// 	return result
// }
// var sub3 = func(n1, n2 int) int {
// 	result := n1 - n2
// 	return result
// }
// var mul3 = func(n1, n2 int) int {
// 	result := n1 * n2
// 	return result
// }
// var div3 = func(n1, n2 int) int {
// 	result := n1 / n2
// 	return result
// }