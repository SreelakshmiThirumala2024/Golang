package main  
import "fmt"

func printOdds(){
	for i:=0;i<10;i++{
		if i%2==0{
			continue
		}else{
			fmt.Printf("\nOdd is %v",i)
		}
	}
}


func main(){
	printOdds()
	// printPrimes(100)
}