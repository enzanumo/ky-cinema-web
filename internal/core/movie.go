package core

import "github.com/enzanumo/ky-theater-web/internal/model"

type MovieInfoService interface {
	GetMovieList() ([]*model.Movie, error)
}
