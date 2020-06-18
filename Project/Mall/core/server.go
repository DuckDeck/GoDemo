package core

import (
	"GoDemo/Project/Mall/global"
	"GoDemo/Project/Mall/initialize"
	"fmt"
	"net/http"

	"time"
)

func RunServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf("0.0.0.0:%d", global.G_CONFIG.System.Addr)

	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	global.G_LOG.Error(s.ListenAndServe())
}
