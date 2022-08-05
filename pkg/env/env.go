package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const day = 24 * time.Hour

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
		return nil, errors.Wrap(err, fmt.Sprintf("failed to convert string %q to int", s))
	}
	e.Retention = time.Duration(r*int(time.Hour)) * day

	return &e, nil
}
