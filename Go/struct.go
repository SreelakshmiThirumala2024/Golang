package main 

import "fmt"


type messageToSend struct{
	phoneNumber int
	message string
}
func test(m messageToSend){
	fmt.Printf("Sending Message: '%s' to %v\n",m.message,m.phoneNumber)
	fmt.Println(":::::::::::::::::::::::::::::::::::::::::::::::::::::")
}
func main() {
	test(messageToSend{phoneNumber: 1234567899,
		message:"Thanks for signing up",
	})
	test(messageToSend{phoneNumber: 9876543211,
		message:"Low Bp",
	})
	test(messageToSend{phoneNumber: 6543217899,
		message:"We are excited to have you",
	})
}