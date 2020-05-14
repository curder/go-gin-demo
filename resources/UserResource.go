package resources

import "github.com/curder/go-gin-demo/models"

type UsersResponse struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserResource(user models.Users) UsersResponse {
	return UsersResponse{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
