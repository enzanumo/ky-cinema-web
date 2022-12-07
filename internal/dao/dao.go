package dao

import (
	"github.com/Masterminds/semver/v3"
	"github.com/enzanumo/ky-theater-web/internal/conf"
	"gorm.io/gorm"
	"sync"

	"github.com/enzanumo/ky-theater-web/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	ds     core.DataService
	onceDs sync.Once
)

func DataService() core.DataService {
	onceDs.Do(func() {
		ds = NewDataService()
		logrus.Infof("using data service")
	})
	return ds
}

type dataServant struct {
	db *gorm.DB
}

func NewDataService() core.DataService {
	// initialize CacheIndex if needed

	db := conf.MustGormDB()

	ds := &dataServant{db}
	return ds
}

func (s *dataServant) Name() string {
	return "Gorm"
}

func (s *dataServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
