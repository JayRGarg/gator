package config

import (
    "fmt"
    "os"
    "io"
    "encoding/json"
    "path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    DbUrl               string      `json:"db_url"`
    CurrentUserName     string      `json:"current_user_name"`
}

func Read() (Config, error) {
    configFilePath, err := getConfigFilePath();
    if err != nil {
        return Config{}, fmt.Errorf("Failed to get configFilePath: %w", err)
    }
    file, err := os.Open(configFilePath)
    if err != nil {
        return Config{}, fmt.Errorf("Failed to open config file: %w", err)
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return Config{}, fmt.Errorf("Failed to read data from file reader: %w", err)
    }

    var conf Config
    err = json.Unmarshal(data, &conf)
    if err != nil {
        return Config{}, fmt.Errorf("Failed to Unmarshal data into Config struct: %w", err)
    }

    return conf, nil
}

func (c *Config) SetUser(newUserName string) error {
    c.CurrentUserName = newUserName
    err := write(c)
    return err
}

func getConfigFilePath() (string, error) {
    configDirPath, err := os.UserHomeDir()
    if err != nil {
        return "", fmt.Errorf("Failed to get User Home Directory: %w", err)
    }
    configFilePath := filepath.Join(configDirPath, configFileName)
    return configFilePath, nil
}

func write(conf *Config) error {
    configFilePath, err := getConfigFilePath()
    if err != nil {
        return fmt.Errorf("Failed to get configFilePath: %w", err)
    }

    jsonConfigData, err := json.MarshalIndent(*conf, "", "    ")
    if err != nil {
        return fmt.Errorf("Failed to convert Config to json: %w", err)
    }
    err = os.WriteFile(configFilePath, jsonConfigData, 0644)
    return err
}
