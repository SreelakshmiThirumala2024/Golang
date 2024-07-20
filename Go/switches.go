package main    
import "fmt" 

func printTypeNum(num interface{}){
	switch v:=num.(type){
		case int:
			fmt.Printf("%T\n:",v)
		case string:
			fmt.Printf("%T\n:",v)
		default: 
			fmt.Printf("%T\n:",v)
		
	}

}
func main(){
	printTypeNum(1)
	printTypeNum("1")
	printTypeNum(struct{}{})
}