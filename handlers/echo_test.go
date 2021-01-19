package handlers

import (
	"testing"

	tg "gopkg.in/tucnak/telebot.v2"
)

func Test_getEchoHandler(t *testing.T) {
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	command, callback := getEchoHandler(responseCallback)
	if command != tg.OnText {
		t.Fatalf("Got wrong command (%v)", command)
	}

	callback(&tg.Message{Text: "Hello"})
	if gotValue != "Echo Hello" {
		t.Fatalf("Got wrong msg (%v)", gotValue)
	}
}
