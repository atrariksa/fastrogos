package repos

import (
	"encoding/json"

	"github.com/atrariksa/fastrogos/rula/errs"
	"github.com/atrariksa/fastrogos/rula/models"
	"gorm.io/gorm"
)

const (
	DuplicateKey string = "1062"
)

type IUserRepo interface {
	IUserRepoRead
	IUserRepoWrite
}

type UserRepo struct {
	DB *gorm.DB
	ICache
}

type IUserRepoWrite interface {
	CreateUser(newUser models.User) (err error)
	UpdateUser(user models.User) (err error)
	DeleteUser(user models.User) (err error)
}

type IUserRepoRead interface {
	GetUserByUsername(user *models.User) (err error)
}

func (ur *UserRepo) CreateUser(newUser models.User) (err error) {
	tx := ur.DB.Debug().Begin()

	err = tx.Create(&newUser).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return
}

func (ur *UserRepo) UpdateUser(user models.User) (err error) {
	tx := ur.DB.Debug().Begin()

	updateUser := tx.Model(&user).Updates(&user)
	err = updateUser.Error
	if err != nil {
		tx.Rollback()
		return
	}

	if updateUser.RowsAffected == 0 {
		tx.Rollback()
		return errs.ErrUserNotFound
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	ur.Del(user.Username)
	return
}

func (ur *UserRepo) DeleteUser(user models.User) (err error) {
	tx := ur.DB.Debug().Begin()

	deleteUser := tx.Debug().Model(&user).Delete(&user)
	err = deleteUser.Error
	if err != nil {
		tx.Rollback()
		return
	}

	if deleteUser.RowsAffected == 0 {
		tx.Rollback()
		return errs.ErrUserNotFound
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	ur.Del(user.Username)
	return
}

func (ur *UserRepo) GetUserByUsername(user *models.User) (err error) {

	bUser, err := ur.Get(user.Username)
	if err != nil {
		tx := ur.DB.Debug()
		err = tx.Where("username = ?", user.Username).First(user).Error
		if err == gorm.ErrRecordNotFound {
			err = errs.ErrUserNotFound
			return
		}

		if err != nil {
			return
		}

		bUser, err = json.Marshal(&user)
		if err != nil {
			return err
		}

		ur.Set(user.Username, bUser)
		return nil
	}

	if bUser != nil {
		err = json.Unmarshal(bUser, &user)
		if err != nil {
			return err
		}
	}

	return
}
