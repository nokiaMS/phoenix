package main

import "fmt"

/*
如果把包名改成 package gocode,则编译运行的时候会报错： runnerw.exe: CreateProcess failed with error 216:
把包名改成package main就没有这个错误了。
*/

/*
go笔试面试题：https://goquiz.github.io/
*/

/*
切片的类型是 []T,数组有固定的长度，而切片没有固定的长度。
可以通过make函数创建一个切片。 func make([]T, len, cap) []T,例如：
	s := make([]byte, 5, 10)  或者  s:= make([]byte, 5) (容量默认与len相同，为5.)
如何获取切片的长度和容量：
	len(s)获取长度，cap(s)获取容量。
*/

func case1()  {
	s := []int{1, 2, 3}		//构造了一个切片，这个切片长度是3，容量也是3.
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}

func main()  {
	fmt.Println("This is a demo program.")
	case1()
}
