package errors

import (
	"fmt"
	"strconv"
)

type Message string

func EntityNotFoundMessage() Message {
	return Message("Entity not found")
}

func IdNotFoundMessage(id interface{}) Message {
	return Message(fmt.Sprintf("ID: %s not found", idToString(id)))
}

func NotEnoughPermissionsMessage(userID interface{}, resource string) Message {
	id := idToString(userID)
	return Message(fmt.Sprintf("User: %s does not have enough permission to %s resource", id, resource))
}

func DuplicatedEntryMessage() Message {
	return Message("Duplicated entry")
}

func idToString(id interface{}) string {
	var formatedID string

	switch typ := id.(type) {
	case string:
		formatedID = typ
	case int64:
		formatedID = strconv.FormatInt(typ, 10)
	case int32:
		formatedID = strconv.FormatInt(int64(typ), 10)
	case uint64:
		formatedID = strconv.FormatUint(typ, 10)
	case uint32:
		formatedID = strconv.FormatUint(uint64(typ), 10)
	default:
		formatedID = fmt.Sprintf("Unexpected type ID")
	}

	return formatedID
}
