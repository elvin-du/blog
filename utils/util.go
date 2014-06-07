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
	C_EXCERPT_TAG = "_excerpt_tag_"
)

func ExcerptContent(in string) (excerpt, content string) {
	i := strings.Index(in, C_EXCERPT_TAG)

	if -1 != i {
		return in[:i], in[i+len(C_EXCERPT_TAG):]
	}
	return in, ""
}

func YYYYMMDD(date string) string {
	feilds := strings.Fields(date)
	if len(feilds) >= 1 {
		return feilds[0]
	}
	return date
}

func YMDHMCN(data interface{}) string {
	date, ok := data.(string)
	if !ok {
		beego.Error("cannot convert interface{} to string")
		return ""
	}

	feilds := strings.Fields(date)
	ymd := strings.Split(feilds[0], "-")
	hms := strings.Split(feilds[1], ":")

	return fmt.Sprintf("%s 年 %s 月 %s 日 %s 时 %s 分", ymd[0], ymd[1], ymd[2], hms[0], hms[1])
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

func Rand(min, max int64) int64 {
	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Int63n(max)
		if r < min {
			rand.Seed(time.Now().UnixNano())
			r := rand.Int63n(max)
			if r >= min {
				return r
			}
		}
	}
}
