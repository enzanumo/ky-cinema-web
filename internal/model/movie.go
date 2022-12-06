package model

import "gorm.io/gorm"

type Movie struct {
	*Model
	Title    string
	Duration int
	ImgUrl   string
	Detail   string
}

func (m *Movie) Create(db *gorm.DB) (*Movie, error) {
	err := db.Create(&m).Error
	return m, err
}

func (m *Movie) Update(db *gorm.DB) error {
	return db.Model(&Movie{}).Where("id = ?", m.Model.ID, 0).Save(m).Error
}

func (m *Movie) Get(db *gorm.DB) (*Movie, error) {
	var movie Movie
	if m.Model != nil && m.Model.ID > 0 {
		db = db.Where("id= ?", m.Model.ID, 0)
	}
	err := db.First(&movie).Error
	if err != nil {
		return &movie, err
	}
	return &movie, nil
}

func (m *Movie) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*Movie, error) {
	var movies []*Movie
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
	if err = db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}
