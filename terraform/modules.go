package terraform

import (
	"bufio"
	"fmt"
	"github.com/debakkerb/rad-lab-cli/config"
	"os"
	"regexp"
	"strings"
)

/**
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

func GetModules() (map[string]*Module, error) {
	moduleDirectories, err := os.ReadDir(fmt.Sprintf("%s/%s", config.Get(config.ParameterDirectory), "modules"))
	if err != nil {
		return nil, err
	}

	modules := make(map[string]*Module)
	for _, module := range moduleDirectories {
		fullPath := fmt.Sprintf("%s/%s/%s", config.Get(config.ParameterDirectory), "modules", module.Name())
		moduleName, err := moduleHumanReadableName(fullPath)
		if err != nil {
			return nil, err
		}

		variables, err := getVariables(fullPath)
		if err != nil {
			return nil, err
		}

		modules[formatModuleName(moduleName)] = &Module{
			Name:      formatModuleName(moduleName),
			Path:      fullPath,
			Variables: variables,
		}
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
func formatModuleName(moduleName string) string {
	re := regexp.MustCompile(`RAD Lab|Module`)
	str := re.ReplaceAllString(moduleName, "")
	return strings.TrimSpace(str)
}
