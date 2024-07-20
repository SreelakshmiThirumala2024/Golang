package main 

import "fmt"

type divideError struct{
	dividend float64
}
// user defined error
func (de divideError) Error() string{
	return fmt.Sprintf("can not divide %f by zero",de.dividend)
}
func divide(dividend,divisor float64) (float64,error){
	if divisor==0{
		return 0,divideError{dividend:dividend}
	}else{
		return dividend/divisor,nil
	}
}
func main(){
	fmt.Println(divide(34,2))
}