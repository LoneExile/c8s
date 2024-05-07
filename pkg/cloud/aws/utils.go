package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/go-ini/ini"
)

func selectTagVal(Tags []types.Tag, TagName string) string {
	for _, tag := range Tags {
		if *tag.Key == TagName {
			return *tag.Value
		}
	}
	return ""
}

func GetAWSProfiles() ([]string, error) {
	var profiles []string

	credsFilePath := config.DefaultSharedCredentialsFilename()

	credsFile, err := ini.Load(credsFilePath)
	if err != nil {
		return nil, err
	}
	for _, section := range credsFile.Sections() {
		// fmt.Println(section.Name())
		// fmt.Println(section.Key("aws_access_key_id").String())
		// fmt.Println(section.Key("aws_secret_access_key").String())
		// fmt.Println(section.Key("region").String())
		profiles = append(profiles, section.Name())
	}

	return profiles, nil
}

func ListAWSProfiles() error {
	l, errr := GetAWSProfiles()
	if errr != nil {
		log.Fatalf("Failed to get AWS profiles: %v", errr)
	}
	for i, val := range l {
		if val == "DEFAULT" {
			l = append(l[:i], l[i+1:]...)
		}
	}

	for _, val := range l {
		log.Printf("Profile: %s\n", val)
	}
	return nil
}
