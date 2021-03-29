// 对自定义的错误进行处理

package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

// 新建自定义的错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// 用来展示更多信息
func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

// 用来展示更多信息
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// 定义 Err结构体类型的 Error 方法（有了该方法就相当于实现了自定义错误）
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error:%s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// 解析自定义的错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}
