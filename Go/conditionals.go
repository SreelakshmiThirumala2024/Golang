package main 

import "fmt"

// func main() {
// 	messageLen := 10
// 	maxMessageLen := 20
// 	fmt.Println("Trying to send a message of length",messageLen)
// 	if messageLen > maxMessageLen {
// 		fmt.Println("Message is too long")
// 	}else {
// 		fmt.Println("Message of correct Size")
// 	}
// }

// func main(){
// 	if length:=getLength(email);length < 1 {
// 		fmt.Println("Email is invalid")
// 	}
// }


// functions
func yearsUnitEvents(age int)(yearsUnitAdult int,yearsUnitDrinking int ,yearsUnitCarRental int){
	yearsUnitAdult= 18 - age
	if yearsUnitAdult<0{
		yearsUnitAdult = 0
	}
	yearsUnitDrinking=21-age
	if yearsUnitDrinking<0{
		yearsUnitDrinking=0
	}
	yearsUnitCarRental=25-age
	if yearsUnitCarRental<0{
		yearsUnitCarRental=0
	}
	return
}
func increment(x int) int {
	x++
	return x
}

func concatenate(s1 string, s2 string) string{
	return s1 + s2
}
func add(x,y int) int{
	return x + y
}
func getNames()(s1 string,s2 string) {
	return "Sumesh", "Ramesh"
}
func main(){
	a:=5
	fmt.Println(concatenate("Hello"," World"))
	fmt.Println(concatenate("Sreelakshmi"," Thirumala"))
	x:=1
	y:=2
	fmt.Println(add(x,y))
	a=increment(a)
	fmt.Println(a)
	x1,_:=getNames()
	fmt.Println(x1)
	x2,x3,x4:=yearsUnitEvents(18)
	fmt.Println(x2,x3,x4 )

}


