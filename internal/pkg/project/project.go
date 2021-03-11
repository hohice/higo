package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hohice/higo/pkg/file"
	"github.com/hohice/higo/pkg/logger"
	"github.com/hohice/higo/pkg/logger/colors"
	"github.com/hohice/higo/pkg/utils"
)

/*
$ tree
.
├── LICENSE.md
├── Makefile
├── README.md
├── api
│   └── README.md
├── assets
│   └── README.md
├── build
│   ├── README.md
│   ├── ci
│   └── package
├── cmd
│   ├── README.md
│   └── _your_app_
├── configs
│   └── README.md
├── deployments
│   └── README.md
├── docs
│   └── README.md
├── examples
│   └── README.md
├── githooks
│   └── README.md
├── init
│   └── README.md
├── internal
│   ├── README.md
│   ├── app
│   │   └── _your_app_
│   └── pkg
│       └── _your_private_lib_
├── pkg
│   ├── README.md
│   └── _your_public_lib_
├── scripts
│   └── README.md
├── test
│   └── README.md
├── third_party
│   └── README.md
├── tools
│   └── README.md
├── vendor
│   └── README.md
├── web
│   ├── README.md
│   ├── app
│   ├── static
│   └── template
└── website
    └── README.md
*/

func NewProject(appname string) error {
	currpath, _ := os.Getwd()
	appPath := filepath.Join(currpath, appname)
	ifneedCreate(appPath)
	logger.Log.Infof(colors.Bold(fmt.Sprintf("create %s", appPath+string(file.Separator))))
	return mkProjectDir(appPath)
}

func ifneedCreate(appPath string) {
	if utils.IsExist(appPath) {
		logger.Log.Errorf(colors.Bold("Application '%s' already exists"), appPath)
		logger.Log.Warn(colors.Bold("Do you want to overwrite it? [Yes|No] "))
		if !utils.AskForConfirmation() {
			os.Exit(2)
		}
	}
}
