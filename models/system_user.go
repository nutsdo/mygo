package models

import (
	"fmt"
	)

type SystemUsers struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(30) notnull unique 'name'" json:"name"`
	Password string `xorm:"varchar(255) notnull default('')" json:"password"`
	Avatar string `xorm:"varchar(255) notnull default('')" json:"avatar"`
	Email string `xorm:"varchar(255) notnull default('')" json:"email"`
	IsEnabled bool `xorm:"Bool notnull default(0)" json:"is_enabled"`
	CreatedAt string `xorm:"DateTime notnull" json:"created_at"`
	UpdatedAt string `xorm:"DateTime notnull" json:"updated_at"`
}

func (m *SystemUsers) GetById(Id int64) *SystemUsers {

	systemUser := new(SystemUsers)

	_, err := DB.Id(Id).Get(systemUser)

	if err != nil {
		fmt.Println(err)
	}

	return systemUser
}

func (m *SystemUsers) GetAll() []*SystemUsers {

	SystemUsers := []*SystemUsers{}

	return SystemUsers
}