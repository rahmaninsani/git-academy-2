package git_academy_2

/*
Hello should return a greeting with the name of the person
*/
func Hello(name string) string {
	return "Hello " + name
}

func Login(username string, password string) bool {
	if username != "bob" || password != "bob" {
		return false
	}

	return true
}
