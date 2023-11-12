package config

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"log/slog"

	"github.com/caarlos0/env/v10"
)

// Config определяет конфигурацию для gophermart.
type Config struct {
	// Уровень логирования.
	Level slog.Level `env:"LOG_LEVEL"`

	// Адресс запуска сервера.
	RunAddress string `env:"RUN_ADDRESS"`

	// Строка подключения к БД.
	DatabaseURI string `env:"DATABASE_URI"`

	// Адресс системы начисления.
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`

	// Путь к файлу с секретным ключом.
	SecretKeyPath string `env:"SECRET_KEY_PATH"`

	// Время жизни токена.
	TokenTTL time.Duration `env:"TOKEN_TTL"`
}

// Clone возвращает копию конфигурации.
func (c *Config) Clone() *Config {
	c2 := *c
	return &c2
}

// Validate возвращает ошибку, если одно из полей конфигурации не валидно.
func (c *Config) Validate() error {
	if c.RunAddress == "" {
		return errors.New("the run address must not be empty")
	}
	if c.DatabaseURI == "" {
		return errors.New("the database uri must not be empty")
	}
	if c.AccrualSystemAddress == "" {
		return errors.New("the address of the accrual system must not be empty")
	}
	if c.SecretKeyPath == "" {
		return errors.New("the path of the secret key path must not be empty")
	}
	if c.TokenTTL < 0 {
		return errors.New("the token lifetime must be greater than or equal to zero")
	}
	return nil
}

// SecretKey возвращает секретный ключ, хранящийся в SecretKeyPath.
func (c *Config) SecretKey() ([]byte, error) {
	src, err := os.ReadFile(c.SecretKeyPath)
	if err != nil {
		return nil, fmt.Errorf("reading a file: %w", err)
	}

	base64 := base64.StdEncoding
	dst := make([]byte, base64.DecodedLen(len(src)))

	_, err = base64.Decode(dst, src)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding: %w", err)
	}

	return dst, nil
}

// Parse парсит переменные окружения и устанавливает их в переданную конфигурацию.
func Parse(c *Config) error {
	c2 := c.Clone()
	err := env.Parse(c)
	if err != nil {
		*c = *c2
		return fmt.Errorf("parsing env: %w", err)
	}
	return nil
}