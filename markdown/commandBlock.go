package markdown

import (
	"strings"
	"regexp"
)

type CommandBlock struct {
	Title string
	Description string
	Command string
}

func rmLastBr(s string) (string) {
	if s[len(s)-1:] == "\n" {
		return s[0:len(s)-1]
	}
	return s
}

func rmLineHeadDollerAndPercent(s string) (string) {
	re := regexp.MustCompilePOSIX("^( )*(\\$|%)( )*")
	return re.ReplaceAllString(s, "")
}

func (c CommandBlock) CommandOneLineString() (string) {
	cmd := c.Command
	cmd = strings.ReplaceAll(c.Command, "\\\n", " ")
	cmd = rmLastBr(cmd)
	cmd = rmLineHeadDollerAndPercent(cmd)
	cmd = strings.ReplaceAll(cmd, "\n", "; ")
	return cmd
}
