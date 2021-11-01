package controller

import (
        "crypto/md5"
        "encoding/json"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
        "log"
        "net/http"
        "regexp"
	"main/model"
	"main/ssql"
	"strconv"
)

func registerRegisterRoutes() {
	http.HandleFunc("/", handleRegister)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// 解析请求
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	// 判断格式
	reg1, _ := regexp.MatchString(`1\d{10}`, user.Mobile)          // 手机号
	reg2, _ := regexp.MatchString(`^[^0-9][\w_]{3,11}`, user.Name) // 用户名必须是4-12位字母、数字或下划线，不能以数字开头
	reg3, _ := regexp.MatchString(`^[\w_]{6,20}`, user.Password)   // 密码必须是6-20位的字母、数字或下划线
	if !(reg1 && reg2 && reg3) {
		//格式不正确
		rst := &model.Result{
			Code: 500,
			Msg:  "输入格式不正确",
			Data: []string{},
		}
		response, _ := json.Marshal(rst)  // json化结果集
		fmt.Fprintln(w, string(response)) // 返回结果
		return
	} else {
		//加密
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
		//添加用户
		err := ssql.AddUser(&user)
		if err != nil {
			//手机号存在
			rst := &model.Result{
				Code: 500,
				Msg:  "手机号已注册",
				Data: []string{},
			}
			response, _ := json.Marshal(rst)  // json化结果集
			fmt.Fprintln(w, string(response)) // 返回结果
			return
		}
		rst := &model.Result{
			Code: 200,
			Msg:  "注册成功",
			Data: []string{strconv.Itoa(user.ID)},
		}
		response, _ := json.Marshal(rst)
		fmt.Fprintln(w, string(response))
		return
	}
}
