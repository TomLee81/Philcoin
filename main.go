package main

import (
	"fmt"
	"net/http"

	"github.com/eunko/yourproject/backend/database"
	"github.com/gorilla/mux"
)

func main() {
	// MongoDB 연결
	database.ConnectMongo("mongodb://localhost:27017")

	// 라우터 초기화
	r := mux.NewRouter()

	// 테스트 라우트
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ 서버가 실행 중입니다!"))
	}).Methods("GET")

	// 서버 실행
	fmt.Println("✅ 서버 실행: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
