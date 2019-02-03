package vipertools

import (
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

// Merge all config files in directory.
// This will fail if invalid config files are in the directory.
// err := ReadDir("configs/")
func ReadDir(directory string) error {

	config_files, err := ioutil.ReadDir(directory)

	if err != nil {
		return err
	}

	for _, info := range config_files {
		if _, err := os.Stat(directory + "/" + info.Name()); err == nil {
			viper.SetConfigFile(directory + "/" + info.Name())
			if err := viper.MergeInConfig(); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
