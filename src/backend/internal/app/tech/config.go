package tech

import (
	"fmt"
	"time"
)

type Postgres struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type Logger struct {
	Encoding string
	Level    string
	File     string
}

type Config struct {
	PG Postgres

	Cost       int
	TokenExp   time.Duration
	DailyBonus int
	SecretKey  string
	Span       time.Duration

	Logger Logger
}

func (d *Postgres) toDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host,
		d.Port, d.User, d.Password, d.DBName)
}
