package fileinput

import (
	"io/ioutil"
	"strings"
)

func LoadThen(fileName string, separator string, handler func(s string)) error {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	records := strings.Split(string(data), separator)

	for _, record := range records {
		if record == "" {
			continue
		}

		handler(record)
	}

	return nil
}
