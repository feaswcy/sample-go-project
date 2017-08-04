package common

func uniqIDGener() *goSnowFlake.IdWorker {
	iw, err := goSnowFlake.NewIdWorker(machineID)
	if err != nil {
		Logger.Println(err.Error())
	}
	return iw
}
