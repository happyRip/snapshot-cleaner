package env

import (
	"os"
	"strconv"
	"time"
)

type Environment struct {
	Username   string
	Password   string
	Repository string
	Retention  time.Duration
}

func New() (*Environment, error) {
	e := Environment{
		Username:   os.Getenv("CLEANER_USERNAME"),
		Password:   os.Getenv("CLEANER_PASSWORD"),
		Repository: os.Getenv("CLEANER_TARGET"),
	}

	s := os.Getenv("CLEANER_RETENTION")
	r, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	e.Retention = time.Duration(r * int(time.Hour))

	return &e, nil
}
