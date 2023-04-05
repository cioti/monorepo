package errors

import (
	"bytes"
	"fmt"
	"runtime"
)

const maxStackDepth = 50

type Error struct {
	err    error
	msg    string
	stack  []uintptr
	frames []StackFrame
}

func New(e interface{}) *Error {
	var err error

	switch e := e.(type) {
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}

	stack := make([]uintptr, maxStackDepth)
	length := runtime.Callers(2, stack[:])
	return &Error{
		err:   err,
		stack: stack[:length],
	}
}

func Wrapf(err error, format string, args ...interface{}) *Error {
	if err == nil {
		return nil
	}

	stack := make([]uintptr, maxStackDepth)
	length := runtime.Callers(2, stack[:])
	return &Error{
		err:   err,
		msg:   fmt.Sprintf(format, args...),
		stack: stack[:length],
	}
}

// Error returns the underlying error's message.
func (e *Error) Error() string {
	if len(e.msg) == 0 {
		return e.err.Error()
	}

	return e.msg + ": " + e.err.Error()
}

// Stack returns the callstack formatted the same way that go does
// in runtime/debug.Stack()
func (err *Error) Stack() []byte {
	buf := bytes.Buffer{}

	for _, frame := range err.StackFrames() {
		buf.WriteString(frame.String())
	}

	return buf.Bytes()
}

// StackFrames returns an array of frames containing information about the
// stack.
func (err *Error) StackFrames() []StackFrame {
	if err.frames == nil {
		err.frames = make([]StackFrame, len(err.stack))

		for i, pc := range err.stack {
			err.frames[i] = NewStackFrame(pc)
		}
	}

	return err.frames
}

// Return the wrapped error (implements api for As function).
func (err *Error) Unwrap() error {
	return err.err
}
