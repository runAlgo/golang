package main

import "fmt"

func main() {
	// simple switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("ON3")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Not Allowed!")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's Work Day")
	// }

	// type switch
	// 	whoAmI := func(i interface{}) {
	// 		switch i := i.(type) {
	// 		case string:
	// 			fmt.Println("This is String Value", i)
	// 		case int:
	// 			fmt.Println("This is Interger Value", i)
	// 		case bool:
	// 			fmt.Println("This is Boolean Value", i)
	// 		case float64:
	// 			fmt.Println("This is Fload32 Value", i)
	// 		}
	// 	}
	// 	whoAmI(4.3)
	// }

	// What is my DataType
	myDataType := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("This is String Value")
		case int:
			fmt.Println("This is Integer Value")
		default:
			fmt.Println("Other")
		}

	}
	myDataType(30)
}
