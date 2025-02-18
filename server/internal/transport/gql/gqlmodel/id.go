package gqlmodel

import (
	"fmt"
	"io"

	"github.com/oklog/ulid"
)

type ID ulid.ULID

func (id *ID) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("ID must be a string")
	}

	ulid, err := ulid.Parse(s)
	if err != nil {
		return err
	}

	*id = ID(ulid)
	return nil
}

func (id ID) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%s"`, ulid.ULID(id).String())
}
