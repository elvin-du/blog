package controllers

import "github.com/astaxie/beego/session"

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}

var globalSessions *session.Manager

const (
	C_HTTP_OK             = 200
	C_HTTP_BAD_REQUEST    = 400
	C_HTTP_UNAUTHORIZED   = 401
	C_HTTP_NOT_FOUND      = 404
	C_HTTP_INTERNAL_ERROR = 500
)

const (
	C_COOKIE_NAME = "__macs__word__"
)
