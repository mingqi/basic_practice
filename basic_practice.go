package main

import (
	"fmt"
)

func main() {
	//jewlist := []string{"a", "a", "E", "A"}
	//tonelist := []string{"A", "b", "d", "c", "B", "Z", "E", "f", "a", "b", "b", "c"}
	numberlist := []int{0, 9}
	//fmt.Println(Jewfinder(jewlist, stonelist))
	//fmt.Println(Reverse(" HALLo, raLPH"))
	fmt.Println(Add(numberlist))
}

func Jewfinder(Jew []string, Stones []string) int {
	Jew_number := 0
	for _, i := range Jew {
		for _, index := range Stones {
			if index == i {
				Jew_number = Jew_number + 1
			}
		}
	}
	return Jew_number
}

func Reverse(String string) string {
	var output string
	for _, i := range []byte(String) {
		if i >= 97 && i <= 122 {
			i = i - 32
			output = output + string(i)
			continue
		} else if i >= 65 && i <= 90 {
			i = i + 32
			output = output + string(i)
			continue
		} else {
			output = output + string(i)
		}
	}
	return output
}

func Add(Intlist []int) []int {
	Intlist[len(Intlist)-1] = Intlist[len(Intlist)-1] + 1
	for i := len(Intlist) - 1; i >= 0; i-- {

		if Intlist[i] > 9 {
			Intlist[i] = 0
			if i == 0 {
				Intlist[0] = 1
				Intlist = append(Intlist, 0)
				break
			}
			Intlist[i-1] = Intlist[i-1] + 1
		}

	}
	return Intlist
}
