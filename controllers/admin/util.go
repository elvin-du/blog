package admin

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()

	log.EnableFuncCallDepth(true)
	log.SetLogger("console", "")
}

var globalSessions *session.Manager

var log = logs.NewLogger(10000)

const (
	C_COOKIE_NAME = "__macs__word__"
)
