package user

import (
	"gopractice/util"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type User struct {
	Id       int64  `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,min=1,max=20"`
	Password string `json:"password" validate:"required,min=6,max=30"`
	Age      int    `json:"age" validate:"required,gte=0,lte=130"`
	Gender   string `json:"gender" validate:"oneof=male female"`
}

func UserExist(username string, password string) bool {
	user := User{
		Username: username,
		Password: password,
	}

	db := util.OpenMySQL()
	if db == nil {
		return false
	}

	tx := db.First(&user)

	return tx.RowsAffected > 0
}

func Rgister(user User) util.Result {
	err := validate.Struct(user)
	if err != nil {
		res := util.Result{
			Code:    -1,
			Message: "valid params error",
		}
		return res
	}

	isExist := UserExist(user.Username, user.Password)

	if isExist {
		result := util.Result{
			Code:    -1,
			Message: "用户已存在！",
		}

		return result
	}

	db := util.OpenMySQL()
	if db == nil {
		result := util.Result{
			Code:    -1,
			Message: "数据库异常！",
		}

		return result
	}

	tx := db.Create(&user)
	if tx.RowsAffected > 0 {
		data := make(map[string]int64)
		data["id"] = user.Id

		result := util.Result{
			Code: 0,
			Data: data,
		}

		return result
	}

	result := util.Result{
		Code:    0,
		Message: "注册失败！",
	}

	return result
}
