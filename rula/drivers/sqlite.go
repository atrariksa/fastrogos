package drivers

import (
	"fmt"

	"github.com/atrariksa/fastrogos/rula/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDBClient(cfg *configs.Config) *gorm.DB {
	dbCfg := cfg.DB.SQLITE
	return buildDBClient(
		dbCfg.Name,
		dbCfg.AdditionalParameters,
	)
}

func buildDBClient(name, additionalParams string) *gorm.DB {
	prepStr := "file:%v%v"
	dbPar := fmt.Sprintf(
		prepStr,
		name,
		additionalParams,
	)

	db, err := gorm.Open(sqlite.Open(dbPar), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
