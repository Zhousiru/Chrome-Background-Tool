package bg

import (
	"errors"
	"os"
	"regexp"
	"strconv"

	"github.com/Zhousiru/Chrome-Background-Tool/lib/config"
)

var bgPreferenceRe = regexp.MustCompile(`(?mi)"custom_background_local_to_device":\s?(false|true)`)

func GetCustomBGPreference() (bool, error) {
	fileData, err := os.ReadFile(config.PreferencePath)
	if err != nil {
		return false, err
	}

	submatch := bgPreferenceRe.FindStringSubmatch(string(fileData))

	if len(submatch) != 2 {
		return false, errors.New("failed to find key")
	}

	value := submatch[1]

	return strconv.ParseBool(value)
}

func SetCustomBGPreference(b bool) error {
	fileData, err := os.ReadFile(config.PreferencePath)
	if err != nil {
		return err
	}

	edited := bgPreferenceRe.ReplaceAllString(
		string(fileData),
		"\"custom_background_local_to_device\":"+strconv.FormatBool(b),
	)

	return os.WriteFile(config.PreferencePath, []byte(edited), 0664)
}

func SetBG(bgData []byte) error {
	return os.WriteFile(config.BGPath, bgData, 0664)
}
