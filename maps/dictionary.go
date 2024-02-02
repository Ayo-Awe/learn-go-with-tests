package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

func (d *Dictionary) Search(word string) (string, error) {
	dict := *d
	definition, found := dict[word]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}
