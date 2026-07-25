package main

import "fmt"



func makeCoffee(i int){
	fmt.Println("Making coffee .....",i)
}

func main(){
	for i := 1; i <=10; i++{

		if(i%2==0){
			continue
		}


		makeCoffee(i)

		if(i==9){
			break
		}
	}

}