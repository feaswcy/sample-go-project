package account

import(
	"log"
)

//注册接口
func register(phone string, role string, password)([]map[string]interface{},err){

}

//判断是否注册
func isRegister(phone string, role string) (bool, []map[string]interface{}, error) {

	users, err := queryUser(phone, role)

	var flag bool
	user := make(map[string]interface{})

	if err != nil {
		flag = false
	} else if len(users) <= 0 {
		flag = false
	} else {
		flag = true
		user = users[0]
	}

	return flag, user, err
}
