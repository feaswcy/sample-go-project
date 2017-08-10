package account

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

//根据时间,手机号,角色生成token
func GenToken(phone string, role string) string {
	crutime := time.Now().Unix()

	h := md5.New()
	str := fmt.Sprintf("%s%s%s", strconv.FormatInt(crutime, 10), phone, role)
	io.WriteString(h, str)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

//检测token是否合法
func CheckToken(token string) bool {

}
