package utils

// # Get 'Action User'
func GetActionUser(user string) string {
	var actionUser string
	if user != "" {
		actionUser = user
	} else {
		actionUser = "system"
	}
	return actionUser
}
