package modules

import (
	"bufio"
	"fmt"
	"github.com/debakkerb/rad-lab-cli/config"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"os"
	"strings"
)

type Module struct {
	Name string
	Path string
}

type Variable struct {
	Name        string
	Value       interface{}
	Type        string
	Description string
}

func StartWizard() error {
	fmt.Println("Create a module instance")
	fmt.Println("#########################")

	modules, err := getModuleNames()
	if err != nil {
		return err
	}

	for _, module := range modules {
		_, err := getVariables(module)
		if err != nil {
			return err
		}
	}
	return nil
}

func getVariables(module *Module) (interface{}, error) {
	moduleDetails, _ := tfconfig.LoadModule(module.Path)

	var variables []*Variable

	for _, value := range moduleDetails.Variables {
		v := &Variable{
			Name:        value.Name,
			Description: value.Description,
			Type:        value.Type,
		}

		variables = append(variables, v)
	}

	return nil, nil
}

func getModuleNames() ([]*Module, error) {
	moduleDirectories, err := os.ReadDir(fmt.Sprintf("%s/%s", config.Get(config.ParameterDirectory), "modules"))
	if err != nil {
		return nil, err
	}

	var modules []*Module
	for _, module := range moduleDirectories {
		fullPath := fmt.Sprintf("%s/%s/%s", config.Get(config.ParameterDirectory), "modules", module.Name())
		moduleName, err := moduleHumanReadableName(fullPath)
		if err != nil {
			return nil, err
		}

		modules = append(modules, &Module{
			Name: moduleName,
			Path: fullPath,
		})
	}

	return modules, nil
}

func moduleHumanReadableName(filePath string) (string, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s", filePath, "README.md"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			header := strings.TrimSpace(strings.TrimLeft(line, "#"))
			return header, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}
