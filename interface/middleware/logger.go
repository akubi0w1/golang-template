package middleware

import (
	"net"
	"net/http"

	"github.com/akubi0w1/golang-sample/log"
)

func AccessLog(handler http.Handler) http.Handler {
	logger := log.New()
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO-akubi: accessLogに何を載せるか
		_host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			_host = r.RemoteAddr
		}
		_method := r.Method
		_query := r.URL.RawQuery
		_path := r.URL.Path

		logger.Info("ip: %s, method: %s, query: %s, pathParam: %s",
			_host,
			_method,
			_query,
			_path,
		)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
