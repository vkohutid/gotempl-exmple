package internals

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	_ = godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file, assuming this is PROD")
	// }
	return os.Getenv(key)
}
