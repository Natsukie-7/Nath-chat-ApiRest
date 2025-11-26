package EnvLoader

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	loader *EnvLoader
	once   sync.Once
)

type EnvLoader struct{}

func Load() *EnvLoader {
	once.Do(func() {
		_ = godotenv.Load()

		loader = &EnvLoader{}
	})

	return loader
}

func (EnvLoader) String(key string, fallback ...string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	log.Fatalf("[ENV] Variável obrigatória ausente: %s", key)
	return ""
}

func (EnvLoader) Int(key string, fallback ...int) int {
	val := os.Getenv(key)
	if val == "" {
		if len(fallback) > 0 {
			return fallback[0]
		}
		log.Printf("[ENV] Variável %s não encontrada, usando 0", key)
		return 0
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("[ENV] %s não é número válido (%s), usando fallback", key, val)
		if len(fallback) > 0 {
			return fallback[0]
		}
		return 0
	}
	return n
}

func (EnvLoader) Bool(key string, fallback ...bool) bool {
	val := strings.ToLower(strings.TrimSpace(os.Getenv(key)))

	if val == "" {
		if len(fallback) > 0 {
			return fallback[0]
		}
		return false
	}

	truthy := map[string]bool{
		"1":  true,
		"on": true,

		"t":    true,
		"true": true,

		"yes": true,
		"y":   true,
	}
	return truthy[val]
}
