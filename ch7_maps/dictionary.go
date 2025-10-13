package main

type IDictionary interface {
	Search(word string) (string, error)
	Add(word, definition string) error
	Update(word, definition string) error
	Delete(word string) error
}

const (
	ErrNotFound         = DictionaryErr("word not found")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d *Dictionary) Search(word string) (string, error) {
	definition, ok := (*d)[word]
	if ok {
		return definition, nil
	} else {
		return "", ErrNotFound
	}
}

func (d *Dictionary) Add(word, definition string) error {
	_, ok := (*d)[word]
	if ok {
		return ErrWordExists
	} else {
		(*d)[word] = definition
		return nil
	}
}

func (d *Dictionary) Update(word, definition string) error {
	_, ok := (*d)[word]
	if ok {
		(*d)[word] = definition
		return nil
	} else {
		return ErrWordDoesNotExist
	}
}

func (d *Dictionary) Delete(word string) error {
	_, ok := (*d)[word]
	if ok {
		delete(*d, word)
		return nil
	} else {
		return ErrWordDoesNotExist
	}
}
