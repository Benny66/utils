package function

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算md5
func CreateMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
