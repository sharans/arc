package arc

import (
	"net/http"
	"os"

	"github.com/twlabs/personal-assistant/arc/logger"
)

func Authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		platformTokenFromRequest := r.Header.Get("Authorization")
		platformTokenStoredInPlanet := getPlatformToken()

		logger.Debugf("Platform Token in the request is : %v", platformTokenFromRequest)
		logger.Debugf("platformTokenStoredInPlanet is : %v", platformTokenStoredInPlanet)

		if platformTokenFromRequest != platformTokenStoredInPlanet {

			w.WriteHeader(http.StatusForbidden)
			writeBody(w, "Invalid Token")
			return

		}
		next.ServeHTTP(w, r)
	})

}

func getPlatformToken() string {
	return os.Getenv("PLATFORM_TOKEN")
}
