package cuslog

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"
)

type JsonFormatter struct {
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {
	if !e.logger.opt.disableApp {
		e.Map["appName"] = e.logger.opt.appName
		e.Map["appType"] = e.logger.opt.appType
	}

	if e.logger.opt.flag != 0 {
		e.Map["flag"] = e.logger.opt.flag
	}

	e.Map["level"] = LevelNameMapping[e.Level]
	e.Map["time"] = e.Time.Format(time.DateTime)
	if e.File != "" && !f.IgnoreBasicFields {
		e.Map["file"] = e.File + ":" + strconv.Itoa(e.Line)
		e.Map["func"] = e.Func
	}

	switch e.Format {
	case FmtEmptySeparate:
		e.Map["msg"] = fmt.Sprint(e.Args...)
	default:
		e.Map["msg"] = fmt.Sprintf(e.Format, e.Args...)
	}

	return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
}
