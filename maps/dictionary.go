package main

/*
Dictionary map holding strings
*/
type Dictionary map[string]string

/*
Search method for Dictionary
*/
func (d Dictionary) Search(word string) string {
	return d[word]
}
