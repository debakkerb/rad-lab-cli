package terraform

import "github.com/hashicorp/terraform-config-inspect/tfconfig"

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

func getVariables(modulePath string) ([]*Variable, error) {
	moduleDetails, _ := tfconfig.LoadModule(modulePath)

	var variables []*Variable

	for _, value := range moduleDetails.Variables {
		v := &Variable{
			Name:        value.Name,
			Description: value.Description,
			Type:        value.Type,
			Value:       value.Default,
		}

		variables = append(variables, v)
	}

	return variables, nil
}
