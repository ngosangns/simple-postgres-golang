package wire

import (
	_ "embed"
	"fmt"
	"reflect"
	"strings"
)

//go:embed .env
var envStr string

type Env struct {
	POSTGRES_HOST string
	POSTGRES_PORT string
	POSTGRES_NAME string
	POSTGRES_USER string
	POSTGRES_PASS string
}

var EnvSingleton = NewSingleton(func() *Env {
	_instance := &Env{}
	v := reflect.ValueOf(_instance).Elem()

	// split envStr into lines
	lines := strings.Split(envStr, "\n")

	// iterate over lines
	for _, line := range lines {
		// Skip empty lines and lines starting with #
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Remove comments
		line = strings.SplitN(line, "#", 2)[0]

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid line in .env file:", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Use reflection to check if the key (field) exists in the struct
		f := v.FieldByName(key)

		if f.IsValid() {
			// If the field exists and is settable, update its value
			f.SetString(value)
		}
	}

	return _instance
})
