package settings

import "os"

type Database struct {
	URL string
}

type Application struct {
	Port string
	Env  string
}

type Settings struct {
	Database    Database
	Application Application
}

var settings *Settings

func New() *Settings {
	if settings != nil {
		return settings
	}

	settings = &Settings{
		Database: Database{
			URL: os.Getenv("DATABASE_URL"),
		},
		Application: Application{
			Port: os.Getenv("PORT"),
			Env:  os.Getenv("ENV"),
		},
	}

	return settings
}

func Get() *Settings {
	return New()
}
