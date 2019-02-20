package rds

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ParseTagsFromString(payload string) (map[string]string, error) {
	tags := make(map[string]string)
	var err error = nil
	if payload != "" {
		err = json.Unmarshal([]byte(payload), &tags)
	}

	return tags, err
}

func ExpandTagString(input string, namespace string, database string) (string, error) {
	expanded := ""
	for inputIdx := 0; inputIdx < len(input); inputIdx++ {
		if string(input[inputIdx]) == "%" {
			if inputIdx+1 == len(input) {
				return "", errors.New("trailing % character is not allowed")
			}
			// else
			inputIdx++
			nextPercentChar := string(input[inputIdx])
			if nextPercentChar == "%" {
				expanded += "%"
			} else if nextPercentChar == "n" {
				expanded += namespace
			} else if nextPercentChar == "d" {
				expanded += database
			} else {
				return "", errors.New("%" + nextPercentChar + " is an unknown expansion")
			}
		} else {
			expanded += string(input[inputIdx])
		}
	}

	return expanded, nil
}

func GetTags(specTags map[string]string, defaultTags map[string]string, namespace string, database string) (map[string]string, error) {
	tags := map[string]string{}

	// concatenate specTags and defaultTags and expand each key / value
	for key, value := range specTags {
		expandedKey, err := ExpandTagString(key, namespace, database)
		if err != nil {
			return map[string]string{}, err
		}
		expandedValue, err := ExpandTagString(value, namespace, database)
		if err != nil {
			return map[string]string{}, err
		}
		tags[expandedKey] = expandedValue
	}
	for key, value := range defaultTags {
		expandedKey, err := ExpandTagString(key, namespace, database)
		if err != nil {
			return map[string]string{}, err
		}
		expandedValue, err := ExpandTagString(value, namespace, database)
		if err != nil {
			return map[string]string{}, err
		}
		tags[expandedKey] = expandedValue
	}

	return tags, nil
}

func GetAwsTags(specTags map[string]string, defaultTags map[string]string, namespace string, database string) ([]rds.Tag, error) {
	awsTags := []rds.Tag{}

	tags, err := GetTags(specTags, defaultTags, namespace, database)
	if err != nil {
		return []rds.Tag{}, err
	}
	for key, value := range tags {
		awsTags = append(awsTags, rds.Tag{
			Key:   aws.String(key),
			Value: aws.String(value),
		})
	}

	return awsTags, nil
}
