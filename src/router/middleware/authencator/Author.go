package authencator

import _ "github.com/jinzhu/gorm/dialects/mysql"

type Author struct {
	UserId   string `gorm:"Column:user_id; primary_key; Type:varchar(32)"`
	Password string `gorm:"Column:password; Type:varchar(128)"`
}

func (author Author) TableName() string {
	return "authors"
}
