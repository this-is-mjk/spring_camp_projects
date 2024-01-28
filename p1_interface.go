package main 

import (
	"fmt"
)

type Cipher interface {
	Encrypt() []int
}
func Encrypter(c Cipher) {
	fmt.Println(c.Encrypt())
}

type myString string
type myArray []int
type myMap map[rune]int

func (ms myString) Encrypt() []int {
	returnArr := []int{}
	for _, value := range ms {
		value := int(value)
		if 65 <= value && value <= 90{
			returnArr = append(returnArr, value - 64)
		} else if 97 <= value && value <= 122 {
			returnArr = append(returnArr, value - 96)
		} else {
			returnArr = append(returnArr, 0)
		}
	}
	return returnArr
}

func (ma myArray) Encrypt() []int {
	returnArr := []int{}
	for number := range ma {
		//fmt.Println(ma[number])
		if ma[number] % 2 == 0 {
			returnArr = append(returnArr, int(ma[number]/2))
		} else {
			returnArr = append(returnArr, int(3 * ma[number] + 1))
		}
	}
	return returnArr
}

// func (mm myMap) Encrypt() []int {
// 	returnArr := []int{}
// 	for key, value := range mm {
// 		//fmt.Println(int(key))
// 		returnArr = append(returnArr, int(key) + value)
// 	}
// 	return returnArr
// }

func Printresult(a Cipher){
	fmt.Println(a.Encrypt())
}

func main() {
	// test data

	// s1 := myString("ABCDEFGhIJKLMNOPQRrSTUVWXYZ-+")
	// a1 := myArray {1, 5, 6, 4, 8, 9, 11, 13, 14}
	// m1 := myMap {'A': 25, 'B': 40}
	// Printresult(m1)
	

	// Input
	var input string
	fmt.Printf("Enter:\n1 for string\n2 for int array\n")
	// fmt.Printf("Enter:\n1 for string\n2 for int array\n3 for map\n")
	fmt.Scanln(&input)
	switch input {
	case "1":
		var s myString
		fmt.Println("Enter the string without spaces: ")
		fmt.Scanln(&s)
		Printresult(s)
	case "2":
		fmt.Printf("Enter size of your array: ")
		var size int
		fmt.Scanln(&size)
		var arr = make([]int, size)
		for i:=0; i<size; i++ {
			fmt.Printf("Enter %dth element: ", i+1)
			fmt.Scanf("%d", &arr[i])
		}
		Printresult(myArray(arr))
	// case "3":
	// 	var mm = make(map[rune]int)
	// 	fmt.Printf("Enter size of your map: ")
	// 	var size int
	// 	fmt.Scanln(&size)
	// 	for i:=0; i<size; i++ {
	// 		var value int
	// 		var key rune
	// 		fmt.Printf("Enter %dth key value: ", i+1)
	// 		fmt.Scanf("%d %d", &key, &value)
	// 		mm[key] = value
	// 	}
	// 	Printresult(myMap(mm))
	}


}