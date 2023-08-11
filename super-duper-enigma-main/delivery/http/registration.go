package http

import (
	"encoding/json"
	"fmt"
	"forummodule/service/users"
	"forummodule/sqllite"
	"gorm.io/gorm"
	"net/http"
)

//запуск через горутину

func Handlereg(db gorm.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "reg.html")
	})

	http.HandleFunc("/handleClick", loginstrangefunc(db))
	http.HandleFunc("/handleregclick", regstrangefunc(db))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	fmt.Println("Server is starting")
}

func loginstrangefunc(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LoginJSONinput(w, r, db)
	}
}

func LoginJSONinput(w http.ResponseWriter, r *http.Request, db gorm.DB) {
	// Check the request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON packet from the request body
	var data sqllite.LoginInput

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var response = map[string]interface{}{}
	//ЛОГИН+++
	login, err := users.ServiceLogin(db, data)
	if err != nil {
		//переадресация на страницу с упс через вызов функции
		fmt.Println("err=", err)
		response = map[string]interface{}{
			"user": "not found",
		}
	} else {
		//переадресация на главную страницу через вызов функции
		response = map[string]interface{}{
			"user": "found",
		}
	}
	fmt.Println("login=", login)
	// Send a response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func regstrangefunc(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RegJSONInput(w, r, db)
	}
}

func RegJSONInput(w http.ResponseWriter, r *http.Request, db gorm.DB) {
	// Check the request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON packet from the request body
	var data sqllite.LoginInput

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	var response = map[string]interface{}{}
	//РЕГИСТРАЦИЯ+++
	users.ServiceRegistrationUser(db, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
