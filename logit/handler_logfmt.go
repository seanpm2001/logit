package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type LogFmtHandler struct {
	BaseHandler
	Timestamp string
	Sort      bool
}

func NewLogFmtHandler() LogFmtHandler {
	return LogFmtHandler{
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
		Sort:        true,
		BaseHandler: NewBaseHandler(),
	}
}

func (config LogFmtHandler) Parse() (*Handler, error) {
	f := logrus.TextFormatter{
		DisableColors:   true,
		DisableSorting:  !config.Sort,
		FullTimestamp:   true,
		TimestampFormat: convertDateFormat(config.Timestamp),
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.formatter = &f
	return h, nil
}

func logfmtParser(meta toml.MetaData, primitive toml.Primitive) (*Handler, error) {
	fconf := NewLogFmtHandler()
	err := meta.PrimitiveDecode(primitive, &fconf)
	if err != nil {
		return nil, fmt.Errorf("parse: %v", err)
	}
	handler, err := fconf.Parse()
	if err != nil {
		return nil, fmt.Errorf("init: %v", err)
	}
	return handler, nil
}

func init() {
	RegisterParser("logfmt", logfmtParser)
}
