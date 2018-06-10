package services

import (
	"encoding/json"
	"errors"
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/missioncontrol/utils"
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/utils/cliutils"
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/utils/config"
	"github.com/jfrog/jfrog-cli-go/jfrog-client/utils/errorutils"
	"github.com/jfrog/jfrog-cli-go/jfrog-client/utils/io/httputils"
	"github.com/jfrog/jfrog-cli-go/jfrog-client/utils/log"
	"net/http"
)

func AddService(serviceType, serviceName string, flags *AddServiceFlags) error {
	data := AddServiceRequestContent{
		Type:        serviceType,
		Name:        serviceName,
		Url:         flags.ServiceDetails.Url,
		User:        flags.ServiceDetails.User,
		Password:    flags.ServiceDetails.Password,
		Description: flags.Description,
		SiteName:    flags.SiteName}
	requestContent, err := json.Marshal(data)
	if err != nil {
		return errorutils.CheckError(errors.New("Failed to execute request. " + cliutils.GetDocumentationMessage()))
	}
	missionControlUrl := flags.MissionControlDetails.Url + "api/v3/services"
	httpClientDetails := utils.GetMissionControlHttpClientDetails(flags.MissionControlDetails)
	resp, body, err := httputils.SendPost(missionControlUrl, requestContent, httpClientDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		return errorutils.CheckError(errors.New(resp.Status + ". " + utils.ReadMissionControlHttpMessage(body)))
	}

	log.Debug("Mission Control response: " + resp.Status)
	return nil
}

type AddServiceFlags struct {
	MissionControlDetails      *config.MissionControlDetails
	Description                string
	SiteName                   string
	ServiceDetails *utils.ServiceDetails
}

type AddServiceRequestContent struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Url         string `json:"url,omitempty"`
	User        string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	SiteName    string `json:"site_name,omitempty"`
	Description string `json:"description,omitempty"`
}
