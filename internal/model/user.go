package model

import "gorm.io/gorm"

type User struct {
	*Model
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Status   int    `json:"status"`
	Balance  int64  `json:"balance"`
	IsAdmin  bool   `json:"is_admin"`
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	return u, err
}

func (u *User) Update(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ?", u.Model.ID, 0).Save(u).Error
}

func (u *User) Get(db *gorm.DB) (*User, error) {
	var user User
	if u.Model != nil && u.Model.ID > 0 {
		db = db.Where("id= ?", u.Model.ID, 0)
	} else if u.Phone != "" {
		db = db.Where("phone = ?", u.Phone, 0)
	} else {
		db = db.Where("username = ?", u.Username, 0)
	}

	err := db.First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (u *User) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*User, error) {
	var users []*User
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
