package obsapi

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashpect/obs-cli/obs-api/utils"
)

func Build(project_name string) error {

	config, err := utils.ReadConfig("config.toml")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return err
	}
	//fmt.Println("Username:", config.ObsCreds.Username)
	//fmt.Println("Password:", config.ObsCreds.Password)

	apiURL := utils.OBS_API + "/build/" + project_name

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL, nil)
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
