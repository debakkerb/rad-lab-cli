package modules

import (
	"fmt"
	"github.com/debakkerb/rad-lab-cli/config"
)

func List() {
	radlabDirectory := config.Get(config.ParameterDirectory)
	fmt.Println(radlabDirectory)
}
