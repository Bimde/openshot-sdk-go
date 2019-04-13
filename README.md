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

## Complete Example
This example shows how to create a project, add a video from an S3 bucket, create a clip with it, change it's x position, and create an export with a different aspect ratio than the input clips.

```
package main

import (
	"github.com/Bimde/openshot-sdk-go/openshot"
)

func main() {
	// Create an OpenShot client with the server location and credentials to login
	client := openshot.New(openshotURL, "username", "password")

	// Create a project with as many customized properties as desired
	// More information at: http://cloud.openshot.org/doc/api_endpoints.html?highlight=location_x#projects
	project, err := client.CreateProject(&openshot.Project{Name: "My Project"})
	if err != nil {
		// deal with error
	}

	// Create a file using a video stored in s3 bucket "mybucket" located at "path/to/file/file_name.mp4" in the bucket
	// More information at: http://cloud.openshot.org/doc/api_endpoints.html?highlight=location_x#files
	file, err := client.CreateFile(project, openshot.CreateFileStruct(openshot.CreateFileS3InfoStruct("file_name.mp4", "path/to/file/", "mybucket")))
	if err != nil {
		// deal with error
	}

	// Create a clip using your new file as it's source
	// More information at: http://cloud.openshot.org/doc/api_endpoints.html?highlight=location_x#clips
	clip, err := client.CreateClip(project, openshot.CreateClipStruct(file, project))
	if err != nil {
		// deal with error
	}

	// Modify clip's x location
	const frame = 120
	const xLocation = 0.5 // Read http://cloud.openshot.org/doc/api_endpoints.html?highlight=location_x#clips for properties and descriptions
	client.AddPropertyPoint(clip, "location_x", frame, xLocation)

	// Remember to call update since adding property points doesn't add them on the server
	// (for efficiency's sake, since many people want to add hundreds of property points!)
	clip, err = client.UpdateClip(clip)
	if err != nil {
		// deal with error
	}

	// Create an export
	export := openshot.CreateDefaultExportStruct(project)
	export.JSON["width"] = 720 // Read http://cloud.openshot.org/doc/api_endpoints.html#exports for available properties
	export, err = client.CreateExport(project, export)
	if err != nil {
		// deal with error
	}

	// Wait until export is ready, either by polling `client.GetExport` or by using a webhook to trigger a seperate handler.
	// More on that here: http://cloud.openshot.org/doc/api_endpoints.html?highlight=location_x#id36

	export, err = client.GetExport(export.ID)
	if err != nil {
		// deal with error
	}

	// Exported video link available at export.Output
}
```

This sdk was purpose-built for [this project](https://github.com/Bimde/fancam-generator/), it's a great place to look for example usages. In particular, check out [this package](https://github.com/Bimde/fancam-generator/tree/master/backend/src/trackingconverter). The tests for this package are also quite thorough and can be used as examples.
