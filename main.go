package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdmin bool   `json:"is_admin"`
}

func main() {
	http.HandleFunc("/", HomeController)

	fmt.Println("Server is running now!")

	http.ListenAndServe(":8000", nil)
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	var isAdmin bool

	if os.Getenv("IS_ADMIN") == "true" {
		isAdmin = true
	} else {
		isAdmin = false
	}

	age, _ := strconv.Atoi(os.Getenv("AGE"))
	user := User{
		Name:    os.Getenv("NAME"),
		Age:     age,
		IsAdmin: isAdmin,
	}

	resp := Response{
		Status:  true,
		Message: "Success!",
		Data:    user,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
