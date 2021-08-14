package main

import (
	"errors"
)

type Dictionary map[string]string
type DictError string

const (
	ErrNotFound   = DictError("could not find word")
	ErrWordExists = DictError("word already exists")
)

func (de DictError) Error() string {
	return string(de)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
