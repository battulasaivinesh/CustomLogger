package constants

import "github.com/battulasaivinesh/CustomLog/common/types"

var LogTypes = struct {
	Warning types.LogType
	Error   types.LogType
	Info    types.LogType
}{
	"Warn",
	"Error",
	"Info",
}

var OutputTypes = struct {
	Console types.OutputType
	File    types.OutputType
}{
	0,
	1,
}
