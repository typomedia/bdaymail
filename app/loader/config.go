package loader

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var err error

func Config(config []byte) {
	viper.SetConfigType("toml") // needed as bytes.NewReader() returns a reader with no file extension
	err = viper.ReadConfig(bytes.NewReader(config))
	if err != nil {
		log.Println(err)
	}

	exec, _ := os.Executable()
	path, _ := filepath.Abs(filepath.Dir(exec))

	// Check if config file exists, if not write it
	if _, err := os.Stat(path + "/config.toml"); os.IsNotExist(err) {
		err = viper.WriteConfigAs(path + "/config.toml")
		if err != nil {
			log.Println(err)
		}
	}

	// Read the external config file
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}
