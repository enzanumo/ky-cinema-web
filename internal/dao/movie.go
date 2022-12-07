package dao

import "github.com/enzanumo/ky-theater-web/internal/model"

func (s *dataServant) GetMovieList() (movies []*model.Movie, err error) {
	tx := s.db.Find(&movies)
	err = tx.Error
	return
}
