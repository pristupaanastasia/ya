package main

import (
	"fmt"
	"os"
)

func eq(m map[byte]int,d map[byte]int) bool{
	for key, _ := range m{
		if m[key] != d[key]{
			return false
		}
	}
	for key, _ := range d{
		if m[key] != d[key]{
			return false
		}
	}
	return true
}

func main(){
	var a string
	var b string


	fmt.Fscan( os.Stdin, &a)
	fmt.Fscan( os.Stdin, &b)

	m := make(map[byte]int)
	for i, _ := range a{
		m[a[i]] = m[a[i]] + 1
	}
	d := make(map[byte]int)
	for j, _ := range b{
		d[b[j]] = d[b[j]] + 1
	}
	if eq(m,d){
		fmt.Println(1)
	}else{
		fmt.Println(0)
	}
}
