package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errorNoValue = errors.New("not found")
var errorCantUpdate = errors.New("Cant update non-existing word")
var errorWordExists = errors.New("That word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNoValue
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	if err == errorNoValue {
		d[word] = def
	} else if err == nil {
		return errorWordExists
	}
	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errorNoValue:
		return errorCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	// key 가 없으면 아무런 동작 하지 않음, error return 하지 않음
	delete(d, word)
}
