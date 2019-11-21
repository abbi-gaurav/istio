package model

import (
	"encoding/json"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/options"
	"istio.io/istio/mixer/template/metric"
	"log"
	"strings"
)

type PassportHeaders = map[string]string
type JsonString = string
type StorageKey = string

type Marshaller struct {
	passportHeaderNames map[string]struct{}
}

func NewMarshaller() *Marshaller {
	passportHeadersString := options.GlobalConfig.PassportHeaders
	if passportHeadersString == "" {
		panic("No passport headers specified")
	}

	keys := strings.Split(passportHeadersString, ",")

	keysMap := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		keysMap[k] = struct{}{}
	}

	return &Marshaller{
		passportHeaderNames: keysMap,
	}
}

func (m *Marshaller) ParseAndValidate(msgs []*metric.InstanceMsg) (JsonString, bool) {
	dk := options.GlobalConfig.DimensionsKey

	passportHeaders := make(map[string]string)

	for _, inst := range msgs {
		key := inst.Dimensions[dk].GetStringValue()
		_, ok := m.passportHeaderNames[key]
		if ok {
			value := inst.Value.GetStringValue()
			passportHeaders[key] = value
		}
	}

	if !m.isValid(passportHeaders) {
		return "", false
	}

	log.Printf("passportHeaderNames %+v\n", m.passportHeaderNames)
	jsonString, err := toJson(passportHeaders)
	if err != nil {
		return "", false
	}

	log.Printf("map %+v\n", jsonString)
	return jsonString, true
}

func (m *Marshaller) isValid(headers PassportHeaders) bool {
	for passportHeaderName := range m.passportHeaderNames {
		value := headers[passportHeaderName]
		if isEmpty(value) {
			log.Printf("Missing element %s in parsed passport headers %+v", passportHeaderName, headers)
			return false
		}
	}
	return true
}

func toJson(headers PassportHeaders) (string, error) {
	jsonBytes, err := json.Marshal(headers)

	if err != nil {
		log.Printf("error in parsing %+v", err)
		return "", err
	}

	return string(jsonBytes), nil
}

func GetStorageKey(msgs []*metric.InstanceMsg) StorageKey {
	dk := options.GlobalConfig.DimensionsKey
	storageKeyName := options.GlobalConfig.StorageKeyName

	for _, inst := range msgs {
		key := inst.Dimensions[dk].GetStringValue()
		if key == storageKeyName {
			value := inst.Value.GetStringValue()
			if isEmpty(value) {
				return ""
			} else {
				return value
			}
		}
	}

	log.Printf("No Storage key %s specified in dimensions", storageKeyName)
	return ""
}

func isEmpty(value string) bool {
	return value == "" || value == options.GlobalConfig.EmptyIdentifier
}
