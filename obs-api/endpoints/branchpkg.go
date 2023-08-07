package obsapi

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ashpect/obs-cli/obs-api/utils"
)

func Branchpkg(project_name, package_name, subproject_name string) error {

	config, err := utils.ReadConfig("config.toml")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return err
	}
	//fmt.Println("Username:", config.ObsCreds.Username)
	//fmt.Println("Password:", config.ObsCreds.Password)

	apiURL := utils.OBS_API + "/source/" + project_name + "/" + package_name

	formData := url.Values{}
	formData.Set("cmd", "branch")
	formData.Set("target_project", project_name+":"+subproject_name)

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	auth := config.ObsCreds.Username + ":" + config.ObsCreds.Password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
	req.Header.Set("Accept", "application/xml")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	fmt.Println("API Response:", string(body))
	return nil

}
