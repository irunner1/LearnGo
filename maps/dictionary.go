package main

// Errors in map.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr for errors in dict.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary map of keys and values.
type Dictionary map[string]string

// Search returns word from dict.
func (d Dictionary) Search(word string) (string, error) {
	definition, status := d[word]

	if !status {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add something to map.
func (d Dictionary) Add(key, value string) error {
	_, status := d.Search(key)

	switch status {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return status
	}
	return nil
}

// Update dict value by key.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete by key from dict.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
