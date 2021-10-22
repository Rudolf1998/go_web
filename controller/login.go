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
	"main/utils"
	"main/ssql"
)

func registerLoginRoutes() {
	http.HandleFunc("/login", handleLogin)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
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
	reg1, _ := regexp.MatchString(`^1\d{10}`, user.Mobile)        // 手机号
	reg3, _ := regexp.MatchString(`^[\w_]{6,20}`, user.Password) // 密码必须是6-20位的字母、数字或下划线
	if !(reg1 && reg3) {
		//格式有误
		rst := model.Result{
			Code: 500,
			Msg:  "登陆失败，手机号或密码不正确1",
			Data: []string{},
		}
		response, _ := json.Marshal(rst)  // json化结果集
		fmt.Fprintln(w, string(response)) // 返回结果
		return
	} else {
		//格式正确，查询手机号和密码是否匹配
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
		user, err := ssql.CheckUserMobileAndPassword(user.Mobile, user.Password)
		if err != nil {
			//手机号或密码不正确
			rst := model.Result{
				Code: 500,
				Msg:  "登陆失败，手机号或密码不正确",
				Data: []string{},
			}
			response, _ := json.Marshal(rst)  // json化结果集
			fmt.Fprintln(w, string(response)) // 返回结果
			return
		} else {
			//用户名和密码正确
			//生成UUID作为Session的id
			uuid := utils.CreateUUID()
			//创建一个Session
			sess := &model.Session{
				SessionID: uuid,
				UserID:    user.ID,
			}
			//将Session保存到数据库中
			ssql.AddSession(sess)
			//创建一个Cookie，让它与Session相关联
			cookie := http.Cookie{
				Name:  "dorm_user",
				Value: uuid,
				//HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			//返回成功消息
			rst := model.Result{
				Code: 200,
				Msg:  "登陆成功",
				Data: []string{},
			}
			response, _ := json.Marshal(rst)  // json化结果集
			fmt.Fprintln(w, string(response)) // 返回结果
			return
		}
	}
}
