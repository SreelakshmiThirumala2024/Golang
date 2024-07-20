package main 

import (
	"fmt"
	"time"
)

func sendMessage(msg message){
	fmt.Println(msg.getMessage())
}
type message interface{
	getMessage() string
}

type birthdayMessage struct{
	birthdayTime time.Time 
	recipientName string
}
func (bm birthdayMessage) getMessage() string{
	return fmt.Sprintf("Hi %s,it is your birthday on %s",bm.recipientName,bm.birthdayTime)
}

type sendingReport struct{
	reportName string
	numberOfReports int
}
func (sr sendingReport) getMessage() string{
	return fmt.Sprintf("Your %s report is ready. Send a copy of %d reports",sr.reportName,sr.numberOfReports)
}
func test(m message){
	sendMessage(m)
	fmt.Println("=================================================")
}
func main(){
	test(birthdayMessage{
		birthdayTime:time.Date(1994,03,21,10,40,0,0,time.UTC),
		recipientName: "Sree",
	})
	test(sendingReport{
		reportName:"Monthly Report",
		numberOfReports:10,
	})
}