package reqHandler

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
