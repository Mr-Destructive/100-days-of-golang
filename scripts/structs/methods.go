package main

import "fmt"

type Mail struct {
	sender     string
	subject    string
	sent       bool
	word_count int
}

func (m Mail) check_spam() bool {
	if m.subject == "" {
		return true
	} else {
		return false
	}
}

func (m Mail) print_spam(spam bool) {
	if spam {
		fmt.Println("Spam!!")
	} else {
		fmt.Println("Safe!!")
	}
}

func main() {
	mail_one := *new(Mail)
	fmt.Printf("Mail one: ")
	is_mail_1_spam := mail_one.check_spam()
	mail_one.print_spam(is_mail_1_spam)

	mail_two := Mail{"xyz@xyz.com", "Golang Structs", true, 100}
	fmt.Printf("Mail two: ")
	is_mail_2_spam := mail_two.check_spam()
	mail_two.print_spam(is_mail_2_spam)
}
