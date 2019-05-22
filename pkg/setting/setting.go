/**
 * @Author: figaro
 * @Description: 项目设置
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2019-05-22 09:16
 */

package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JWTSecret string

)

func init()  {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini':%v",err)
	}

	LoadBase()
	LoadServer()
	LoadAPP()

}

func LoadBase()  {
	RunMode = Cfg.Section("").Key("RunMode").MustString("debug")
}

func LoadServer()  {
	sec,err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v",err)
	}

	RunMode = Cfg.Section("").Key("RunMode").MustString("debug")

	HTTPPort = sec.Key("HTTP_Port").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_Timeout").MustInt(60))*time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_Timeout").MustInt(60))*time.Second
}

func LoadAPP()  {
	sec,err := Cfg.GetSection("app")

	if err != nil {

		log.Fatalf("Fail to get section 'app' : %v",err)

	}

	JWTSecret = sec.Key("JWT_Secret").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_Size").MustInt(10)
}