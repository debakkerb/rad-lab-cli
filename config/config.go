package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Parameter int8

const (
	ParameterRegion Parameter = iota
	ParameterBillingAccount
	ParameterOrganization
	ParameterZone
	ParameterDirectory
	ParameterAdminProject
	ParameterAdminBucket
)

func (c Parameter) String() string {
	switch c {
	case ParameterRegion:
		return "region"
	case ParameterBillingAccount:
		return "billing-account"
	case ParameterOrganization:
		return "organization"
	case ParameterZone:
		return "zone"
	case ParameterDirectory:
		return "rad-lab-dir"
	case ParameterAdminProject:
		return "admin-project"
	case ParameterAdminBucket:
		return "admin-bucket"
	default:
		return ""
	}
}

var allowedConfigParameters = []string{"rad-lab-dir", "region", "billing-account", "organization", "zone", "admin-project", "admin-bucket"}

func init() {
	configDirectory, err := checkConfigDirectory()
	if err != nil {
		log.Fatalf("Error while creating configuration file: %v", err)
	}

	viper.AddConfigPath(configDirectory)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.SafeWriteConfig()
			if err != nil {
				log.Fatalf("Error while creating configuration file: %v", err)
			}
		}
	}
}

func checkConfigDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "radlab")
	err = os.MkdirAll(configDir, 0700)

	return configDir, err
}

func SetConfigParameter(name, value string) {
	if isAllowed(name) {
		viper.Set(name, value)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Failed to write to the config file: ", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Error while writing config parameter %s, only these values are allowed: %s", name, strings.Join(allowedConfigParameters[:], ", "))
	}
}

func isAllowed(name string) bool {
	for _, v := range allowedConfigParameters {
		if v == name {
			return true
		}
	}
	return false
}

func Get(name Parameter) string {
	return fmt.Sprintf("%s", viper.Get(name.String()))
}
