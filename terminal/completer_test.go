package terminal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCompleterToString(t *testing.T) {
	c := NewCompleter()
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
