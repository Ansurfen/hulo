package terminal

import (
	"fmt"
	"testing"
)

func TestTerminalToString(t *testing.T) {
	term := NewHulo()
	fmt.Println(term.toString())
}
