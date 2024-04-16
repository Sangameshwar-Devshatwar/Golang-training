package folder1

func Login(usernm string, pass string) string {
	var myUsername = "Sangam@123"
	var myPass = "1234"
	if usernm == myUsername {
		if pass == myPass {
			return "logged in successfully :-welcome to GO"
		} else {
			return "Your password is wrong"
			// reveal_password()
		}
	} else {
		return "Your username is wrong"
	}
}
