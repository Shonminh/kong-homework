package yaml

import "gopkg.in/yaml.v2"

func VerifyYamlContent(context string) (data interface{}, err error) {
	if context == "" {
		return "", nil
	}
	if err = yaml.Unmarshal([]byte(context), &data); err != nil {
		return nil, err
	}
	return data, nil
}
