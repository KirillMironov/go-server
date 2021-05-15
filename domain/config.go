package domain

type Config struct {
	Port string
	Database struct {
		ConnectionString string
	}
	Security struct {
		JWTKey string
		AllowedOrigin string
	}
}
