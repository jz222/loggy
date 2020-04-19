package models

// Keys contains all environment variables.
type Keys struct {
	MONGO_URI     string
	MONGO_DB_NAME string
	PORT          string
	SECRET        string
	CLIENT_URL    string
	IS_SELFHOSTED bool
}
