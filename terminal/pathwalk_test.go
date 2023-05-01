package terminal

import (
	"fmt"
	"testing"
)

func TestPathWalk(t *testing.T) {
	pw := NewPathWalk(true)
	last := pw.Pwd()
	pw.Walk("..")
	pw.Walk("..")
	pw.Walk("..")
	fmt.Println(pw.Pwd())
	pw.Walk(last)
	fmt.Println(pw.Pwd())
	pwl := NewPathWalk(false, last)
	pwl.Walk("../..")
	fmt.Println(pwl.Pwd())
}
