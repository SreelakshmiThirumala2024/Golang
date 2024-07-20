package main 
import "fmt"

func getMessageCosts(messages []string) []float64{
	costs:=make([]float64,len(messages))
	for i:=0;i<len(messages);i++{
		message:=messages[i]
		cost:=float64(len(message))*0.01
		costs[i]=cost
	}
	return costs
}
func test(messages [3]string){
	costs:=getMessageCosts(messages[:])
	fmt.Println("Messages:")
	for i:=0;i<len(messages);i++{
		fmt.Printf("-%v\n",messages[i])
	}

	fmt.Println("Costs:")
	for i:=0;i<len(costs);i++{
		fmt.Printf("-%v\n",costs[i])
	}

}
func main(){
	test([3]string{"heloo how are you","how do you dear friend","go away"})
}