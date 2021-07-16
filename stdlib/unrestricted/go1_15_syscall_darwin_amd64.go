// Code generated by 'yaegi extract syscall'. DO NOT EDIT.

// +build go1.15,!go1.16

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall/syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Exec":         reflect.ValueOf(syscall.Exec),
		"Exit":         reflect.ValueOf(syscall.Exit),
		"ForkExec":     reflect.ValueOf(syscall.ForkExec),
		"Kill":         reflect.ValueOf(syscall.Kill),
		"PtraceAttach": reflect.ValueOf(syscall.PtraceAttach),
		"PtraceDetach": reflect.ValueOf(syscall.PtraceDetach),
		"RawSyscall":   reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":  reflect.ValueOf(syscall.RawSyscall6),
		"Shutdown":     reflect.ValueOf(syscall.Shutdown),
		"StartProcess": reflect.ValueOf(syscall.StartProcess),
		"Syscall":      reflect.ValueOf(syscall.Syscall),
		"Syscall6":     reflect.ValueOf(syscall.Syscall6),
		"Syscall9":     reflect.ValueOf(syscall.Syscall9),
	}
}
