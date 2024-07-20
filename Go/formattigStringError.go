package main    
import "fmt"

func getSMSErrorString(cost float64,recipient string) string{
	return fmt.Sprintf("SMS that cost $%v to be sent to %s:",cost,recipient)
}
func test(cost float64,recipient string){
	s:=getSMSErrorString(cost,recipient)
	fmt.Println(s)
	fmt.Println("\n=============================")
}
func main(){
	test(1.4,"+1 (435) 555 0923")
	test(2.1,"+2 (702) 555 3452")
}