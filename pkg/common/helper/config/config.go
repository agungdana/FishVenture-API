package config

type AppConfig struct {
	Name           string
	Host           string
	Port           string
	Debug          string
	FirebaseConfig FirebaseConfig
}

type DbConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type FirebaseConfig struct {
	FireBase string
}
