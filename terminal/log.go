package terminal

import (
	"fmt"

	"github.com/ansurfen/cushion/utils"
	"go.uber.org/zap"
)

func init() {
	utils.InitLoggerWithDefault()
}

type HuloLogger struct {
	term *HuloTerminal
}

func NewHuloLogger() *HuloLogger {
	return &HuloLogger{}
}

func (logger *HuloLogger) Error(msg string) string {
	zap.S().Error(msg)
	return logger.term.Patiner.errColor.Render(fmt.Sprintf("[ERROR] %s", msg))
}

func (logger *HuloLogger) Info(msg string) string {
	zap.S().Info(msg)
	return msg
}

func (*HuloLogger) Fatal() {}
