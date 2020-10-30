package main

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strings"
)

func main(){
	var cadenas []string
	var n int64
	var i int64
	var auxliarString string

	i = 0

	fmt.Print("Escriba cuantos strings quiere escribir: ")
	fmt.Scanln(&n)

	for i < n {
		fmt.Print("Escriba un string: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		auxliarString = scanner.Text()
		cadenas = append(cadenas,auxliarString)
		i++
	}

	sort.Strings(cadenas)

	file, err := os.Create("asecendente.txt")

	if err != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString(strings.Join(cadenas,"\n"))
	
	sort.Sort(sort.Reverse(sort.StringSlice(cadenas)))

	file2, err2 := os.Create("descendente.txt")

	if err2 != nil{
		fmt.Println(err2)
		return
	}

	defer file2.Close()

	file2.WriteString(strings.Join(cadenas,"\n"))

}