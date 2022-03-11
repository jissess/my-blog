package models

type User struct {
	Model
	Name     string `gorm:"size:32;unique_index" json:"name"`
	Email    string `gorm:"size:50;unique_index" json:"email"`
	Password string `json:"-"`
	Avatar   string `json:"avatar"`
	Role     int    `gorm:"default:1" json:"role"` //0-代表管理员，1-代表正常用户
}

func QueryUserByEmailAndPwd(email, pwd string) (user User, err error) {
	return user, db.Where("email = ? and password = ?", email, pwd).Take(&user).Error
}

func QueryUserByName(name string) (user User, err error) {
	return user, db.Where("name = ?", name).Take(&user).Error
}

func QueryUserByEmail(email string) (user User, err error) {
	return user, db.Where("email = ?", email).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Save(user).Error
}
