package obsapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	utils "github.com/ashpect/obs-cli/obs-api/utils"
)

func About() error {

	client := &http.Client{}

	apiURL := utils.OBS_API + "/about"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

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
