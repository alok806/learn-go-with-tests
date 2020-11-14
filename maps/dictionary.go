package main

/*
ErrNotFound
ErrWordExists
Const error messages
*/
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("Word already exists")
)

/*
DictionaryErr type
*/
type DictionaryErr string

/*
Error string for DictionaryErr type
*/
func (e DictionaryErr) Error() string {
	return string(e)
}

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

/*
Add method for Dictionary
*/
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
