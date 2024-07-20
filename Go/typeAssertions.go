package main  
import "fmt"

func getExpenseReport(e expense) (string ,float64){
	em,ok:=e.(email)
	if ok{
		return em.toAddress,em.cost()
	}
	s,ok:=e.(sms)
	if ok{
		return s.toPhoneNumber,s.cost()
	}
	return "",0.0
}
func (e email) cost() float64{
	if !e.isSubscribed{
		return 0.05*float64(len(e.body))
	}
	return 0.01*float64(len(e.body))

}
func (s sms) cost() float64{
	if !s.isSubscribed{
		return 0.1*float64(len(s.body))
	}
	return 0.03*float64(len(s.body))

}
func (e email) print(){
	fmt.Println(e.body)
}
func (s sms) print(){
	fmt.Println(s.body)
}
type expense interface{
	cost() float64
}
type printer interface{
	print()
}
func print(p printer){
	p.print()
}
type email struct{
	isSubscribed bool
	body string
	toAddress string
}
type sms struct{
	isSubscribed bool
	body string
	toPhoneNumber string
}
func testType(e expense){
	address,cost:=getExpenseReport(e)
	switch e.(type){
	case email:
		fmt.Printf("Report : The email going to %s will cost: %.2f",address,cost)
		fmt.Println("_---------------------------------_")
	}
}
func test(e expense, p printer){
	fmt.Printf("Printing with cost; $%.2f.....\n",e.cost())
	p.print()
	fmt.Println("------------------------------------")
}

func main(){
	e:=email{
		isSubscribed: true,
		body:"Hello there!",
		toAddress:"abc@gmail.com",
	}
	test(e,e)
	testType(e)
	s:=sms{
		isSubscribed: false,
		body:"Where are you",
		toPhoneNumber: "123456789",
	}
	test(s,s)
}