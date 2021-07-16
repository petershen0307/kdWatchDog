package handlers

import (
	"testing"
)

func Test_getEchoHandler(t *testing.T) {
	mailbox := make(chan Mail, 10)
	h := NewHandler(mailbox, nil)
	h.echo(&Mail{fromMsg: "Hello"})
	mail := <-h.mailbox
	if mail.toMsg != "Echo Hello" {
		t.Fatalf("Got wrong msg (%v)", mail.toMsg)
	}
}
