package auth

import "net/http"

func VerifyAuthentication(w http.ResponseWriter, r *http.Request, auth *Authorization) {
	auth.IsVerified = true
}
