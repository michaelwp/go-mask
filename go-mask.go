package go_mask

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Option struct {
	*Mask
	Word *string
	*Json
}

type Mask struct {
	Char    string
	Length  int
	Prepend string
	Append  string
}

type Json struct {
	Key string
}

type JSON json.RawMessage

func MaskingString(sentence *string, opt ...*Option) {
	if opt == nil {
		return
	}

	for _, o := range opt {
		if o.Word == nil {
			continue
		}

		if o.Mask == nil {
			o.Mask = &Mask{}
		}

		if o.Char == "" {
			o.Char = "*"
		}

		if o.Length <= 0 {
			o.Length = 3
		}

		word := *o.Word
		word = strings.TrimSpace(word)

		mask := strings.Repeat(o.Char, o.Length)
		mask = fmt.Sprintf("%s%s%s", o.Prepend, mask, o.Append)
		*o.Word = mask

		if sentence == nil || word == "" {
			continue
		}

		*sentence = strings.Replace(*sentence, word, mask, -1)
	}
}

// MaskingJSON : `keys` is case-sensitive !!!
func MaskingJSON(jsonData []*JSON, opt ...*Option) error {
	for _, j := range jsonData {
		jMap, err := jsonToMap(j)
		if err != nil {
			return err
		}

		if opt != nil {
			for _, o := range opt {
				data := jMap[o.Key]
				if data != "" {
					o.Word = &data
					MaskingString(nil, o)
					jMap[o.Key] = data
				}
			}
		} else {
			for jk := range jMap {
				data := jMap[jk]
				o := Option{
					Word: &data,
				}
				MaskingString(nil, &o)
				jMap[jk] = data
			}
		}

		jsonByte, err := json.Marshal(jMap)
		if err != nil {
			return err
		}

		jsonString := string(jsonByte)
		*j = *stringToJson(&jsonString)
	}

	return nil
}

// MaskingJSONSlice : `keys` is case-sensitive !!!
func MaskingJSONSlice(jsonData []*JSON, opt ...*Option) error {
	for _, j := range jsonData {
		jMaps, err := JsonToSliceMap(j)
		if err != nil {
			return err
		}

		sliceJson, err := sliceMapToSliceJson(jMaps)
		if err != nil {
			return err
		}

		err = MaskingJSON(sliceJson, opt...)
		if err != nil {
			return err
		}

		sliceMap := make([]map[string]string, len(sliceJson))
		for i, sJson := range sliceJson {
			toMap, err := jsonToMap(sJson)
			if err != nil {
				return err
			}

			sliceMap[i] = toMap
		}

		jData, err := sliceMapToJson(sliceMap)
		if err != nil {
			fmt.Println(err)
		}

		*j = *jData
	}

	return nil
}
