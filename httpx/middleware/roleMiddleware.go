package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/cloud01-wu/cgsl/httpx/model"
	"github.com/cloud01-wu/cgsl/httpx/role"
	"github.com/cloud01-wu/cgsl/httpx/server"
)

func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseObject := model.Response{}
		routeID := mux.CurrentRoute(r).GetName()
		serverInstance := server.Get()

		// username forwarded by API gateway
		usernameString := r.Header.Get("X-User-Name")
		// role identity forwarded by API gateway
		roleString := r.Header.Get("X-User-Role")

		tuple, ok := serverInstance.Routes.Get(routeID)
		if !ok {
			errObject := model.Error{
				Status: http.StatusInternalServerError,
				Detail: "Could not find role definition",
			}
			responseObject.Errors = append(responseObject.Errors, errObject)
			responseFunc(w, r, responseObject)
			return
		} else {
			routeItem := tuple.(server.RouteItem)
			if len(routeItem.Roles) == 0 {
				// no roles are specified
				// pass
				next.ServeHTTP(w, r)
			} else {
				// role verification
				granted := false
				for _, allowRole := range routeItem.Roles {
					if allowRole == role.IRole(roleString) {
						granted = true
						break
					}
				}

				// deny
				if !granted {
					errObject := model.Error{
						Status: http.StatusForbidden,
						Detail: "Authorization denied",
					}
					responseObject.Errors = append(responseObject.Errors, errObject)
					responseFunc(w, r, responseObject)
					return
				}

				// pass
				// set key information to context for passing to next handler
				context.Set(r, "username", usernameString)
				context.Set(r, "role", roleString)

				next.ServeHTTP(w, r)
			}
		}
	})
}

func responseFunc(w http.ResponseWriter, _ *http.Request, resultObject model.Response) {
	result, err := json.Marshal(resultObject)
	if err != nil {
		errObject := model.Error{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}
		resultObject.Errors = append(resultObject.Errors, errObject)
		fmt.Fprintln(w, string(result))
	} else {
		fmt.Fprintln(w, string(result))
	}
}
