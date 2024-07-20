package main 

import (
	"fmt"
	"errors"
)
func divide(x,y float64) (float64,error){
	if y==0{
		return 0.0,errors.New("CAN'T DIVIDE BY ZERO")
	}
	return x/y,nil
}
func main(){
	v,err:=divide(3,7)
	fmt.Printf("the divide value is %f and is there is an error %v:",v,err)
}