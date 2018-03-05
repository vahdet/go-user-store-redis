package interfaces

import "github.com/vahdet/go-user-store-redis/models"
type (
	UserDal interface {
		Get(id int64) (*models.User, error)
		Create(user *models.User) error
		Update(id int64, user *models.User) error
		Delete(id int64) error
		Count(id int64) (int64, error)
	}
)
