package main

import "fmt"

// 作用于变量上的方法实际上是不区分变量到底是指针还是值的。
// 当碰到接口类型的值，这会变得有点复杂，原因是接口变量存储的具体的值是不可寻址的，
// 使用不当，编辑器会报错。
type List []int

func (l List) Len() int  {
	return len(l)
}


func (l *List)Append(val int)  {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int)  {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool  {
	return l.Len()*10 > 42
}

func main()  {
	// a bare value  值类型变量
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//	List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)

	if LongEnough(lst) { // lst 变量，实现Lener接口， 也就是可以调用Len()
		fmt.Printf("- lst is long enough\n")
	}

	// a pointer value 指针类型变量
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("plst is long enough\n")
	}

	/* 在接口上调用方法时，必须有和方法定义是相同的接受者类型或者是可以从具体类型P直接可以辨识的：
	*	指针方法可以通过指针调用
	*	值方法可以通过值调用
	*	接受者是值的方法可以通过指针调用，因为指针会的后弦破解引用
	*	接受者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
	*
	*/

	// 将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此变量上调用，因此不正确的值在编译期就会失败

	/*  go语言规范定义接口方法集
	*	类型T的可调用方法集包含接受者为T或者T的所有方法集
	*	类型T的可调用方法集包含接受者为T的所有方法
	*	类型T的可调用方法集不包含接受者为*T的方法
	*/

}









