package terminal

const (
	FailDailDaemon = "fail to dial daemon. see detail, input hulo log"

	// Command
	NotFoundCmd    = "not found command"
	Exit           = "exit"
	ClearInterrupt = "clear interrupt"
	SetInterrupt   = "set interrupt"
	Restart        = "restart"

	// Registry
	FailLoadComplete = "fail to load complete"
	FailLoadLoader   = "fail to load loader"
	FailLoadLib      = "fail to load lib"

	UusupportPlatform = "support current platform"

	NoException = ""
)

// Exception is the unpacking of error, which is used to fit in lua enviroment.
// e.g. err.Error() is a kind of Exception type
type Exception string
