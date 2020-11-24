package api

import (
	"errors"
	"fmt"
)

const (
	TYPE_ERROR  string = "error"
	TYPE_ACTION string = "action"

	ACTION_HELLO          string = "hello"
	ACTION_HELLO_RESPONSE string = "hello_response"
	ACTION_JOIN_GAME      string = "join_game"
	ACTION_HOST_GAME      string = "host_game"
	ACTION_START_GAME     string = "start_game"
)

var ActionParameters map[string][]string = map[string][]string{
	ACTION_HELLO:          []string{"name"},
	ACTION_HELLO_RESPONSE: []string{"name", "token"},
}

func ValidateActionAndParameter(action string, args map[string]string) error {
	parameters, ok := ActionParameters[action]
	if !ok {
		return errors.New(fmt.Sprintf("'%s' is Not a Valid Action", action))
	}

	for _, param := range parameters {
		if _, ok = args[param]; !ok {
			return errors.New(fmt.Sprintf("'%s' is Missing from Argument list for action %s. Arguments are: '%s'.", param, action, parameters))
		}
	}
	return nil
}
