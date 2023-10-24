package database

import (
	"fmt"
	"time"

	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/bagashiz/gommerce/internal/pkg/database/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Mysql is a wrapper for GORM mysql driver
type Mysql struct {
	Conn *gorm.DB
	cfg  *config.Database
}

// newMysql creates a new GORM mysql driver instance
func newMysql(cfg *config.Database) (DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// sets the maximum amount of time a connection may be reused.
	maxLifeTime := time.Duration(cfg.MaxLifeTime) * time.Second
	sqlDB.SetConnMaxLifetime(maxLifeTime)

	return &Mysql{
		db,
		cfg,
	}, nil
}

// Migrate runs the auto migration for the database.
func (m *Mysql) Migrate() error {
	return m.Conn.AutoMigrate(
		dao.User{},
		dao.Address{},
		dao.Shop{},
		dao.Category{},
		dao.Product{},
		dao.ProductPhoto{},
		dao.ProductLog{},
		dao.Transaction{},
		dao.TransactionDetail{},
	)
}

// DB returns the underlying database connection.
func (m *Mysql) DB() *gorm.DB {
	return m.Conn
}

// Close closes the database connection.
func (m *Mysql) Close() error {
	sqlDB, err := m.Conn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
