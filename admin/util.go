package admin

import (
	"errors"
	"fmt"
	"github.com/debakkerb/rad-lab-cli/config"
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

/*
 * Checks if a parameter has been passed in via the command line.  If not, check the local configuration
 * for a value.  If neither have been provided, throw an error
 */
func checkConfigValue(value string, parameter config.Parameter) (string, error) {
	if value == "" {
		configValue := config.Get(parameter)
		if configValue != "" {
			return configValue, nil
		}
	} else {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("error while getting value for parameter %s: not passed via cli nor does it exist in the local config", parameter.String()))

}
