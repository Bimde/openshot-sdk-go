package main

import (
	"github.com/Bimde/openshot-sdk-go/openshot"
)

const (
	username    = "demo-cloud"
	password    = "demo-password"
	openshotURL = "http://cloud.openshot.org/"
)

func main() {
	// Create an OpenShot client with the server location and credentials to login
	client := openshot.New(openshotURL, username, password)

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
