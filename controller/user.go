package controller
  
import (
    "net/http"
    "text/template" //导入模版包
)

func registerUserRoutes() {
        http.HandleFunc("/user", handleUser)
}

func handleUser(w http.ResponseWriter, r *http.Request) {

    // t := template.New("index")

    // t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")

    t, _ := template.ParseFiles("./templates/index.html")
    data := map[string]string{
        // "name":    "zeta",
        // "someStr": "这是一个开始",
    }

    t.Execute(w, data)
}

