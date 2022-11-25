# EvalGOMohamed

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func helloHandler1(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())
		fmt.Fprintf(w, "> %s", time.Now().Format("15:04"))
	}
}

func helloHandler2(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())
		fmt.Fprintf(w, "> %d", rand.Intn(1000))
	}
}

func helloHandler3(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())
		fmt.Fprintf(w, "> ")
		p := rand.Perm(100)
		for _, r := range p[:15] {
			if p := rand.Perm(100); p[0] < 10 {
				fmt.Fprintf(w, "%03d ", r)
			} else {
				fmt.Fprintf(w, "%d ", r)
			}
		}

	}

}

// func helloHandler4(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodPost:
// 		rand.Seed(time.Now().UnixNano())
// 		fmt.Scanf("%s", req.Body)
// 		fmt.Fprintf(w, "> %s", req.Body)

// }

func main() {
	http.HandleFunc("/", helloHandler1)
	http.HandleFunc("/dice", helloHandler2)
	http.HandleFunc("/dices", helloHandler3)
	// http.HandleFunc("/randomize-words", helloHandler4)
	http.ListenAndServe(":4567", nil)
}
