package api

import (
	"net/http"
	"github.com/gorilla/mux"
)

// requestRouter is a wrapper for the router. Change this if you'd like to use a different router such as httprouter or roll your own net/http router.
type requestRouter struct {
	*mux.Router
}

// routerInitFunction exists so we can change the type of router if we'd like, we just need to change the implementation of this function rather than the things that rely on the router (which is presumably more).
func (r *requestRouter) routerInit() {
	r.Router = mux.NewRouter()
}

// getParams again exists so that if we ever decide to switch routers, this is an example of how we can easily change implementation but not have to change functionality
func (r *requestRouter) getParams(req *http.Request) map[string]string {
	return mux.Vars(req)
}
