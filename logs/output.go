package logs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	textReset = "\033[0m"
	textBold  = "\033[1m"
	textDim   = "\033[2m"
)

var textColor = map[string]string{
	"trace":    "\033[1;35m",
	"debug":    "\033[1;34m",
	"info":     "\033[1;36m",
	"warn":     "\033[1;33m",
	"error":    "\033[1;31m",
	"critical": "\033[1;31m",
	"panic":    "\033[1;31m",
}

func (e Entry) dumpJSON() string {
	j, err := json.Marshal(e.fields)
	if err != nil {
		Error(err.Error())
	}
	return fmt.Sprintf("%s\n", string(j))
}

func (e Entry) dumpText() string {
	text := ""
	if t, ok := e.fields["time"].(time.Time); ok {
		text += fmt.Sprintf("%s[%s]%s", textDim, t.Format(time.RFC3339), textReset)
	}
	if level, ok := e.fields["level"]; ok {
		text += fmt.Sprintf("[%s%s%s] ", textColor[level.(string)], strings.ToUpper(level.(string)), textReset)
	}
	if msg, ok := e.fields["msg"]; ok {
		text += fmt.Sprintf("%s%s%s ", textBold, msg, textReset)
	}
	for _, k := range e.keys {
		if k != "time" && k != "level" && k != "msg" {
			text += fmt.Sprintf("%s[%s=\"%v\"]%s", textDim, k, e.fields[k], textReset)
		}
	}
	text = strings.TrimSpace(text)
	return fmt.Sprintf("%s\n", text)
}
