package interfaces

import "github.com/vahdet/tafalk-user-store/models"
type (
	UserService interface {
		Get(id int64) (*models.User, error)
		Create(user *models.User) (*models.User, error)
		Update(id int64, user *models.User) (*models.User, error)
		Delete(id int64) (*models.User, error)
		Count(id int64) (int64, error)
	}
)
