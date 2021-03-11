package project

import (
	"github.com/hohice/higo/pkg/file"
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

func mkProjectDir(appPath string) error {
	file.MustCreateDir(appPath)
	file.MustCreateDir(appPath + string(file.Separator) + "api")
	file.MustCreateDir(appPath + string(file.Separator) + "assets")

	var buildpath = appPath + string(file.Separator) + "build"
	file.MustCreateDir(buildpath)
	file.MustCreateDir(buildpath + string(file.Separator) + "ci")
	file.MustCreateDir(buildpath + string(file.Separator) + "pkg")

	file.MustCreateDir(appPath + string(file.Separator) + "cmd")
	file.MustCreateDir(appPath + string(file.Separator) + "configs")
	file.MustCreateDir(appPath + string(file.Separator) + "deployments")
	file.MustCreateDir(appPath + string(file.Separator) + "docspath")
	file.MustCreateDir(appPath + string(file.Separator) + "examples")
	file.MustCreateDir(appPath + string(file.Separator) + "githooks")
	file.MustCreateDir(appPath + string(file.Separator) + "init")

	var internalpath = appPath + string(file.Separator) + "internal"
	file.MustCreateDir(internalpath)
	file.MustCreateDir(internalpath + string(file.Separator) + "app")
	file.MustCreateDir(internalpath + string(file.Separator) + "pkg")
	file.MustCreateDir(appPath + string(file.Separator) + "pkg")
	file.MustCreateDir(appPath + string(file.Separator) + "scripts")
	file.MustCreateDir(appPath + string(file.Separator) + "test")
	file.MustCreateDir(appPath + string(file.Separator) + "third_party")
	file.MustCreateDir(appPath + string(file.Separator) + "tools")
	file.MustCreateDir(appPath + string(file.Separator) + "vendor")

	var webpath = appPath + string(file.Separator) + "web"
	file.MustCreateDir(webpath)
	file.MustCreateDir(webpath + string(file.Separator) + "app")
	file.MustCreateDir(webpath + string(file.Separator) + "static")
	file.MustCreateDir(webpath + string(file.Separator) + "template")
	file.MustCreateDir(appPath + string(file.Separator) + "website")

	return nil
}
