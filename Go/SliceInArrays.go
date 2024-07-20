package main
import (
	"fmt"
	"errors"
)

const(
	planFree="free"
	planPro="pro"
)
func getMessagesWithRetriesForPlan(plan string)([]string,error){
	allMessages:=getMessagesWithRetries()
	if plan==planPro{
		return allMessages[:],nil
	}
	if plan==planFree{
		return allMessages[0:2],nil
	}
	return nil,errors.New("Unsupported plan")
}
func getMessagesWithRetries() [3]string{
	return [3]string{
		"Hello",
		"You are cute",
		"goodbye",
	}
}
func send(name string,doneAt int,plan string){
	fmt.Printf("Sending message to %v...",name)
	fmt.Println()
	messages,err:=getMessagesWithRetriesForPlan(plan)
	if err==nil{
		for i:=0;i<len(messages);i++{
			msg:=messages[i]
			fmt.Printf("Sending message %v:",msg)
			fmt.Println()
		}
	}else{
		fmt.Println("Error:",err)
	}
	

}
func main(){
	send("John",10,planFree)
	send("Beena",100,planPro)
	send("Binu",10000,"hai")
}
