package main  
import "fmt"
func bulkSend(numMessages int) (float64){
	sum:=0.0
	for i:=0.0;i<float64(numMessages);i++ {
		sum+=1.0+(0.01*i)
	}
	return sum
}
func test(num int){
	defer fmt.Println("\n::::::::::::::::::::::::::::::::::::::::\n")
	fmt.Printf("Sending %v messages:",num)
	cost:=bulkSend(num)
	fmt.Printf("\nCost of sending %v is %v",num,cost)
}
func main(){
	test(5)
	test(10)
}