package buz

import (
	"net/http"
	"github.com/pkg/errors"
	"github.com/KelvinChen684/Go-000/Week01/data"
)

func authenticate(writer http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		errors.Wrap(err, "authenticate: request parse form error")	// 标准库函数调用出错，使用Wrap包装堆栈信息
		writer.WriteHeader(400)
	}
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		errors.WithMessage(err, "authenticate: no such user")	// UserByEmail中已Wrap包装err，此处只包装上下文信息
		writer.WriteHeader(401)
	}
	writer.Header().Set("name", user)
	writer.WriteHeader(200)
}
