package main

import "fmt"

func main(){
	age := 15

	if(age<18){
		fmt.Println("Too young")
	}else if(age >=18 && age<=30){
		fmt.Println("Adult!")
	}else if (age>30 && age<60) {
		println("Working")		
	}else{
		println("Old")
	}

}
