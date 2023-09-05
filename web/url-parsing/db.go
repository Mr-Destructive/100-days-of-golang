package main

import (
	"fmt"
	"net/url"
)

func main() {
	//postgres db url
	dbUrl, err := url.Parse("postgres://admin:pass1234@localhost:5432/mydb")
	if err != nil {
		panic(err)
	}
	fmt.Println(dbUrl)
	fmt.Println("Scheme/Protocol = ", dbUrl.Scheme)
	fmt.Println("User = ", dbUrl.User)
	//fmt.Println("User = ", dbUrl.User.String())
	fmt.Println("Username = ", dbUrl.User.Username())
	password, _ := dbUrl.User.Password()
	fmt.Println("Password = ", password)
	fmt.Println("Host = ", dbUrl.Host)
	fmt.Println("HostName = ", dbUrl.Hostname())
	fmt.Println("Port = ", dbUrl.Port())
	fmt.Println("DB Name = ", dbUrl.Path)
}
