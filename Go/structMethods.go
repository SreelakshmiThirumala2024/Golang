package main 
import "fmt"

type authenticationInfo struct{
	username string
	password string
}
func (a authenticationInfo) getBasicAuth() string{
	return fmt.Sprintf("Authorization: Basic %s:%s",a.username,a.password)
}

func test(authentic authenticationInfo){
	fmt.Println(authentic.getBasicAuth())
	fmt.Println(".....................................")
}
func main(){
	a:=authenticationInfo{
		username:"admin",
		password:"1234",
	}
	test(a)
}
