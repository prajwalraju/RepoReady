package utils

import (
	"fmt"
	"os"
)

func GenerateLicense(licenseDirectory string, licenseType string) error {
	licenses, error := GetLicenses()
	if error != nil {
		fmt.Println("An erro occurred while getting Licenses", error)
		return error
	}

	if licenseType == "" {
		var licenseNames []string
		for _, l := range licenses {
			licenseNames = append(licenseNames, l.SpdxId)
		}
		licenseType, error = TakeOptionInput("Enter the license type : ", true, licenseNames)
		if error != nil {
			fmt.Println("An error occurred while taking license type", error)
			return error
		}
	}

	var licenseContent LicenseBody

	for _, l := range licenses {
		if l.SpdxId == licenseType {
			licenseContent, error = GetLicenseContent(l.Url)
			if error != nil {
				fmt.Println("An error occurred while getting license content", error)
				return error
			}
		}
	}

	filePath := licenseDirectory + "/LICENSE"

	// Open the file for writing (create if it doesn't exist)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	if err = writeToFile(file, licenseContent.Body); err != nil {
		fmt.Println("Error in writing to file:", err)
		return err
	}

	return nil
}
