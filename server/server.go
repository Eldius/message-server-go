package server

import (
	"fmt"
	"net/http"

	"github.com/Eldius/cors-interceptor-go/cors"
	"github.com/eldius/message-server-go/logger"
)

func Start(appPort int) {
	host := fmt.Sprintf(":%d", appPort)
	logger.Logger().Infof("starting app at '%s'", host)
	if err := http.ListenAndServe(host, cors.CORS(RequestIdInterceptor(Routes()))); err != nil {
		logger.Logger().Panic(err.Error())
	}
}
