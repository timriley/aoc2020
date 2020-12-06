package fileinput

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Load(fileName string, separator string, parser func(s string) (interface{}, error)) ([]interface{}, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return []interface{}{}, err
	}

	records := strings.Split(string(data), separator)

	var results []interface{}

	for _, record := range records {
		if record == "" {
			continue
		}

		parsed, err := parser(record)

		if err != nil {
			return results, fmt.Errorf("error parsing record '%v'", record)
		}

		results = append(results, parsed)
	}

	return results, nil
}
