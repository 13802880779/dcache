package rest

import (
	"net/http"

	"github.com/dcache/cache"
	"github.com/gorilla/mux"
)

type RestApiServer struct {
	port     string
	bindAddr string
	router   *mux.Router
	ctable   *cache.CacheTable
}

func NewRestApiServer(port string, ct *cache.CacheTable) *RestApiServer {
	println("new restapiserver")
	return &RestApiServer{port: port, router: mux.NewRouter(), ctable: ct}
}
func (ras *RestApiServer) Start() {
	println("api server start at:" + ras.port)
	cacheopt := CacheOptAPI{ctable: ras.ctable, router: ras.router}
	cacheopt.register()
	http.ListenAndServe(":"+ras.port, ras.router)
}
