package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr      string
	MongoURI        string
	EthRPCURL       string
	ContractAddress string
	JWTSecret       string
	OCRAPIKey       string
}

// Load reads configuration from environment variables or .env file
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables directly")
	}

	cfg := &Config{
		ServerAddr:      getEnv("SERVER_ADDR", ":8080"),
		MongoURI:        getEnv("MONGO_URI", "mongodb://localhost:27017"),
		EthRPCURL:       getEnv("ETH_RPC_URL", ""),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
		JWTSecret:       getEnv("JWT_SECRET", ""),
		OCRAPIKey:       getEnv("OCR_API_KEY", ""),
	}

	// Validate required fields
	if cfg.EthRPCURL == "" || cfg.ContractAddress == "" {
		log.Fatalf("Missing required environment variables: ETH_RPC_URL or CONTRACT_ADDRESS")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
