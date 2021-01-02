package bot

import (
	"testing"
)

func Test_echo(t *testing.T) {
	got := echo("hello")
	if got != "Echo hello" {
		t.Fatalf("Got: %s", got)
	}
}
