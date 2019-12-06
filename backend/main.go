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
	graphQLAPIPort := mustInt(getEnv("GRAPHQL_API_PORT", "8080"))
	httpAPIPort := mustInt(getEnv("HTTP_API_PORT", "80"))

	keyGenBufferSize := mustInt(getEnv("KEY_GEN_BUFFER_SIZE", "50"))
	kgsHostname := getEnv("KEY_GEN_HOSTNAME", "localhost")
	kgsPort := mustInt(getEnv("KEY_GEN_PORT", "8080"))

	facebookClientID := getEnv("FACEBOOK_CLIENT_ID", "")
	facebookClientSecret := getEnv("FACEBOOK_CLIENT_SECRET", "")
	facebookRedirectURI := getEnv("FACEBOOK_REDIRECT_URI", "")

	googleClientID := getEnv("GOOGLE_CLIENT_ID", "")
	googleClientSecret := getEnv("GOOGLE_CLIENT_SECRET", "")
	googleRedirectURI := getEnv("GOOGLE_REDIRECT_URI", "")

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

	serviceConfig := cmd.ServiceConfig{
		RecaptchaSecret:      recaptchaSecret,
		GithubClientID:       githubClientID,
		GithubClientSecret:   githubClientSecret,
		FacebookClientID:     facebookClientID,
		FacebookClientSecret: facebookClientSecret,
		FacebookRedirectURI:  facebookRedirectURI,
		GoogleClientID:       googleClientID,
		GoogleClientSecret:   googleClientSecret,
		GoogleRedirectURI:    googleRedirectURI,
		JwtSecret:            jwtSecret,
		WebFrontendURL:       webFrontendURL,
		GraphQLAPIPort:       graphQLAPIPort,
		HTTPAPIPort:          httpAPIPort,
		KeyGenBufferSize:     keyGenBufferSize,
		KgsHostname:          kgsHostname,
		KgsPort:              kgsPort,
	}

	rootCmd := cmd.NewRootCmd(
		dbConfig,
		serviceConfig,
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
