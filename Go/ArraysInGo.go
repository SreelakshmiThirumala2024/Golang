package main  

import "fmt"

func getMessagesWithRetries() [3]string{
	return [3]string{
		"Hello",
		"You are cute",
		"goodbye",
	}
}
func send(name string,doneAt int){
	fmt.Printf("Sending message to %v...",name)
	fmt.Println()
	messages:=getMessagesWithRetries()
	for i:=0;i<len(messages);i++{
		msg:=messages[i]
		fmt.Printf("Sending message %v:",msg)
		fmt.Println()
	}

}
func main(){
	send("John",10)
}