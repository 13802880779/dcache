package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dcache/cache"

	"github.com/gorilla/mux"
)

type KV struct {
	Key   string
	Value string
}
type CacheOptAPI struct {
	ctable *cache.CacheTable
	router *mux.Router
}

func (coapi *CacheOptAPI) Get(w http.ResponseWriter, r *http.Request) { //query an item with key
	if coapi.ctable == nil {
		json.NewEncoder(w).Encode("{}")
		return
	}
	params := mux.Vars(r)
	if k, ok := params["key"]; ok {
		println("========>key:", k)
		v := coapi.ctable.Get(k)
		if v != nil {
			kval := KV{Key: k, Value: v.Value.(string)}
			println("========>", kval.Value)
			json.NewEncoder(w).Encode(kval)
			return
		}

	}
	json.NewEncoder(w).Encode("{}")
}
func Put(w http.ResponseWriter, r *http.Request) { //add or update an item

}

func (coapi *CacheOptAPI) Delete(w http.ResponseWriter, r *http.Request) {
	if coapi.ctable == nil {
		json.NewEncoder(w).Encode("{}")
		return
	}
	params := mux.Vars(r)
	if k, ok := params["key"]; ok {
		coapi.ctable.RemoveCacheItem(k)
	}
	json.NewEncoder(w).Encode("{}")
}

func (coapi *CacheOptAPI) register() {
	println("register api func")
	coapi.router.HandleFunc("/kv/{key}", coapi.Get).Methods("GET")
	//coapi.router.HandleFunc("/kv?k={key}", coapi.Get).Methods("GET")
	coapi.router.HandleFunc("/kv/{key}", coapi.Delete).Methods("DELETE")
}
