package main

import (
	"fmt"
	"os"
)
type graf struct{
	x int
	y int
}
func main(){
	var a int
	var mas []graf
	var min int
	var start int
	var end int

	fmt.Fscan( os.Stdin, &a)

	mas = make([]graf, a)
	for i:=0;i<a;i++{
		fmt.Fscanln( os.Stdin, &mas[i].x, &mas[i].y)
	}
	fmt.Fscan( os.Stdin, &min)
	fmt.Fscan( os.Stdin, &start)
	fmt.Fscan( os.Stdin, &end)
}
