package auth 

func extractSession() string {
	return "loggedinUser"
}

func GetSession() string {
	return extractSession()
}