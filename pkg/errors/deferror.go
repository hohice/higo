package errors

// 自定义的错误类型
type DefineError struct {
	msg string
}

func (d *DefineError) Error() string {
	return d.msg
}

func NewDefineError(msg string) *DefineError {
	return &DefineError{msg}
}

/*
[ref](/usr/local/go/src/fmt/errors.go)
func Errorf(format string, a ...interface{}) error {
 p := newPrinter()
 p.wrapErrs = true
 p.doPrintf(format, a)
 s := string(p.buf)
 var err error
 if p.wrappedErr == nil {
  err = errors.New(s)
 } else {
  err = &wrapError{s, p.wrappedErr}
 }
 p.free()
 return err
}
*/

type DefWarpError struct {
	DefineError
	err error
}

func NewDefWarpError(msg string, err error) *DefWarpError {
	return &DefWarpError{DefineError{msg: msg}, err}
}

func (d *DefWarpError) Error() string {
	return d.msg
}

func (d *DefWarpError) Unwrap() error {
	return d.err
}

func (d *DefWarpError) Stack() []string {
	var i = 0
	var buff = make([]string, 4)

	buff[i] = d.msg
	derr := d.err
	for derr != nil {
		i = i + 1
		if i >= len(buff) {
			newbuff := make([]string, 2*len(buff))
			copy(newbuff, buff)
			buff = newbuff
		}
		buff[i] = derr.Error()

		if dderr, ok := derr.(*DefWarpError); ok {
			derr = dderr.Unwrap()
		} else {
			derr = nil
		}
	}

	return buff[:i+1]
}
