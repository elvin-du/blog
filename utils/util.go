package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	C_EXCERPT_TAG = "__EXCERPT__TAG__"
)

//template func
func Excerpt(in string) string {
	i := strings.Index(in, C_EXCERPT_TAG)
	if -1 != i {
		return in[:i]
	}
	return in
}

func YYYYMMDD(date string) string {
	feilds := strings.Fields(date)
	if len(feilds) >= 1 {
		return feilds[0]
	}
	return date
}

func UUID() string {
	gid := os.Getgid()
	uid := os.Getuid()
	pid := os.Getpid()
	t := time.Now().Unix()

	uuid := int64(gid) + int64(uid) + int64(pid) + t + rand.Int63()
	id := fmt.Sprintf("cookie:%f", uuid)

	h := md5.New()
	_, err := h.Write([]byte(id))
	if nil != err {
		beego.Error(err)
		return "abkjlj;lkjl;kjfdaslkjf;lkajsdf;kdasjfal;"
	}
	return hex.EncodeToString(h.Sum(nil))
}
