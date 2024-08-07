package main

import (
	"fmt"
)


type messageToSend struct{
	message string
	sender user
	recipient user
}
type user struct{
	name string
	number int
}

func canSendMessage(mToSend messageToSend) bool{
	if mToSend.sender.name==""{
		return false
	}
	if mToSend.recipient.name==""{
		return false
	}
	if mToSend.sender.number==0{
		return false
	}
	if mToSend.recipient.number==0{
		return false
	}
	return true

}

func main() {
	fmt.Println(canSendMessage(messageToSend{
		message:"HAI",
		sender:user{name:"",number:0},
		recipient:user{name:"hai",number:0},
	}))
}