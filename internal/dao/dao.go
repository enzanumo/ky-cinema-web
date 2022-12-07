package dao

import (
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
