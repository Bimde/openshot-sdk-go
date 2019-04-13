[![GoDoc](https://godoc.org/github.com/Bimde/openshot-sdk-go?status.svg)](https://godoc.org/github.com/Bimde/openshot-sdk-go)
[![Build Status](https://travis-ci.com/Bimde/openshot-sdk-go.svg?branch=master)](https://travis-ci.com/Bimde/openshot-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bimde/openshot-sdk-go)](https://goreportcard.com/report/github.com/Bimde/openshot-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Bimde/openshot-sdk-go/blob/master/LICENSE)

# OpenShot SDK for Go
Hello! This sdk provides an easy-to-use go interface for the [OpenShot Cloud API](http://cloud.openshot.org/doc/index.html).

## Installation

```
go get github.com/Bimde/openshot-sdk-go/openshot
```

## How to Use
Add `import "github.com/Bimde/openshot-sdk-go/openshot"` to your file.

Create a new OpenShot client: `client := openshot.New(baseUrl, username, password)`

Create a project: `project := openShot.CreateProject(&openshot.Project{Name: "My Project"})`

Start adding files, clips, animations, transitions and creating exports of your work!

Look at our [![GoDoc](https://godoc.org/github.com/Bimde/openshot-sdk-go?status.svg)](https://godoc.org/github.com/Bimde/openshot-sdk-go) for a complete overview of the functionality available thus far.

This sdk was purpose-built for [this project](https://github.com/Bimde/fancam-generator/), it's a great place to look for example usages. In particular, check out [this package](https://github.com/Bimde/fancam-generator/tree/master/backend/src/trackingconverter). The tests for this package are also quite thorough and can be used as examples.
