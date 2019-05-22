/**
 * @Author: figaro
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2019-05-21 11:02
 */

package main

import (
	"fmt"
	"know/pkg/setting"
	"know/routers"
	"net/http"
)

func main() {

	router := routers.InitRouter()


	s := &http.Server{
		Addr:fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,

	}

	s.ListenAndServe()

}
