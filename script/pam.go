package main

import (
	"../src/pam"
	"bufio"
	"errors"
	"fmt"
	"github.com/bgentry/speakeasy"
	"log"
	"os"
)

// This example uses whatever default PAM service configuration is available
// on the system, and tries to authenticate any user. This should cause PAM
// to ask its conversation handler for a username and password, in sequence.
//
// This application will handle those requests by displaying the
// PAM-provided prompt and sending back the first line of stdin input
// it can read for each.
//
// Keep in mind that unless run as root (or setuid root), the only
// user's authentication that can succeed is that of the process owner.
func main() {
	t, err := pam.StartFunc("", "", func(s pam.Style, msg string) (string, error) {
		switch s {
		case pam.PromptEchoOff:
			//return "aaa", fmt.Errorf("bbb")
			return speakeasy.Ask(msg)
		case pam.PromptEchoOn:
			fmt.Print(msg + " ")
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return "", err
			}
			return input[:len(input)-1], nil
		case pam.ErrorMsg:
			log.Print(msg)
			return "", nil
		case pam.TextInfo:
			fmt.Println(msg)
			return "", nil
		}
		return "", errors.New("Unrecognized message style")
	})
	if err != nil {
		log.Fatalf("Start: %s", err.Error())
	}
	err = t.Authenticate(0)
	if err != nil {
		log.Fatalf("Authenticate: %s", err.Error())
	}
	fmt.Println("Authentication succeeded!")
}
