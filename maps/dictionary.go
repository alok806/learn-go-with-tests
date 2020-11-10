package main

import (
	"errors"
)

/*
ErrNotFound message
*/
var ErrNotFound = errors.New("could not find the word you were looking for")

/*
Dictionary map holding strings
*/
type Dictionary map[string]string

/*
Search method for Dictionary
*/
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
