package main

import (
	"fmt"
	"sync"
)

func sendMail(address string) {
	fmt.Println("Sending mail to", address)
}

func main() {
	emails := []string{
		"recipient1@example.com",
		"recipient2@example.com",
		"xyz@example.com",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(emails))

	for _, email := range emails {
		mail := email
		go func(m string) {
			sendMail(m)
			wg.Done()
		}(mail)
	}
	wg.Wait()
	fmt.Println("All emails queued for sending")
	// Do other stuff
}
