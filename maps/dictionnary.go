package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionnary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

type Dictionnary map[string]string
