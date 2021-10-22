package controller

import (
	"fmt"
	"net/http"
	"main/ssql"
)

func registerListRoutes() {
	http.HandleFunc("/list", handleList)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	flag, _ := ssql.IsLogin(r)
	if flag{
		fmt.Fprintf(w, "宿舍名\t楼号\t床位数\t空余床位数\n")
		dorms, _ := ssql.GetDorms()
		for _,dorm:= range dorms{
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n",dorm.Name,dorm.Number,dorm.Amount,dorm.Remain)
		}
	}else{
		fmt.Fprintf(w,"请先登录")
	}
}
