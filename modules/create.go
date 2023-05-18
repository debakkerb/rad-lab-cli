package modules

import (
	"fmt"
	"github.com/debakkerb/rad-lab-cli/terraform"
	"github.com/manifoldco/promptui"
	"sort"
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

	_, _, err = prompt.Run()
	if err != nil {
		return err
	}

	return nil
}

func getModuleNames(modules map[string]*terraform.Module) []string {
	moduleNames := make([]string, 0, len(modules))

	for _, value := range modules {
		moduleNames = append(moduleNames, value.Name)
	}

	sort.Strings(moduleNames)

	return moduleNames
}

func getMissingValues(module *terraform.Module) map[string]interface{} {

	for _, value := range module.Variables {
		//if value.Value == nil {
		fmt.Printf("%s: %s\n", value.Name, value.Value)
		//}
	}

	return nil
}
