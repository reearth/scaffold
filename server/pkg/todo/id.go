package todo

import "github.com/oklog/ulid/v2"

type ID ulid.ULID

func NewID() ID {
	return ID(ulid.Make())
}

func ParseID(s string) (ID, error) {
	id, err := ulid.Parse(s)
	return ID(id), err
}

type IDList []ID
