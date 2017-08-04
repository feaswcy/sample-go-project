package controller

type Controller interface {
	Deal(param map[string]string) map[string]interface{}
}

