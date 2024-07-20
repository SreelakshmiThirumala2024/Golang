package main 

import "fmt"

type Truck struct{
	bedsize int
	car
}
type car struct{
	model string
	make string
}
// func main(){
	// lanesTruck:=Truck{
	// 	bedsize:10,
	// 	car:car{
	// 		model:"Camry",
	// 		make:"Toyota",
	// 	},
	// }
// 	fmt.Println(lanesTruck.model)
// }

type sender struct{
	rateLimit int
	userna
}
type userna struct{
	name string
	number int
	
}

func test(s sender){
	fmt.Println("Sender Name:",s.name)
	fmt.Println("Sender Number:",s.number)
	fmt.Println("Sender RateLimit:",s.rateLimit)
	fmt.Println("_____________________________________")

}
func main(){
	test(sender{
		rateLimit:10000,
		userna: userna{
			name:"Sree",
			number:1234567890,
		},
	})
	test(sender{
		rateLimit:20000,
		userna: userna{
			name:"Hanna",
			number:10987654321,
		},
	})
	lanesTruck:=Truck{
		bedsize:10,
		car:car{
			model:"Camry",
			make:"Toyota",
		},
		
	}
	fmt.Println(lanesTruck.model)

}