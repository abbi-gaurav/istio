package model

import (
	"encoding/json"
	"istio.io/api/policy/v1beta1"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/options"
	"log"
)

type PassportHeaders = map[string]string
type JsonString = string
type StorageKey = string

func ParseAndValidate(properties map[string]*v1beta1.Value) (JsonString, bool) {
	passportHeaders := make(map[string]string)
	for k, v := range properties {
		strValue := v.GetStringValue()
		if isEmpty(strValue) {
			log.Printf("Empty header %s", k)
			return "", false
		}
		passportHeaders[k] = strValue
	}

	jsonString, err := toJson(passportHeaders)
	if err != nil {
		return "", false
	}

	log.Printf("map %+v\n", jsonString)
	return jsonString, true
}

func toJson(headers PassportHeaders) (string, error) {
	jsonBytes, err := json.Marshal(headers)

	if err != nil {
		log.Printf("error in parsing %+v", err)
		return "", err
	}

	return string(jsonBytes), nil
}

func GetStorageKey(properties map[string]*v1beta1.Value) StorageKey {
	storageKey := options.GlobalConfig.StorageKeyName
	value := properties[storageKey].GetStringValue()

	if isEmpty(value) {
		log.Printf("No Storage key %s specified in dimensions", storageKey)
		return ""
	} else {
		return value
	}
}

func isEmpty(value string) bool {
	return value == "" || value == options.GlobalConfig.EmptyIdentifier
}
