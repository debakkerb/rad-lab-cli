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
	ParameterDefaultProjectLabels
)

func (c Parameter) String() string {
	switch c {
	case ParameterDefaultProjectLabels:
		return "default-project-labels"
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

type configParameter struct {
	description string
	value       string
}

var parameters map[string]*configParameter

func init() {
	configDirectory, err := checkConfigDirectory()
	if err != nil {
		log.Fatalf("Error while creating configuration file: %v", err)
	}
	readConfiguration(configDirectory)
	initParameters()
}

func readConfiguration(configDirectory string) {
	viper.AddConfigPath(configDirectory)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.SafeWriteConfig()
			if err != nil {
				log.Fatalf("Error while creating configuration file: %v", err)
			}
		}
	}
}

func initParameters() {
	parameters = make(map[string]*configParameter)

	addStringValue(ParameterBillingAccount, "Billing account ID that will be attached to all projects related to RAD Lab.")
	addStringValue(ParameterDirectory, "Local directory where the RAD Lab directory has been cloned.")
	addStringValue(ParameterRegion, "Default region for all resources deployed by RAD Lab.")
	addStringValue(ParameterZone, "Default zone for all resources deployed by RAD Lab.")
	addStringValue(ParameterOrganization, "Organization ID where all RAD Lab projects will be created.")
	addStringValue(ParameterAdminBucket, "Name of the Google Cloud Storage bucket that will store all the RAD Lab state files.")
	addStringValue(ParameterAdminProject, "Project ID which is the Admin project for all RAD Lab resources.")
	addStringValue(ParameterDefaultProjectLabels, "Default labels to add to all RAD Lab projects.")
}

func addStringValue(parameter Parameter, description string) {
	parameters[parameter.String()] = &configParameter{
		description: description,
		value:       fmt.Sprintf("%s", viper.Get(parameter.String())),
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

func getConfigParameterNamesAsString(delimiter string) string {
	keys := make([]string, len(parameters))
	i := 0
	for k := range parameters {
		keys[i] = k
		i++
	}
	return strings.Join(keys, delimiter)
}

func SetString(name string, value interface{}) {
	if isAllowed(name) {
		viper.Set(name, value)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Failed to write to the config file: ", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Error while writing config parameter %s, only these values are allowed: %s", name, getConfigParameterNamesAsString(","))
	}
}

func isAllowed(name string) bool {
	for key, _ := range parameters {
		if key == name {
			return true
		}
	}
	return false
}

func Get(name Parameter) string {
	return fmt.Sprintf("%s", viper.Get(name.String()))
}

func Show() {
	settings := viper.AllSettings()
	for key, value := range settings {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func Usage() string {
	var output strings.Builder
	for k, v := range parameters {
		output.WriteString(fmt.Sprintf("%s: %s\n", k, v.description))
	}
	return output.String()
}
