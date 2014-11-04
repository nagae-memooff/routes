package routes
import (
	"net/http"
)

var DefaultMux = New()

func Get(patten string, handler http.HandlerFunc) {
	DefaultMux.Get(patten, handler)
}
func Post(patten string, handler http.HandlerFunc) {
	DefaultMux.Post(patten, handler)
}
func Put(patten string, handler http.HandlerFunc) {
	DefaultMux.Put(patten, handler)
}
func Delete(patten string, handler http.HandlerFunc) {
	DefaultMux.Del(patten, handler)
}
