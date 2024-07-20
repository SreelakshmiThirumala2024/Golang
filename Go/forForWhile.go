package main 
import "fmt"

func getMaxMessagesToSend(costMultiplier float64,maxCostInPennies int) int{
	actualCostInPennies:=1.0
	maxMessagesToSend:=0
	for actualCostInPennies<float64(maxCostInPennies){
		maxMessagesToSend++
		actualCostInPennies*=costMultiplier
	}
	return maxMessagesToSend
}
func test(costMultiplier float64, maxCostInPennis int){
	defer fmt.Println("\n===============================\n")
	fmt.Printf("\nCost Multiplier is %v",costMultiplier)
	fmt.Printf("\nMax Cost Pennis is %v:",maxCostInPennis)
	fmt.Printf("\nMax messages can send with %v and %v is %v",costMultiplier,maxCostInPennis,getMaxMessagesToSend(costMultiplier,maxCostInPennis))
}
func main(){
	test(2.0,100)
	test(50.0,2000)
}