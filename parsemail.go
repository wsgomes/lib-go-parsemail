package parsemail

import (
	"fmt"
	"net/mail"
)

func Run() {
	addr, err := mail.ParseAddress("john.doe@")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addr)
	}
}
