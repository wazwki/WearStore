package middleware

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/wazwki/WearStore/api-gateway/pkg/jwtutil"
)

type ProtectedRoute struct {
	Path   string
	Method string
}

func AuthMiddleware(jwtService *jwtutil.JWTUtil, protectedRoutes []ProtectedRoute) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, route := range protectedRoutes {
				if matchesRoute(r.URL.Path, route.Path) && r.Method == route.Method {
					token := r.Header.Get("Authorization")
					if token == "" {
						http.Error(w, "missing token", http.StatusUnauthorized)
						return
					}
					_, err := jwtService.ValidateToken(r.Context(), token, true)
					if err != nil {
						http.Error(w, "invalid token", http.StatusUnauthorized)
						return
					}
					break
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isProtectedRoute(path string, protectedPaths []string) bool {
	for _, protectedPath := range protectedPaths {
		if strings.HasPrefix(path, protectedPath) {
			return true
		}
	}
	return false
}

func matchesRoute(requestPath, routePath string) bool {
	router := mux.NewRouter()
	router.Handle(routePath, nil)
	req, _ := http.NewRequest(http.MethodGet, requestPath, nil)
	match := &mux.RouteMatch{}
	return router.Match(req, match)
}
