package migration

import (
	"github.com/1Panel-dev/1Panel/agent/global"
	"github.com/1Panel-dev/1Panel/agent/init/migration/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Init() {
	m := gormigrate.New(global.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.AddTable,
		migrations.InitHost,
		migrations.InitSetting,
		migrations.InitBackupAccount,
		migrations.InitImageRepo,
		migrations.InitDefaultGroup,
		migrations.InitDefaultCA,
		migrations.InitPHPExtensions,
		migrations.AddTask,
	})
	if err := m.Migrate(); err != nil {
		global.LOG.Error(err)
		panic(err)
	}
	global.LOG.Info("Migration run successfully")
}
