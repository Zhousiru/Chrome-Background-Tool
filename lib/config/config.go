package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Zhousiru/Chrome-Background-Tool/lib/util"
)

type Config struct {
	Basic struct {
		ChromeProfilePath string `toml:"chrome_profile_path"`
	} `toml:"basic"`
}

var (
	Loaded         = new(Config)
	PreferencePath string
	BGPath         string
)

func init() {
	Load()
}

func Load() {
	path := filepath.Join(util.GetExecutableDir(), "config.toml")
	configData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	_, err = toml.Decode(string(configData), Loaded)
	if err != nil {
		panic(err)
	}

	PreferencePath = filepath.Join(Loaded.Basic.ChromeProfilePath, "Preferences")
	BGPath = filepath.Join(Loaded.Basic.ChromeProfilePath, "background.jpg")
}
