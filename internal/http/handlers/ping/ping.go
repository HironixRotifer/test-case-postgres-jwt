package ping

import (
	"io"
	"log"
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "working as well"); err != nil {
			log.Println("service is not working")
		}
	}
}
