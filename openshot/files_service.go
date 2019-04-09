package openshot

import (
	"fmt"

	"github.com/Bimde/httputils"
)

const (
	filesEndpoint        = "/projects/%d/files/"
	fileEndpoint         = "/files/%d/"
	s3DefaultFilesFolder = "files/"
	s3DefaultBucket      = "fancamgenerator"
)

// GetFiles returns a list of all files created for a particular project
func (o *OpenShot) GetFiles(project *Project) (*Files, error) {
	log := getLogger("GetFiles")
	var files Files
	httputils.Get(log, o.filesURL(project.ID), nil, &files)
	return &files, nil
}

// CreateFile adds file to openshot from location on s3. The projectURL of the
// given file (if empty) is overriden with one matching the specified projectID.
// The URL (if empty) is overriden with "files/NAME"
func (o *OpenShot) CreateFile(project *Project, file *FileUploadS3) (*File, error) {
	log := getLogger("CreateFile")
	setDefaults(file, project)
	var createdFile File
	httputils.Post(log, o.filesURL(project.ID), file, &createdFile)
	return &createdFile, nil
}

// CreateFileStruct creates a minimum file struct required for intput to CreateFile
func CreateFileStruct(testFileName string) *FileUploadS3 {
	return &FileUploadS3{JSON: FileS3Info{Name: testFileName}}
}

func setDefaults(file *FileUploadS3, project *Project) {
	if file.ProjectURL == "" {
		file.ProjectURL = project.URL
	}
	if file.JSON.URL == "" {
		file.JSON.URL = s3DefaultFilesFolder + file.JSON.Name
	}
	if file.JSON.Bucket == "" {
		file.JSON.Bucket = s3DefaultBucket
	}
}

// DeleteFile deletes the file from openshot and associated storage
func (o *OpenShot) DeleteFile(fileID int) error {
	log := getLogger("DeleteFile")
	return httputils.Delete(log, o.fileURL(fileID), nil, nil)
}

func (o *OpenShot) filesURL(projectID int) string {
	return fmt.Sprintf(o.BaseURL+filesEndpoint, projectID)
}

func (o *OpenShot) fileURL(fileID int) string {
	return fmt.Sprintf(o.BaseURL+fileEndpoint, fileID)
}
