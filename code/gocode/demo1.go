package main

import (
	"fmt"
)

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

当切片达到了最大容量，再通过append()函数对切片进行追加的话会重新创建一个新的底层数组。
for i,v := range s {}, 这种方式中i返回的是索引，v返回的是指定位置的值的拷贝，因此对v的修改并不会改变原始值。
for i := range s {}, 这种方式返回的是索引。

切片对象与其底层数组的地址是不同的，但是切片的元素与对应的底层数组的元素是相同的（地址相同，为同一个元素）。
*/

func case1()  {
	s := []int{1, 2, 3}		//构造了一个切片，这个切片长度是3，容量也是3.
	ss := s[1:]				//ss的值是{2，3}
	fmt.Printf("%p\n",&ss)		//打印ss地址。
	fmt.Printf("%p\n",&ss[0])  //ss的第一个元素的地址。
	ss = append(ss, 4)		//此处需要特别注意，当容量满了的话，如果再对切片扩充，那么底层数组会重新分配，分配一个新的底层数组来追加新的值。
									//所以这里ss下面对应的底层数组已经不是原来的底层数组了。
	fmt.Printf("%p\n",&ss)
	fmt.Printf("%p\n",&ss[0])  //ss的第一个元素的地址。

	for _, v := range ss {
		v += 10		//此种方法返回的元素是值传递，是原始值的拷贝，因此对v复制并不会改变原始值。
	}
	fmt.Println(ss)

	for i := range ss {		//此种方法返回值的值是ss切片的索引，对切片指定位置修改，返回的是切片指定位置的元素。
		ss[i] += 10
	}
	fmt.Println(ss)
	fmt.Println(s)
}

func f() (int, int) {
	return 1,2
}

/*
如果x被声明了，y没有被声明，那么下面哪个语句是正确的。

:=的左边必须有新的符号，但是不必全部变量都是新的符号。
对于:=,左边必须至少有一个是新的符号；对于=，左边有新符号和没有新符号都可以编译通过。
*/
func case2()  {
	var x = 3	//声明x
	//x, _ := f()	//此句是错误的，因为符号 := 的左边必须有新的符号。
	x, _ = f()
	x, y := f()
	x, y = f()
	fmt.Println(x,y)
}

/*纠正行A和行B的两处错误。*/
func case3() {
	//var m map[string]int //A	此处的错误是对于m只进行了声明，但是没有进行初始化。
	var m map[string]int = make(map[string]int)
	m["a"] = 1
	//if v := m["b"]; v != nil { //B	v是值类型，不是指针类型，因此不能把nil赋值给v。
	if v:= m["b"]; v!= 0 {
		println(v)
	}
}

/*
声明全局变量，下面哪个才是对的。
var g
var g G
var g = G{}
g := G{}
*/
func case4() {
	//var g 		//错误，没有变量类型。
	//var g int		//正确。
	//var g = G{}	//正确。
	//g := G{}		//:=方式的赋值只能用在函数中的局部变量中，不能用在全局变量中。
}

//https://blog.csdn.net/pplin/article/details/70241075
func case5() {
	s := "123"
	ps := &s
	b := []byte(*ps)	//重新分配了一份内存并用*ps进行初始化，所以b与s已经是两个独立的内存了。
	pb := &b

	s += "4"	//导致底层存储重新分配内存空间。
	*ps += "5"	//导致底层存储重新分配内存空间。
	b[1] = '0'

	println(*ps)			//12345
	println(string(*pb))	//103
}

func main()  {
	fmt.Println("This is a demo program.")
	//case1()
	//case2()
	//case3()
	case5()
}
