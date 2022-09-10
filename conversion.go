package go_mask

import "encoding/json"

func JsonToSliceMap(j *JSON) (m []map[string]string, err error) {
	return m, json.Unmarshal(*j, &m)
}

func jsonToMap(j *JSON) (m map[string]string, err error) {
	return m, json.Unmarshal(*j, &m)
}

func stringToJson(s *string) *JSON {
	res := JSON(*s)
	return &res
}

func sliceMapToSliceJson(m []map[string]string) ([]*JSON, error) {
	sliceJson := make([]*JSON, len(m))

	for i, jMap := range m {
		j, err := json.Marshal(jMap)
		if err != nil {
			return nil, err
		}

		jsonString := JSON(j)
		sliceJson[i] = &jsonString
	}

	return sliceJson, nil
}

func sliceMapToJson(m []map[string]string) (*JSON, error) {
	sliceJsonByte, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	sliceJson := JSON(sliceJsonByte)
	return &sliceJson, nil
}
