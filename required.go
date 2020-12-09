package vipertools

import (
	"github.com/spf13/viper"
)

var required = map[string]bool{}

func SetRequired(name string) {
	required[name] = true
}

func RemoveRequired(name string) {
	required[name] = false
}

func CheckRequired(p ...*viper.Viper) error {
	if len(v) > 0 {
		v := p[0]

		return nil
	}

	return nil
}
