package config

import (
	"errors"
	"fmt"
	"github.com/debakkerb/rad-lab-cli/validator"
	"github.com/manifoldco/promptui"
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
	ParameterParentID
)

func (c Parameter) String() string {
	switch c {
	case ParameterDefaultProjectLabels:
		return "default-project-labels"
	case ParameterRegion:
		return "region"
	case ParameterBillingAccount:
		return "billing-account-id"
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
	case ParameterParentID:
		return "parent-id"
	default:
		return ""
	}
}

type configParameter struct {
	name        Parameter
	label       string
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

func InitLocalConfiguration() error {
	fmt.Println("Initialising RAD Lab CLI configuration")
	fmt.Println("######################################")

	if err := promptForParameter(*parameters[ParameterBillingAccount.String()], func(input string) error {
		v := validator.New()

		v.Check(input != "", ParameterBillingAccount.String(), "can't be empty")

		if !v.Valid() {

		}
		return nil
	}); err != nil {

	}

	if err := promptForParameter(*parameters[ParameterRegion.String()], func(input string) error {
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func promptForParameter(parameter configParameter, validate func(input string) error) error {
	prompt := promptui.Prompt{
		Label:    parameter.label,
		Validate: validate,
	}

	value, err := prompt.Run()
	if err != nil {
		return err
	}

	err = SetString(parameter.name.String(), value)
	if err != nil {
		return err
	}

	return nil
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

	addStringValue(ParameterBillingAccount, "Billing Account", "Billing account ID that will be attached to all projects related to RAD Lab.")
	addStringValue(ParameterDirectory, "RAD Lab Directory", "Local directory where the RAD Lab directory has been cloned.")
	addStringValue(ParameterRegion, "Region", "Default region for all resources deployed by RAD Lab.")
	addStringValue(ParameterZone, "Zone", "Default zone for all resources deployed by RAD Lab.")
	addStringValue(ParameterOrganization, "Organization ID (organizations/123456789)", "Organization ID where all RAD Lab projects will be created.")
	addStringValue(ParameterAdminBucket, "RAD Lab Admin Storage Bucket", "Name of the Google Cloud Storage bucket that will store all the RAD Lab state files.")
	addStringValue(ParameterAdminProject, "RAD Lab Admin Project", "Project ID which is the Admin project for all RAD Lab resources.")
	addStringValue(ParameterDefaultProjectLabels, "Project labels", "Default labels to add to all RAD Lab projects.")
	addStringValue(ParameterParentID, "Parent ID (organizations/1234, folders/1234)", "Default Parent ID where all RAD Lab resources will be created")
}

func addStringValue(parameter Parameter, label, description string) {
	parameters[parameter.String()] = &configParameter{
		name:        parameter,
		label:       label,
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

func SetString(name string, value interface{}) error {
	if isAllowed(name) {
		viper.Set(name, value)
		err := viper.WriteConfig()
		if err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("Error while writing config parameter %s, only these values are allowed: %s", name, getConfigParameterNamesAsString(", ")))
	}
	return nil
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
	if len(settings) == 0 {
		fmt.Println("No config parameters have been set.  Initialise the local configuration by running 'radlab config init'")
	} else {
		for key, value := range settings {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}

func Usage() string {
	var output strings.Builder
	for k, v := range parameters {
		output.WriteString(fmt.Sprintf("%s: %s\n", k, v.description))
	}
	return output.String()
}
