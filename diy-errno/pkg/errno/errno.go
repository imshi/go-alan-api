// 自定义错误类型的详细定义及处理
package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

type Err struct {
	Code    int
	Message string
	Err     error
}

// 定义 Errno 结构体类型的 Error 方法（有了该方法就相当于实现了自定义错误） - 用户端展示
func (err Errno) Error() string {
	return err.Message
}

// 定义 Err结构体类型的 Error 方法（有了该方法就相当于实现了自定义错误） - 服务端详细日志
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error:%s", err.Code, err.Message, err.Err)
}

// 用来展示更多信息 - string
func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

// 用来展示更多信息 - Stringf：接收任意类型的可变参数进行格式化输出
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// 【核心函数 1】：新建自定义的错误，返回一个 Err 类型的指针
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// 【核心函数 2】：解析自定义的错误，返回两个值：错误状态码 和 错误内容的字符串
func DecodeErr(err error) (int, string) {
	// error为 nil返回自定义的 OK状态码和 OK信息
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

// 错误类型判断
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}
