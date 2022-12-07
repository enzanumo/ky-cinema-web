// Core service implement base gorm+mysql/postgresql/sqlite3.
// Jinzhu is the primary developer of gorm so use his name as
// pakcage name as a saluter.

package dao

import (
	"github.com/Masterminds/semver/v3"
	"github.com/enzanumo/ky-theater-web/internal/conf"
	"github.com/enzanumo/ky-theater-web/internal/core"
)

var (
	_ core.DataService = (*dataServant)(nil)
)

type dataServant struct {
	core.WalletService
	core.UserManageService
}

func NewDataService() core.DataService {
	// initialize CacheIndex if needed

	db := conf.MustGormDB()

	ds := &dataServant{
		WalletService:     newWalletService(db),
		UserManageService: newUserManageService(db),
	}
	return ds
}

func (s *dataServant) Name() string {
	return "Gorm"
}

func (s *dataServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
