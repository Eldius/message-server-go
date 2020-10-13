package server

import (
	"fmt"
	"net/http"

	"github.com/Eldius/message-server-go/logger"
)

func Start(appPort int) {
	host := fmt.Sprintf(":%d", appPort)
	logger.Logger().Infof("starting app at '%s'", host)
	if err := http.ListenAndServe(host, CORS(RequestIdInterceptor(Routes()))); err != nil {
		logger.Logger().Panic(err.Error())
	}
}
