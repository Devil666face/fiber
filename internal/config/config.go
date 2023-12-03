package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	https = "https"
)

type Config struct {
	IP                string `env:"IP" env-default:"127.0.0.1"`
	PortHTTP          uint   `env:"PORT_HTTP" env-default:"8000"`
	PortHTTPS         uint   `env:"PORT_HTTPS" env-default:"4443"`
	SqliteDB          string `env:"SQLITE_DB" env-default:"db.sqlite3"`
	AllowHost         string `env:"ALLOW_HOST" env-default:"localhost"`
	UseTLS            bool   `env:"TLS" env-default:"false"`
	TLSKey            string `env:"TLS_KEY" env-default:"server.key"`
	TLSCrt            string `env:"TLS_CERT" env-default:"server.crt"`
	PostgresUse       bool   `env:"POSTGRES" env-default:"false"`
	PostgresHost      string `env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort      string `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresDB        string `env:"POSTGRES_DB" env-default:"db"`
	PostgresUser      string `env:"POSTGRES_USER" env-default:"superuser"`
	PostgresPassword  string `env:"POSTGRES_PASSWORD" env-default:"Qwerty123!"`
	Superuser         string `env:"SUPERUSER" env-default:"superuser@local.lan"`
	SuperuserPassword string `env:"SUPERUSER_PASSWORD" env-default:"Qwerty123!"`
	MaxQueryPerMinute int    `env:"MAX_QUERY_PER_MINUTE" env-default:"50"`
	CookieKey         string `env:"COOKIE_KEY" env-default:"VtsTmAz5I7LUM3N2NA4J7eX1XC/gNzA8DUK1Ocssowo="`
	// Use `openssl rand -base64 32` for get CookieKey
	ConnectHTTP   string
	ConnectHTTPS  string
	HTTPSRedirect string
}

func New() *Config {
	cfg := Config{}
	// if err := cleanenv.ReadEnv(&cfg); err != nil {
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			slog.Error(fmt.Sprintf("Env variable not found: %s", err))
			//nolint:revive //If not env's not set - close app
			os.Exit(1)
		}
	}
	cfg.ConnectHTTP = fmt.Sprintf("%v:%v", cfg.IP, cfg.PortHTTP)
	cfg.ConnectHTTPS = fmt.Sprintf("%v:%v", cfg.IP, cfg.PortHTTPS)
	cfg.HTTPSRedirect = fmt.Sprintf("%s://%s:%d", https, cfg.AllowHost, cfg.PortHTTPS)
	return &cfg
}
