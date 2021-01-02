package bot

import (
	"testing"

	tg "gopkg.in/tucnak/telebot.v2"
)

func Test_getEchoHandler(t *testing.T) {
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		if what != "Echo Hello" {
			t.Fatalf("Got wrong msg (%v)", what)
		}
	}
	command, callback := getEchoHandler(responseCallback)
	if command != tg.OnText {
		t.Fatalf("Got wrong command (%v)", command)
	}

	callback(&tg.Message{Text: "Hello"})
}
