package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")
var ErrWordExists = errors.New("cannot add word, word already exists in dictionary")
var ErrWordDoesNotExist = errors.New("cannot update word, word doesn't exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
