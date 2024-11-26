package persistence

import (
	"fmt"
	"os"
	"reflect"

	"github.com/ecommerce/user/domain/contracts"
	"github.com/ecommerce/user/infrastructure/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Container struct {
	DB       *gorm.DB
	bindings map[reflect.Type]interface{}
}

func NewContainer() *Container {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_DATABASE_USERNAME"),
		os.Getenv("MYSQL_DATABASE_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Database connection failed! %v", err.Error()))
	}

	if os.Getenv("APP_ENV") == "testing" {
		var tables []string
		sqlDB.Raw("SHOW TABLES").Scan(&tables)
		for _, table := range tables {
			sqlDB.Migrator().DropTable(table)
		}
	}

	// TODO: check this more
	sqlDB.AutoMigrate(&models.User{}, &models.AccessToken{})

	return &Container{
		DB: sqlDB,
		bindings: map[reflect.Type]interface{}{
			reflect.TypeOf((*contracts.UserRepository)(nil)):        NewSqlUserRepository(sqlDB),
			reflect.TypeOf((*contracts.AccessTokenRepository)(nil)): NewSqlAccessTokenRepository(sqlDB),
		},
	}
}

func (c *Container) Resolve(interfaceType reflect.Type) (interface{}, error) {
	if implementation, ok := c.bindings[interfaceType]; ok {
		return implementation, nil
	}

	return nil, fmt.Errorf("no binding found for type %v", interfaceType)
}
