package modules

import (
	"fmt"
	"github.com/debakkerb/rad-lab-cli/terraform"
	"github.com/manifoldco/promptui"
)

func StartWizard() error {
	fmt.Println("Create a module instance")
	fmt.Println("#########################")

	modules, err := terraform.GetModules()
	if err != nil {
		return err
	}

	prompt := promptui.Select{
		Label: "Select Module",
		Items: getModuleNames(modules),
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	fmt.Printf("You chose %s", result)

	return nil
}

func getModuleNames(modules []*terraform.Module) []string {
	var moduleNames []string

	for _, value := range modules {
		moduleNames = append(moduleNames, value.Name)
	}

	return moduleNames
}
