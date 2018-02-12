package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/vahdet/go-user-store/models"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
	"github.com/vahdet/go-user-store/dal/interfaces"
)

var validate *validator.Validate

type UserService struct {
	dal interfaces.UserDal
}

func NewUserService(dal interfaces.UserDal) *UserService {
	return &UserService{dal}
}

func (s *UserService) Get(id int64) (*models.User, error) {
	return s.dal.Get(id)
}

func (s *UserService) Create(user *models.User) (*models.User, error) {
	// validation of the input
	if err := validate.Struct(user); err != nil {
		valErr := err.(validator.ValidationErrors)
		log.WithFields(log.Fields{
			"username": user.Name,
		}).Error(fmt.Sprintf("validation failed: '%#v'", valErr))
		return nil, err
	}
	// Data Access Layer call
	if err := s.dal.Create(user); err != nil {
		log.WithFields(log.Fields{
			"username": user.Name,
		}).Error(fmt.Sprintf("creation failed: '%#v'", err))
		return nil, err
	}
	return s.dal.Get(user.Id)
}

func (s *UserService) Update(id int64, user *models.User) (*models.User, error) {
	// validation of the input
	if err := validate.Struct(user); err != nil {
		valErr := err.(validator.ValidationErrors)
		log.WithFields(log.Fields{
			"username": user.Name,
		}).Error(fmt.Sprintf("validation failed: '%#v'", valErr))
		return nil, err
	}
	// Data Access Layer call
	if err := s.dal.Update(id, user); err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Error(fmt.Sprintf("Update failed: '%#v'", err))
		return nil, err
	}
	return s.dal.Get(user.Id)
}

func (s *UserService) Delete(id int64) (*models.User, error) {
	// Check if exists
	user, err := s.dal.Get(id)
	if err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Error(fmt.Sprintf("getting failed: '%#v'", err))
		return nil, err
	}
	// Data Access Layer call
	err = s.dal.Delete(id)
	return user, err
}

func (s *UserService) Count(id int64) (int64, error) {
	return s.dal.Count(id)
}
