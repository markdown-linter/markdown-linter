package utils

import "io/ioutil"

func ReadFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
