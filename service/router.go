package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RespData struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func writeResp(status bool, msg string) []byte {
	respData := RespData{
		Status: status,
		Msg:    msg,
	}
	resp, err := json.Marshal(respData)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}
	{
		fmt.Println(r.Form.Get("username"))
	}
	if strings.ToLower(r.Form.Get("username")) != "admin" || r.Form.Get("password") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		w.Write(writeResp(false, "Error when logging in"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(writeResp(true, "Succeed to login"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(writeResp(false, "Succeed to logout"))
}
