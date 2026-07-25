package main

import "fmt"

func chnageValue(x *int){
	*x= 1000
}

func arrayModificationWithoutPointer(arr [5]int){
	arr[0] = 1000
	fmt.Println("inside with out pointer function ",arr)
}

func arrayModificationWithPointer(arr *[5]int){
	arr[0] = 1000
	fmt.Println("inside with pointer function ",arr)
}

func main(){

	arr := [5]int{1,2,3,4,45}
	// arrayModificationWithPointer(&arr)
	arrayModificationWithoutPointer(arr)
	fmt.Println(arr)

	a := 10
	fmt.Println("a value ",a)
	b := &a
	fmt.Println("b value ", b)
	
	a = 100

	fmt.Println("after change a value ",a)
	fmt.Println("after chnage a value b value is ", *b)

	// de-referenceing in pointer 
	*b = 1000
	fmt.Println("after change a value ",a)
	fmt.Println("after chnage a value b value is ", *b)

	y := 20
	fmt.Println(y)

	chnageValue(&y)
	// after change 
	fmt.Println(y)


}
