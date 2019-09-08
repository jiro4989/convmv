package errorcode

type ErrorCode int

const (
	OK ErrorCode = iota
	FailedBinding
	CouldNotReadFile
	IllegalMonitorConfig
	CouldNotReadDir
)
