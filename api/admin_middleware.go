package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanoohlaing1997/online-ordering-items/models"
)

type contextKey string

const authUser contextKey = "auth_user"

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has a JSON content type
		// contentType := r.Header.Get("Content-Type")
		// if !strings.Contains(contentType, "application/json") {
		// 	http.Error(w, "Invalid Content-Type, expected application/json", http.StatusUnsupportedMediaType)
		// 	return
		// }

		vars := mux.Vars(r)
		userID := vars["user_id"]
		if userID == "" {
			http.Error(w, "Missing user id in route parameter", http.StatusBadRequest)
			return
		}

		intUserID, err := StringToUint64(userID)
		if err != nil {
			http.Error(w, "Error converting string to uint64", http.StatusInternalServerError)
			return
		}

		user, err := models.NewDatabaseManager().GetUserByID(intUserID)
		if err != nil {
			http.Error(w, "User not found!!! ", http.StatusNotFound)
			return
		}

		if user.Status != 1 {
			http.Error(w, "User is not admin user", http.StatusBadRequest)
			return
		}

		// Create a new context with user information
		ctx := context.WithValue(r.Context(), authUser, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
