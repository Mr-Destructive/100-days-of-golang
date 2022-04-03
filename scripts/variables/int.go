package main 

import "fmt"

func main() {

    // by default a int type is dependent on computer architecture
    // it can be either int64 or int32 
    var likes int = 140
    fmt.Println(likes)
    fmt.Printf("type = %T\n",likes)

    //The below code will give error as limit overflow for int8
    //var age int8 = 140
    //fmt.Println("Age = ",age) 
}
