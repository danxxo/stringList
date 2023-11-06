package main

import (
	"fmt"
	stringList "stringList/stringList"
)

func main() {
	myList := stringList.New("WABBALABBADUBDUB")
	myList.Append('!')
	myList.Prepend('?')
	myList.Insert('1', 1)
	myList.Insert('0', 0)
	myList.Insert('9', 9)
	char, err := myList.At(2)
	fmt.Println(string(char))

	char, err = myList.At(0)
	fmt.Println(string(char))

	char, err = myList.At(9)
	fmt.Println(string(char))

	err = myList.Insert('d', 100)
	fmt.Println(err)

	myList.Remove(0)
	fmt.Println(myList.String())
	myList.Remove(myList.Len() - 1)
	char, err = myList.At(myList.Len() - 1)
	fmt.Println(string(char))

	fmt.Println(myList.String(), myList.Len())
	fmt.Println(myList.IndexOf('?'))
	fmt.Println(myList.IndexOf('9'))
	myList.Append('+')
	fmt.Println(myList.IndexOf('+'), myList.Len())
	sub, err := myList.Substring(0, myList.Len()-1)
	fmt.Println(myList.String())
	fmt.Println(sub.String())

	sub, err = myList.Substring(5, 10)
	fmt.Println(myList.String())
	fmt.Println("    ", sub.String())

	fmt.Println(myList.Replace('B', 'y', 2).String())
	fmt.Println(myList.Replace('B', '@', 0).String())

	myEmptyList := stringList.New("")
	fmt.Println(myEmptyList.At(0))
	fmt.Println(myEmptyList.String())
	myEmptyList.Insert('s', 0)
	fmt.Println(myEmptyList.String())
	myEmptyList.Remove(0)
	fmt.Println(myEmptyList.String())
	fmt.Println(myEmptyList.Equals(myList))
	fmt.Println(myEmptyList.IndexOf('p'))
}
