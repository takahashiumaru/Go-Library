package app

import (
	"log"
	"os"
	"time"

	"go-library/helper"
	"go-library/model/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	// "api_ski:mdnfjkt45@tcp(103.103.192.24:4000)/ski?parseTime=true"
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// RUN before_auto_migrate.sql
	helper.RunSQLFromFile(database, "app/database/before_auto_migrate.sql")

	err = database.AutoMigrate(

		// GOLANG
		&domain.Book{},
		&domain.Category{},
		&domain.Publisher{},
		&domain.User{},
		&domain.Session{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	// RUN after_auto_migrate.sql
	helper.RunSQLFromFile(database, "app/database/after_auto_migrate.sql")

	// Delete Constraint
	database.Exec("ALTER TABLE books DROP FOREIGN KEY fk_books_category")
	database.Exec("ALTER TABLE books DROP FOREIGN KEY fk_books_publisher")

	// Add Constraint
	database.Exec("ALTER TABLE books ADD CONSTRAINT fk_books_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT ON UPDATE RESTRICT")
	database.Exec("ALTER TABLE books ADD CONSTRAINT fk_books_publisher FOREIGN KEY (publisher_id) REFERENCES publishers(id) ON DELETE RESTRICT ON UPDATE RESTRICT")

	return database
}
