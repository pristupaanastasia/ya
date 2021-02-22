package main

import (
	"fmt"
	"os"
)

type graph struct {
	x int
	y int
	name int
	len int
}

func Sqrt(x int) int {
	z := 1.0
	check := 0.0
	b := float64(x)
	for i:=0; i <= 6; i++	{
		z = z - (z*z - b)/(2*z)
		if i != 0 {
			if check - z < 0.00000001 {
				break
			}
		}
		check = z
	}
	return int(z)
}

func solve(g []graph, max int, start int, end int){

	for i,_:= range g{
		g[i].len = (g[i].x + g[start -1].x)*(g[i].x + g[start -1].x)  + (g[i].y + g[start-1].y)* (g[i].y + g[start-1].y)
		g[i].len = Sqrt(g[i].len)
		if g[i].len > max{
			g[i].len = -1
		}


	}

	fmt.Println(g)
}


func main(){

	var k int
	var max int
	var start int
	var end int
	fmt.Fscan( os.Stdin, &k)
	g := make([]graph,k)
	for i:=0; i<k;i++{
		fmt.Fscan( os.Stdin, &g[i].x,&g[i].y)
		g[i].name = i + 1
	}
	fmt.Fscan( os.Stdin, &max)
	fmt.Fscan( os.Stdin, &start)
	fmt.Fscan( os.Stdin, &end)
	solve(g,max,start,end)

}