package logger

//nolint:gochecknoglobals
var ModuleDirectory []string

func init() {
	ModuleDirectory = []string{"iGoStyle/"}
	UnitTestMode()
}
