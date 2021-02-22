package main

import (
	"fmt"
	"os"
)

type graph struct {
	x int
	y int
	name int

}


func solve(g []graph, max int, start int, end int){

	m := make([][]int,len(g))
	for j,_:= range m{
		m[j] = make([]int,len(g))
	}
	for i,_:= range g{
		for k,_ := range g {
			buf := g[i].x-g[k].x
			buf2 := g[i].y-g[k].y
			if buf < 0 {
				buf = -buf
			}
			if buf2 < 0 {
				buf2 = -buf2
			}
			m[i][k] = buf + buf2
			if m[i][k] > max || i ==k{
				m[i][k] = -1
			}
		}
	}
	if m[start-1][end -1] > 0{
		fmt.Println(1)
		return
	}
	alg := make([]int,0)
	buf := make([]int,0)
	for j,_ := range m[start-1]{
		if m[start-1][j] >0 && j != start -1 {
			alg = append(alg, j)
		}
	}
	if len(alg) ==0{
		fmt.Println(-1)
		return
	}
	min :=1
	k := -1

	for {
		k=k+1

		for j,_:=range m[alg[k]] {
			if m[j][alg[k]] > 0  && (j == end -1 || alg[k] ==end-1){
				min = min +1
				buf = append(buf, j)
				fmt.Println(min)
				return
			}
			if m[j][alg[k]] >0 && j!=alg[k] && j!=start -1 && (len(buf) ==0 ||alg[k]!= buf[len(buf) -1]){

				buf = append(buf, j)
			}
		}
		if len(buf) == 0{
			fmt.Println(-1)
			return
		}
		if k == len(alg) -1{
			alg = nil
			alg = buf
			buf = make([]int,0)
			min = min +1
			k = -1
		}
		if min > len(g) || len(alg) == 0{
			fmt.Println(-1)
			return
		}
	}

	fmt.Println(min)
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