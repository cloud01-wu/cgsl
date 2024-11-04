package utils

import (
	"errors"
	"fmt"
	"strings"
)

type Urn struct {
	Service  string
	Resource string
	Account  string
	Object   string
}

func NewURN(service string, resource string, account string, object string) string {
	return fmt.Sprintf("urn:%v:%v:%v:%v", service, resource, account, object)
}

func ParseURN(urn string) (Urn, error) {
	var err error = nil
	var urnObject = Urn{}

	tokens := strings.Split(urn, ":")
	if len(tokens) != 5 || tokens[0] != "urn" {
		err = errors.New("INVALID URN")
	}

	urnObject.Service = tokens[1]
	urnObject.Resource = tokens[2]
	urnObject.Account = tokens[3]
	urnObject.Object = tokens[4]

	return urnObject, err
}
