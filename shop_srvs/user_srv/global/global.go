package global

import (
	"gorm.io/gorm"
	"shop_srvs/user_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig  *config.NacosConfig  = &config.NacosConfig{}
)

//func init() {
//	dsn := "root:root@tcp(192.168.0.104:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold:             time.Second, // Slow SQL threshold
//			LogLevel:                  logger.Info, // Log level
//			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
//			ParameterizedQueries:      false,       // Don't include params in the SQL log
//			Colorful:                  true,        // Disable color
//		},
//	)
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//	})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//}
