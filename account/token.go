package account

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

//根据时间戳和给定参数生成Token
func genToken(phone string, role string) string {
	crutime := time.Now().Unix()

	h := md5.New()
	str := fmt.Sprintf("%s%s%s", strconv.FormatInt(crutime, 10), phone, role)
	io.WriteString(h, str)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
