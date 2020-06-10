package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webgo/model"
)

func WebServerBase() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", loginTask)
	http.HandleFunc("/register", registerTask)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe(":8899", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func loginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...")

	//模拟延时
	//time.Sleep(2 * time.Second)

	//获取客户端通过GET/POST方式传递的参数
	_ = req.ParseForm()

	param_username, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		_, _ = fmt.Fprintln(w, "非法访问")
		return
	}

	result := model.NewBaseJsonBean()
	username := param_username[0]
	password := param_password[0]

	s := "userName:" + username + ",password:" + password
	fmt.Println(s)

	if username == "zhangsan" && password == "123456" {
		result.Code = 200
		result.Msg = "登录成功"
	} else {
		result.Code = 201
		result.Msg = "用户名或密码不正确"
	}

	bytes, _ := json.Marshal(result)
	_, _ = fmt.Fprintln(w, string(bytes))
}

func registerTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("register task running")

	result := model.NewBaseJsonBean()

	//获取客户端通过GET/POST方式传递的参数
	err := req.ParseForm()

	if err != nil {
		fmt.Println("获取参数错误")

		result.Code = 202
		result.Msg = "获取参数错误"
		bytes, _ := json.Marshal(result)
		_, _ = fmt.Fprintln(w, string(bytes))
		return
	}

	param_userAccount, found1 := req.Form["account"]
	param_userPassword, found2 := req.Form["password"]
	param_userPhone, found3 := req.Form["phone"]
	param_authCode, found4 := req.Form["authCode"]

	if !found1 {
		fmt.Println("账号不能为空")
		result.Code = 201
		result.Msg = "账号不能为空"
		bytes, _ := json.Marshal(result)
		_, _ = fmt.Fprintln(w, string(bytes))
		return
	}

	if !found2 {
		fmt.Println("密码不能为空")
		result.Code = 201
		result.Msg = "密码不能为空"
		bytes, _ := json.Marshal(result)
		_, _ = fmt.Fprintln(w, string(bytes))
		return
	}

	if !found3 {
		fmt.Println("手机号不能为空")
		result.Code = 201
		result.Msg = "手机号不能为空"
		bytes, _ := json.Marshal(result)
		_, _ = fmt.Fprintln(w, string(bytes))
		return
	}

	if !found4 {
		fmt.Println("验证码不能为空")
		result.Code = 201
		result.Msg = "验证码不能为空"
		bytes, _ := json.Marshal(result)
		_, _ = fmt.Println(w, string(bytes))
		return
	}

	account := param_userAccount[0]
	password := param_userPassword[0]
	phone := param_userPhone[0]
	code := param_authCode[0]

	if account == "" && password == "" && phone == "" && code == "" {

	}

}
