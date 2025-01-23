package gerror

import (
	"gofly/utils/tools/command"
	"gofly/utils/tools/gcode"
)

// IIs is the interface for Is feature.
type IIs interface {
	Error() string
	Is(target error) bool
}

// IEqual is the interface for Equal feature.
type IEqual interface {
	Error() string
	Equal(target error) bool
}

// ICode is the interface for Code feature.
type ICode interface {
	Error() string
	Code() gcode.Code
}

// IStack is the interface for Stack feature.
type IStack interface {
	Error() string
	Stack() string
}

// ICause is the interface for Cause feature.
type ICause interface {
	Error() string
	Cause() error
}

// ICurrent is the interface for Current feature.
type ICurrent interface {
	Error() string
	Current() error
}

// IUnwrap is the interface for Unwrap feature.
type IUnwrap interface {
	Error() string
	Unwrap() error
}

const (
	// commandEnvKeyForBrief is the command environment name for switch key for brief error stack.
	commandEnvKeyForBrief = "gf.gerror.brief"

	// commaSeparatorSpace is the comma separator with space.
	commaSeparatorSpace = ", "
)

var (
	// isUsingBriefStack is the switch key for brief error stack.
	isUsingBriefStack bool
)

func init() {
	value := command.GetOptWithEnv(commandEnvKeyForBrief)
	if value == "1" || value == "true" {
		isUsingBriefStack = true
	}
}
