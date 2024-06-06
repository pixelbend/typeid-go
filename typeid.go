package typeid

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"regexp"
	"strings"
)

type TypeId string

func (ti TypeId) String() string {
	return string(ti)
}
func (ti TypeId) Type() string {
	return strings.Split(string(ti), "_")[0]
}

func (ti TypeId) Id() string {
	return strings.Split(string(ti), "_")[1]
}

func (ti TypeId) Length() int {
	return len(string(ti))
}

func Make(typeName string) (TypeId, error) {
	if err := validateTypeName(typeName); err != nil {
		return "", err
	}

	return TypeId(fmt.Sprintf(`%s_%s`, typeName, ulid.Make().String())), nil
}

func Parse(typeName string, typeId string) (TypeId, error) {
	if err := validateTypeName(typeName); err != nil {
		return "", err
	}

	typeIdParts := strings.Split(typeId, "_")
	if len(typeIdParts) != 2 {
		return "", fmt.Errorf("type id must contain two parts: type and ulid separated by an underscore")
	}

	if typeIdParts[0] != typeName {
		return "", fmt.Errorf("type id name '%s' does not match the given type name '%s'", typeIdParts[0], typeName)
	}

	if _, err := ulid.Parse(typeIdParts[1]); err != nil {
		return "", fmt.Errorf("type id invalid ulid '%s': %w", typeIdParts[1], err)
	}

	return TypeId(typeId), nil
}

func validateTypeName(typeName string) error {
	typeName = strings.TrimSpace(typeName)
	if typeName == "" {
		return fmt.Errorf("type name must not be empty")
	}

	if strings.HasPrefix(typeName, "_") || strings.HasSuffix(typeName, "_") {
		return fmt.Errorf("type name must not start or end with an underscore")
	}

	typeNameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !typeNameRegex.MatchString(typeName) {
		return fmt.Errorf("type name must not contain invalid")
	}

	return nil
}
