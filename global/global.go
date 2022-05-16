package global

import (
	"douyin/common/config"

	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)
