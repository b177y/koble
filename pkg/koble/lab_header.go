package koble

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

var blue = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var magBold = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
var mag = color.New(color.FgHiMagenta).SprintFunc()
var MAXPRINTWIDTH = 100

func barText(char byte, msg string, length int) string {
	remaining := length - len(msg) - 4
	if remaining <= 0 {
		remaining = 0
	}
	space := " "
	if msg == "" {
		space = "="
	}
	if remaining%2 != 0 {
		msg = msg + space
	}
	msg = space + msg + space
	padding := strings.Repeat(string(char), remaining/2)
	return blue(fmt.Sprintf(" %s%s%s \n", padding, msg, padding))
}

func itemText(key, value string, width int) string {
	if value == "" {
		value = "<unknown>"
	}
	remaining := width - len(key+value) - 3
	if remaining <= 0 {
		remaining = 0
	}
	padding := strings.Repeat(" ", remaining)
	return fmt.Sprintf(" %s:%s%s \n", magBold(key), padding, mag(value))
}

func itemTextArray(key string, values []string, width int) string {
	if len(values) >= 2 {
		key = key + "s"
	}
	return itemText(key, strings.Join(values, ", "), width)
}

func (lab *Lab) Header() string {
	var header string
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return fmt.Sprintf("Could not get terminal size to render lab header: %v", err)
	}
	if width > MAXPRINTWIDTH {
		width = MAXPRINTWIDTH
	}
	header += barText('=', "Starting Lab", width)
	header += itemText("Lab Directory", lab.Directory, width)
	header += itemText("Created At", lab.CreatedAt, width)
	header += itemText("Version", lab.KobleVersion, width)
	header += itemTextArray("Author", lab.Authors, width)
	header += itemTextArray("Email", lab.Emails, width)
	header += itemTextArray("Web", lab.Web, width)
	header += itemText("Description", lab.Description, width)
	header += barText('=', "", width)
	return header + "\n"
}
