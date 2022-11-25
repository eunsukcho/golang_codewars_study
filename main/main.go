package main

import (
	"log"
	"strings"
)

func main() {
	//ReverseWords("row row row your boat")
	//Multiple3And5(20)
	//SquaresInRect(20, 14)
	//QueueTime([]int{2, 2, 3, 3, 4, 4}, 2)
	//SpinWords("Hey fellow war
	//DigitalRoot(202223262356162950)
	log.Println(CamelCase("test case"))
}

func CamelCase(s string) string {
	str_arr := strings.Fields(s)

	for index, str := range str_arr {
		str_arr[index] = strings.Title(str)
	}
	return strings.Join(str_arr, "")
}

func DigitalRoot(n int) int {
	sum := 0

	for n > 0 {
		sum = sum + (n % 10)
		n = n / 10

		if n == 0 {
			sum = (sum / 10) + (sum % 10)
			if sum/10 > 0 {
				n = sum
				sum = 0
			}
		}
	}

	return sum
}

func SpinWords(str string) string {
	str_arr := strings.Fields(str)
	var retrun_arr []string

	for _, str := range str_arr {
		if len(str) >= 5 {
			retrun_arr = append(retrun_arr, ReverseWordFiveMoreLetter(str))
		} else {
			retrun_arr = append(retrun_arr, str)
		}
	}

	return strings.Join(retrun_arr, " ")
}

func ReverseWordFiveMoreLetter(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func QueueTime(customers []int, n int) int {
	//q := make(chan int, 2)

	return 0
}

func SquaresInRect(lng int, wdth int) []int {
	var return_int = []int{}

	if lng != wdth {

		for {

		}
	}

	log.Println(return_int)

	return return_int
}

func Multiple3And5(number int) int {
	var return_number int
	for i := 1; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			return_number = return_number + i
		} else if i%3 == 0 || i%5 == 0 {
			return_number = return_number + i
		}

	}

	return return_number
}

/*
func InArray(array1 []string, array2 []string) []string {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	var r = []string{"arp", "live", "strong"}

	return array2
}
*/
func ReverseWords(str string) string {
	str_arr := strings.Split(str, " ")

	var reverse_arr []string
	for i := len(str_arr) - 1; i >= 0; i-- {
		reverse_arr = append(reverse_arr, str_arr[i])
	}

	return strings.Join(reverse_arr, " ") // reverse those words
}

func GetCount(str string) (count int) {
	// Enter solution here
	for _, var_str := range strings.ToLower(str) {
		if var_str >= 97 || var_str <= 122 {
			switch var_str {
			case 'a', 'e', 'i', 'o', 'u':
				count++
			default:

			}
		}
	}

	return count
}
