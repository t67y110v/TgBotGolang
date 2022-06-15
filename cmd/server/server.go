package main

import (
	"encoding/json"
	"flag"
	. "goBot/repository/DB/mongodb"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Guid struct {
	Guid string `json:"guid"`
}

type MsgJson struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	infoLog.Printf("Соединение с БД")
	InitDB()
	r := mux.NewRouter()
	r.Handle("/api/get-token", GetTokensHandler).Methods("POST")
	r.Handle("/api/refresh-token", RefreshTokenHandler).Methods("PUT")
	r.Handle("/api/check-token", CheckTokenHandler).Methods("POST")
	addr := flag.String("addr", ":4000", "localhost")
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  r,
	}
	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

var GetTokensHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	infoLog.Printf("Запрос на получение токена")
	body, _ := io.ReadAll(r.Body)
	var guid Guid
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Host", "localhost")

	err := json.Unmarshal(body, &guid)
	if err != nil {
		errHandler(err, "Ошибка при разборе json", &w)
		return
	}

	SendTokenResponse(guid.Guid, &w, InsertRefreshToken)
	infoLog.Printf("Токен успешно сгенерирован")
})

var CheckTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Host", "localhost")
	var token Tokens
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errHandler(err, "Ошибка при чтении тела запроса", &w)
		return
	}
	err = json.Unmarshal(body, &token)
	if err != nil {
		errHandler(err, "Ошибка при разборе json", &w)
		return
	}
	_, err = ParseVerifiedAccessToken(token.Access)
	if err != nil {
		errHandler(err, "Ошибка валидации access токена", &w)
		return
	}
	w.WriteHeader(http.StatusOK)
	message, _ := json.Marshal(MsgJson{Status: 1, Msg: "Валидация прошла успешно!"})
	_, err = w.Write(message)
	infoLog.Printf("Валидация прошла успешно!")
})

var RefreshTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Host", "localhost")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errHandler(err, "Ошибка при чтении тела запроса", &w)
	}
	token, err := DecodingJsonToken(body)
	if err != nil {
		errHandler(err, "Ошибка при разборе json", &w)
		return
	}

	if token.Refresh == "" || token.Access == "" {
		errHandler(nil, "Отсутствует токен(ы)", &w)
		return
	}

	if claims, err := ParseVerifiedAccessToken(token.Access); claims == nil || err != nil {
		errHandler(err, "Ошибка валидации access токена", &w)
		return
	} else {
		if err := RefreshTokenValidate(claims.Guid, token.Refresh); err == nil {
			SendTokenResponse(claims.Guid, &w, UpdateRefreshToken)
		} else {
			errHandler(err, "Ошибка валидации refresh токена", &w)
		}
	}
})

func SendTokenResponse(guid string, w *http.ResponseWriter, query func(string, string) error) {
	if guid == "" {
		errHandler(nil, "Поле guid пустое или отсутствует", w)
		return
	}

	access, err := GetNewAccessToken(guid)
	if err != nil {
		errHandler(err, "Ошибка при генерации access токена", w)
		return
	}

	refresh, err := CreateRefreshToken(guid, query)
	if err != nil {
		errHandler(err, "Ошибка при создании refresh токена", w)
		return
	}

	response, err := TokenEncodingJson(Tokens{Status: 1, Access: access, Refresh: refresh, Guid: guid})
	(*w).WriteHeader(http.StatusCreated)
	_, err = (*w).Write(response)
	infoLog.Printf("Токен успешно сгенерирован")
}

func errHandler(err error, errText string, w *http.ResponseWriter) {
	errorLog.Println(err)
	(*w).WriteHeader(http.StatusBadRequest)
	message, _ := json.Marshal(MsgJson{Status: 0, Msg: errText})
	_, _ = (*w).Write(message)
	return
}
