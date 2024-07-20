package main 

import "fmt"

func maxMessages(thresh float64) int{
	totalCost:=0.0
	for i:=0.0;;i++{
		totalCost+=1.0+(0.01*i)
		if totalCost>thresh{
			return int(i)
		}
	}
	return int(totalCost)
}
func test(thresh float64){
	defer fmt.Println("\n=============================\n")
	fmt.Println("Threshold is %v:",thresh)
	max:=maxMessages(thresh)
	fmt.Printf("\nMaxium number of messages can be send %v:",max)
}
func main(){
	test(10)
	test(20)
}