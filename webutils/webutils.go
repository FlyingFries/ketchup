package webutils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/sebest/xff"
)

func MustGetPublicIP(r *http.Request) string {
	ipstr, _, err := net.SplitHostPort(xff.GetRemoteAddr(r))
	if !net.ParseIP(ipstr).IsGlobalUnicast() || err != nil {
		if os.Getenv("GO_ENV") != "dev" {
			panic("ip is not public!")
		} else {
			fmt.Println("USING GO_ENV == dev!!!!!")
		}
	}
	return ipstr
}

func IsXHR(r *http.Request) bool {
	return strings.ToLower(r.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

//RenderJSON renders JSON
func RenderJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
