package openshot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	//baseURL = "http://cloud.openshot.org"
	baseURL     = "http://18.234.247.207"
	loggingName = "openshot"
)

// OpenShot is the main entry point into the sdk
type OpenShot struct {
	BaseURL string
}

// New creates a new instance of OpenShot with default settings
func New(BaseURL string) *OpenShot {
	return &OpenShot{BaseURL: BaseURL}
}

func getLogger(method string) *log.Entry {
	return log.WithFields(log.Fields{
		"method": fmt.Sprintf("%s#%s", loggingName, method),
	})
}
