package converter

import (
	"crypto/md5"
	"fmt"
)

func StringToMD5(str string) string {
	h := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", h[:8])
}
