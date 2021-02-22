package main

import (
	"fmt"
	//"io"

	//"io"
	"os"
	//"strconv"
)

func vivod(m []int){
	for _,val := range m{
		if val == 0{
			fmt.Print("(")
		}else{
			fmt.Print(")")
		}
	}
	fmt.Println("")
}

func bal(m []int) int{

	var stack []int
	for i:= len(m)-1; i >-1;i-- {
		if m[i] == 1{
			stack =  append(stack, 1)
		}
		if m[i] == 0 && len(stack) > 0 && (stack[len(stack) -1] == 1){
			stack = stack[:len(stack) -1]
		}else{
			if m[i] == 0 && len(stack) == 0{
				return -1
			}
		}
	}
	if len(stack) == 0{
		return 0
	}else{
		return -1
	}

}
func double_arifm(m []int) []int{
	buf := 0
	for i:= len(m)-1; i >-1;i--{
		if i == len(m) -1{
			m[i] = m[i] +1
		}
		if m[i] > 1{
			buf = 1
			m[i] = 0
		}else{
			if m[i] == 1 && buf ==1{
				m[i] = 0
			}else {
				m[i] = m[i] + buf
				buf = 0
			}
		}
	}
	//fmt.Println(m)
	return m
}

func main(){
		var a int


		fmt.Fscan( os.Stdin, &a)

		balance:=0
		m := make([]int, a *2)
		for j:=0;j<a*2;j++{
			if j < a{
				m[j] = 0
			}else{
				m[j] = 1
			}
		}
		for {
			balance = bal(m)
			if balance == 0{
				vivod(m)
			}
			m = double_arifm(m)
			if m[0] == 1{
				break
			}
		}

}
