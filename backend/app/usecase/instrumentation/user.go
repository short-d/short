package instrumentation

import "github.com/short-d/short/backend/app/entity"

func getUserID(user *entity.User) string {
	if user == nil {
		return "anonymous"
	}
	return user.ID
}
