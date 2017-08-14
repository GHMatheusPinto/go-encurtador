package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/caiofilipini/encurtador/url"
)

var (
	porta int
	urlBase string
)

func init() {
	porta = 8888
	urlBase = fmt.Sprintf("http://localhost:%d", porta)
}

type Headers map[string]string

func Encurtador(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responderCom(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
			})
		return
	}
}

func main() {
	/*responsável por receber uma URL, criar um identificador 
	curto e retorna a URL curta*/
	http.HandleFunc("/api/encurtador", Encurtador)
	/*responsável por receber um identificador curto e 
	redirecionar para a URL original*/
	http.HandleFunc("/r", Redirecionador)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", porta), nil))
}