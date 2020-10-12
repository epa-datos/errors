package errors

import (
	"github.com/jinzhu/gorm"
	"strings"
)

func ParseGormErr(e error) error {
	strError := e.Error()

	if e == gorm.ErrRecordNotFound {
		return Build(
			Operation("Entity not found"),
			NotFound,
		)
	}
	if strings.Contains(strError, "Duplicate") {
		return Build(
			Operation("Insert new record"),
			Conflict,
			Message(searchKeywordValue(strError, "key")+" : "+searchKeywordValue(strError, "entry")+" ya existe."),
		)
	}

	return e
}

func searchKeywordValue(s string, keyword string) string {
	trimText := strings.Replace(s, " ", "", -1)
	startIndex := strings.Index(trimText, keyword)

	if startIndex == -1 {
		return ""
	}

	var result string

	for i := startIndex + len(keyword) + 1; i < len(trimText); i++ {
		singleQoute := "'"

		if string(trimText[i]) == singleQoute {
			break
		}

		result += string(trimText[i])
	}

	return result
}
