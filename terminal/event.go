package terminal

type HuloEventBus struct {
	term *HuloTerminal
}

func NewHuloEventBus() *HuloEventBus {
	return &HuloEventBus{}
}

func (bus *HuloEventBus) BeforeRead() {

}

func (bus *HuloEventBus) AfterRead() {

}

func (bus *HuloEventBus) BeforeExec() {}

func (bus *HuloEventBus) AfterExec() {}

type Event interface{}

const ()
