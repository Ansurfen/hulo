package terminal

type Signal struct {
	Op uint8
}

const (
	LJMP = iota
	SJMP
	CLI
	STI
	NOP
)
