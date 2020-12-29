package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

func NewLogger(cpath string) (*logrus.Logger, error) {
	config := NewConfig()
	_, err := toml.DecodeFile(cpath, &config)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	log := logrus.New()

	formatter, err := config.Formatter()
	if err != nil {
		return nil, fmt.Errorf("cannot init formatter: %v", err)
	}
	log.SetFormatter(formatter)

	level, err := config.Level()
	if err != nil {
		return nil, fmt.Errorf("cannot parse level: %v", err)
	}
	log.SetLevel(level)

	return log, nil
}
