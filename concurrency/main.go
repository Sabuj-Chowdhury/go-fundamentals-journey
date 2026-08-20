package main

import (
	"fmt"
	"time"
)

func transferFile() {
	fmt.Println("Transferring files .....")
	time.Sleep(3 * time.Second)
	fmt.Println("transfer completed")
}

func sendUrl() {
	fmt.Println("Sending url ......")
	time.Sleep(1 * time.Second)
	fmt.Println("url sent!")
}

func sendEmail() {
	fmt.Println("Sending email ...")
	time.Sleep(1 * time.Second)
	fmt.Println("Email sent!")
}

func main() {
	start := time.Now()

	transferFile()
	sendUrl()
	sendEmail()

	fmt.Println("Total time took to complete ", time.Since(start))

}
