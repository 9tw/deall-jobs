package repository

import (
	"deall/features/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{
		db: db,
	}
}

// AddUser implements domain.Repository
func (rq *repoQuery) AddUser(newUser domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.UserCore{}, err
	}

	return ToDomain(cnv), nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(userID uint) (domain.UserCore, error) {
	var resQry User

	if err := rq.db.First(&resQry, "ID = ?", userID).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}
	if err := rq.db.Unscoped().Delete(&resQry).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}

	res := ToDomain(resQry)
	return res, nil
}

// GetMyUser implements domain.Repository
func (rq *repoQuery) GetMyUser(userID uint) (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "id = ?", userID).Error; err != nil {
		log.Error("error on get my user", err.Error())
		return domain.UserCore{}, err
	}
	res := ToDomain(resQuery)
	return res, nil
}

// GetUser implements domain.Repository
func (rq *repoQuery) GetUser(existUser domain.UserCore) (domain.UserCore, error) {
	var cnv User
	if err := rq.db.Table("users").First(&cnv, "email = ?", existUser.Email).Error; err != nil {
		log.Error("error on get user login", err.Error())
		return domain.UserCore{}, nil
	}
	res := ToDomain(cnv)
	return res, nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(updatedUser domain.UserCore, userID uint) (domain.UserCore, error) {
	var cnv User = FromDomain(updatedUser)
	if err := rq.db.Table("users").Where("id = ?", userID).Updates(&cnv).Error; err != nil {
		log.Error("error on updating user", err.Error())
		return domain.UserCore{}, err
	}

	res := ToDomain(cnv)
	return res, nil
}

func (rq *repoQuery) Show() ([]domain.UserCore, error) {
	var resQry []User

	if err := rq.db.Table("users").Select("id", "fullname", "email", "role").Where("role = ?", 0).Model(&User{}).Find(&resQry).Error; err != nil {
		return nil, err
	}

	res := ToDomainArray(resQry)
	return res, nil
}
