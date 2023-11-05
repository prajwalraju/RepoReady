package impl

import (
	"fmt"
	"os"

	"github.com/prajwalraju/RepoReady/dto"
	"github.com/prajwalraju/RepoReady/utils"
)

func GenerateLicense(repoInput dto.RepoInput, licenseType string) (dto.RepoInput, error) {
	licenses, error := utils.GetLicenses()
	if error != nil {
		fmt.Println("An erro occurred while getting Licenses", error)
		return repoInput, error
	}

	if licenseType == "" {
		var licenseNames []string
		for _, l := range licenses {
			licenseNames = append(licenseNames, l.SpdxId)
		}
		licenseType, error = utils.TakeOptionInput("Enter the license type : ", true, licenseNames)
		if error != nil {
			fmt.Println("An error occurred while taking license type", error)
			return repoInput, error
		}
	}

	var licenseContent utils.LicenseBody

	for _, l := range licenses {
		if l.SpdxId == licenseType {
			repoInput.StringMetadata["License"] = l.Name
			licenseContent, error = utils.GetLicenseContent(l.Url)
			if error != nil {
				fmt.Println("An error occurred while getting license content", error)
				return repoInput, error
			}
		}
	}

	filePath := repoInput.Name + "/LICENSE"

	// Open the file for writing (create if it doesn't exist)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return repoInput, err
	}
	defer file.Close()

	if err = utils.WriteToFile(filePath, licenseContent.Body); err != nil {
		fmt.Println("Error in writing to file:", err)
		return repoInput, err
	}

	return repoInput, nil
}
