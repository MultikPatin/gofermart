package app

import "net/http"

func isAllowedMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return false
	} else {
		return true
	}
}
