package main

import (
	"log"
	"os"
	"short/cmd"
	"short/dep"
	"strconv"

	"github.com/byliuyang/app/fw"
	"github.com/joho/godotenv"
)

func main() {
	autoLoadEnv()

	host := getEnv("DB_HOST", "localhost")
	portStr := getEnv("DB_PORT", "5432")
	port := mustInt(portStr)
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "short")
	recaptchaSecret := getEnv("RECAPTCHA_SECRET", "")
	githubClientID := getEnv("GITHUB_CLIENT_ID", "")
	githubClientSecret := getEnv("GITHUB_CLIENT_SECRET", "")
	jwtSecret := getEnv("JWT_SECRET", "")
	webFrontendURL := getEnv("WEB_FRONTEND_URL", "")

	cmdFactory := dep.InjectCommandFactory()
	dbConnector := dep.InjectDBConnector()
	dbMigrationTool := dep.InjectDBMigrationTool()

	dbConfig := fw.DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}
	githubConfig := cmd.GithubConfig{
		ClientID:     githubClientID,
		ClientSecret: githubClientSecret,
	}

	rootCmd := cmd.NewRootCmd(
		dbConfig,
		recaptchaSecret,
		githubConfig,
		jwtSecret,
		webFrontendURL,
		cmdFactory,
		dbConnector,
		dbMigrationTool,
	)
	cmd.Execute(rootCmd)
}

func autoLoadEnv() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}
	return val
}

func mustInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
