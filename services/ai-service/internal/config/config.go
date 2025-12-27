package config

import "os"

type Config struct {
	Port        string
	JWTSecret   string
	OpenAIKey   string
	OpenAIModel string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8082"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		OpenAIKey:   os.Getenv("OPENAI_API_KEY"),
		OpenAIModel: getEnv("OPENAI_MODEL", "gpt-4o-mini"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
