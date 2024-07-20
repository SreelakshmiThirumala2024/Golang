package main 
import "fmt"

type employee interface{
	getName() string
	getSalary() int
}

type constructor struct{
	name string
	hourlyPay int
	hoursPerYear int
}

func (c constructor) getName() string{
	return c.name
}

func (c constructor) getSalary() int{
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct{
	name string
	salary int
}
func (f fullTime) getName() string{
	return f.name
}
func (f fullTime) getSalary() int{
	return f.salary
}
func test(e employee){
	fmt.Println(e.getName(),e.getSalary())
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}

func main(){
	test(constructor{
		name:"John",
		hourlyPay: 100,
		hoursPerYear: 2000,
	})
	test(fullTime{
		name:"Sree",
		salary:100000,
	})
}