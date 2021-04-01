// 实际开发中，一个错误类型通常包含两部分：Code 部分，用来唯一标识一个错误；Message 部分，用来展示错误信息，这部分错误信息通常供前端直接展示。这两部分映射在 errno 包中即为 &Errno{Code: 0, Message: "OK"}
// 自定义错误类型
package errno

var (
	// 系统错误
	// 格式：错误类型（1）- 服务模块（00）- 具体的错误码（0x）
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// 用户操作错误
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
