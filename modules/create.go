package modules

import (
	"fmt"
	"github.com/debakkerb/rad-lab-cli/config"
	"github.com/debakkerb/rad-lab-cli/terraform"
	"github.com/manifoldco/promptui"
	"sort"
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

func StartWizard(variablesFile string) error {
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

	billingAccountID := config.Get(config.ParameterBillingAccount)
	parentID := config.Get(config.ParameterParentID)

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
