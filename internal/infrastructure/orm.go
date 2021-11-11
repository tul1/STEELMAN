package infrastructure

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/sirupsen/logrus"
)

type ORM interface {
	Conn() *gorm.DB
	Close() error
}

type PostgreSQL struct {
	conn *gorm.DB
	log  *logrus.Logger
}

func NewPostgreSQL(DSN string, log *logrus.Logger) (*PostgreSQL, error) {
	var err error

	c := &PostgreSQL{log: log}
	c.conn, err = gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger: gormlogger.New(
			log,
			gormlogger.Config{
				SlowThreshold: 200 * time.Millisecond, // Slow SQL threshold
				LogLevel:      gormlogger.Warn,        // Log level
				Colorful:      false,                  // Disable color
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := c.conn.DB()
	if err != nil {
		return nil, err
	}

	// It's the default, 0 means no idle which means it will almost close the socket after each request.
	sqlDB.SetMaxIdleConns(2)
	// Will autoclose socket after idle, available since Go 1.15.
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)
	// Before Go1.15 way, it is the absolute lifetime of a connection.
	// So even if it is active, it will be closed. Don't think it is needed now but as Go1.15-Go1.15.2 is bugged,
	// it is mandatory for now, else it will ignore SetConnMaxIdleTime.
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	return c, nil
}

func (c *PostgreSQL) Close() {
	panic("Unimplemented method in GORM")
}

func (c *PostgreSQL) Conn() *gorm.DB {
	return c.conn
}
