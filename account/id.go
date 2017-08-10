package account

import (
	"fmt"

	"github.com/zheng-ji/goSnowFlake"
)

//SnowFlake算法，生成唯一id
func GenUniqID() *goSnowFlake.IdWorker {
	iw, err := goSnowFlake.NewIdWorker(machineID)
	if err != nil {
		Logger.Println(err.Error())
	}
	fmt.Println(iw)
	return iw
}

//检测id是否合法
func CheckId(id *goSnowFlake.IdWorker) bool {

}
