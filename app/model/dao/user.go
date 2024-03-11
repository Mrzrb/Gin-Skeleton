package dao

type User struct {
	NickName string `gorm:"column:nickName; type:varchar(255)"`
	Email    string `gorm:"column:email; type:varchar(255)"`
	Password string `gorm:"column:password; type:varchar(255)"`
	Model
}
