package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

// Middleware to check content type as JSON
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported media type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Middleware to add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		cookie := http.Cookie{Name: "ServerTime(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Check if method is POST
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		//Your resoure creation logic goes here. For now it is plain
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		// Tell everything is fine
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		// Say method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
	}
}
func main() {
	http.HandleFunc("/city", mainLogic)
	http.ListenAndServe(":8080", nil)
}
