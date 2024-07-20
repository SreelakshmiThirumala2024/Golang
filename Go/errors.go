package main   
import "fmt"

func sendSMSToCouple(msgToCustomer,msgToSpouse string) (float64,error){
	costForCustomer,err:=sendSMS(msgToCustomer)
	if err!=nil{
		return 0.0,err
	}
	costForSpouse,err:=sendSMS(msgToSpouse)
	if err!=nil{
		return 0.0 ,err
	}
	return costForCustomer+costForSpouse,nil
}
func sendSMS(message string)(float64, error){
	const maxTextLen=25
	const costPerChar=.0002
	if len(message) > maxTextLen{
		return 0.0,fmt.Errorf("can't send text over %v characters",maxTextLen)
	}
	return costPerChar *  float64(len(message)),nil
}
func test(msgToCustomer, msgToSpouse string){
	defer fmt.Println("\n==========================================")
	fmt.Println("Message for the customer:",msgToCustomer)
	fmt.Println("Message for the spouse:",msgToSpouse)
	totalcost,err:=sendSMSToCouple(msgToCustomer,msgToSpouse)
	if err!=nil{
		fmt.Println("Error",err)
	}
	fmt.Printf("Total cost for sending to a couple %f:",totalcost)
	
}
func main(){
	test("I love you","But i hate you")
}