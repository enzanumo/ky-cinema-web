package dao

import (
	"sync"

	"github.com/enzanumo/ky-theater-web/internal/core"
	"github.com/enzanumo/ky-theater-web/internal/dao/jinzhu"
	"github.com/sirupsen/logrus"
)

var (
	ds                      core.DataService
	onceTs, onceDs, onceOss sync.Once
)

func DataService() core.DataService {
	onceDs.Do(func() {
		var v core.VersionInfo
		ds, v = jinzhu.NewDataService()

		logrus.Infof("use %s as data service with version %s", v.Name(), v.Version())
	})
	return ds
}
