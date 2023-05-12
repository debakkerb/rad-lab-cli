package modules

import (
	"bufio"
	"fmt"
	"github.com/debakkerb/rad-lab-cli/config"
	"os"
	"regexp"
	"strings"
)

type Module struct {
	Name string
	Path string
}

type TerraformVariable struct {
	VariableName string
	Description  string
	Type         string
	Default      string
}

func StartWizard() error {
	fmt.Println("Create a module instance")
	fmt.Println("#########################")

	modules, err := getModuleNames()
	if err != nil {
		return err
	}

	for _, module := range modules {
		variables, err := getVariables(module)
		if err != nil {
			return err
		}

		for _, variable := range variables {
			fmt.Println(variable.VariableName)
		}
	}
	return nil
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

func getVariables(module *Module) ([]*TerraformVariable, error) {
	fileContent, err := os.ReadFile(fmt.Sprintf("%s/%s", module.Path, "variables.tf"))
	if err != nil {
		return nil, err
	}

	variables := parseVariables(string(fileContent))
	for _, variable := range variables {
		fmt.Println(variable.VariableName)
	}
	return nil, nil
}

func parseVariables(fileContent string) []*TerraformVariable {
	pattern := regexp.MustCompile(`variable\s+"([^"]+)"\s+{([^}]+)}`)
	matches := pattern.FindAllStringSubmatch(fileContent, -1)
	var variables []*TerraformVariable

	for _, match := range matches {
		variableName := match[1]
		variables = append(variables, &TerraformVariable{
			VariableName: variableName,
		})
	}

	return variables
}
