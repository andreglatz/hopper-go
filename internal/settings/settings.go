package settings

import "os"

type Database struct {
	URL string
}

type Settings struct {
	Database Database
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
	}

	return settings
}

func Get() *Settings {
	return New()
}
