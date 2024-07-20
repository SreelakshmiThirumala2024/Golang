package main   
import "fmt"

func fizzbuzz(){
	for i:=0 ;i<=100;i++{
		if i%3==0 && i%5==0{
			fmt.Println("FizzBuzz multiple of 3 and 5 :")
			fmt.Printf("%v",i);
			fmt.Println("\n==============================\n")
		}else if i%3==0 || i%5==0{
			if i%3==0{
				fmt.Println("FizzBuzz multiple of 3:")
				fmt.Printf("%v",i);
				fmt.Println("\n==============================\n")
			}else{
				fmt.Println("FizzBuzz multiple of 5:")
				fmt.Printf("%v",i);
				fmt.Println("\n==============================\n")
			}
		}
	}
}
func main(){
	fizzbuzz()
}