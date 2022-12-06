package general_goutils

import (
	"golang.org/x/crypto/bcrypt"
	"os"
)

const configPath = "../configurations"
const configFileName = "appsettings.txt"

func HashPassword(password string) (string, error) {
	Logger.Info("hashing password")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "failed to hash password", err
	}
	return string(bytes), nil
}

func CheckPassword(providedPassword, existingPassword string) (bool, string, error) {
	Logger.Info("checking password match")
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(providedPassword))
	if err != nil {
		Logger.Error("failed to match passwords: " + err.Error())
		return false, "wrong password", err
	}
	return true, "password match", nil
}

func IsPasswordHashed(password string) bool {
	Logger.Info("checking if password is already hashed")
	cost, err := bcrypt.Cost([]byte(password))

	if err != nil {
		Logger.Error("failed to find cost")
		return false
	}
	return cost == 14
}

// TODO....this below has to be dealt with

// func createConfigurationsDirectory() bool {
// 	dirExists := DoesDirectoryExist("../configurations") // check directory existence
// 	isCreated := false
// 	if !dirExists {
// 		Logger.Info("configurations directory does not exist")
// 		Logger.Info("creating configurations directory")
// 		err := os.MkdirAll(configPath, os.ModePerm)
// 		if err != nil {
// 			Logger.Error("failed to create configurations directory: " + err.Error())
// 		} else {
// 			Logger.Info("configurations directory created successfully")
// 			isCreated = true
// 		}
// 	}
// 	return isCreated
// }

// CreateAppSettingsFile creates file a
func createAppSettingsFile(isConfigDirPresent bool, encryptedConfig string) bool {
	isCreated := false

	if !isConfigDirPresent {
		Logger.Info("../configurations directory not found")
		return false
	}

	f, err := os.Create(configPath + "/" + configFileName)

	if err != nil {
		Logger.Error("failed to create encrypted config file")
		return false
	} else {

		writeConfigsToFile := func() { // todo...to be refactored as a standalone
			Logger.Info("writing image base64 to file")
			writeString, err := f.WriteString(encryptedConfig)
			if err != nil {
				Logger.Error("failed to write configs to appsettings file")
			}
			if writeString > 0 { // todo...confirm if this makes sense
				Logger.Info("configs written to appsettings file")

				err := f.Close()
				if err != nil {
					Logger.Error("created file not closed properly")
					return
				}

				isCreated = true
			}
		}
		writeConfigsToFile()
	}
	return isCreated
}

//SaveEncryptedConfig saves encrypted config to disk
func SaveEncryptedConfig(encryptedConfig string) {
	// createAppSettingsFile(createConfigurationsDirectory(), encryptedConfig)
}
