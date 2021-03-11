package errors

import (
	"fmt"
	"os"
	"runtime/debug"
)

var PANICERROR = NewDefineError("panic errors occur\n")

func DefaultPanic() {
	/*
		不透明的错误处理不推荐
		err := errors.New("panic errors occur\n")
		if err != nil {
		 fmt.Println(err)
		  return
		}
	*/
	panic(PANICERROR)
}

func Panic(err error) {
	/*
		不透明的错误处理不推荐
		err := errors.New("panic errors occur\n")
		if err != nil {
		 fmt.Println(err)
		  return
		}
	*/
	panic(err)
}

var PanicHandler = func() {
	var ifwarp = false
	if r := recover(); r != nil {
		var err error
		switch r := r.(type) {
		case *DefWarpError:
			err = r
			ifwarp = true
		case *DefineError:
			err = r
		case error:
			err = r
		default:
			err = fmt.Errorf("%+v", r)
		}
		// stack := make([]byte, 1024)
		// enable stack all
		// length := runtime.Stack(stack, true)
		if ifwarp {
			fmt.Fprintln(os.Stderr, "PANIC RECOVER:")
			printWarpStack(r)
		} else {
			fmt.Fprintln(os.Stderr, "PANIC RECOVER:", err)
			//fmt.Printf("[%s] %s\n", "PANIC RECOVER", err)
		}
		// equal when enable stack all
		debug.PrintStack()
	}
}

func printWarpStack(ierr interface{}) {
	if p, ok := ierr.(*DefWarpError); ok {
		for _, msg := range p.Stack() {
			fmt.Fprintln(os.Stderr, msg)
		}
	}
}
