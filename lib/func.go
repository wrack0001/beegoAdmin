package lib

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) (md5str string) {
	has := md5.Sum([]byte(str))
	md5str = fmt.Sprintf("%x", has)
	return
}
