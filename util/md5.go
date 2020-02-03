package util

import (
	"crypto/md5"
	"fmt"
)

func Md5(data []byte) (result string){
	result = fmt.Sprintf("%x", md5.Sum(data))
	return
}