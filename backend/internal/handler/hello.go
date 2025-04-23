package handler

import (
	"fmt"
	"net/http"
	// "log"
)



func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API básica funcionando!")
}
