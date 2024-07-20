package main 

import "fmt"

// func main(){
// 	congrats:="Happy Birthday"
// 	fmt.Println(congrats)
// }
// func main(){
// 	penniestText:=2.0
// 	fmt.Printf("The type of penniestTest is %T\n",penniestText)
// }
// func main(){
// 	averageOpenRate,displayMessage:=.23,"is the average open rate"
// 	fmt.Println(averageOpenRate,displayMessage)
// }

// Convert the float to int

// func main(){
// 	accountAge:=2.6
// 	accountAgeInt:=int(accountAge)
// 	fmt.Println("Your account has existed for",accountAgeInt,"Years")
// }

// Constants

// func main(){
// 	const premiumPlanName = "Premium Plan"
// 	const basicPlanName = "Basic Plan"
// 	fmt.Println("plan",premiumPlanName)
// 	fmt.Println("plan",basicPlanName)

// }

// computed contants

// func main(){
// 	const secondsInMinute=60
// 	const minutesInHour = 60
// 	const secondsInHour = secondsInMinute * minutesInHour
// 	fmt.Println("No of seconds in an Hour:",secondsInHour)
// }

// formating strings in Go

// func main(){
// 	const name="Sreelakshmi"
// 	const openRate=30.5
// 	msg:=fmt.Sprintf(
// 		"Hi %s,Your open Rate is %.1f percent",
// 		name,
// 		openRate,
// 	)
// 	fmt.Println(msg)

// }


// conditionals

func main(){
	const height = 9
	if height > 6 {
		fmt.Println("You are super tall")
	}else if height > 4 {
		fmt.Println("You are tall enough")
	}else{
		fmt.Println("you are not tall enough")
	}
}