package middlew

import (
	"net/http"
)

/* es el middlew que me permite conocer el esta de la base de datos */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}
