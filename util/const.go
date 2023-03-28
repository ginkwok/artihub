package util

import "io/fs"

type key string

const (
	LoggerKey key = "logger"
	MySQLKey  key = "mysql"
)

const (
	PERM_OF_DIR  fs.FileMode = 0775
	PERM_OF_FILE fs.FileMode = 0644
)
