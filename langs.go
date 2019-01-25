package ergo

import "github.com/oklog/ulid"

type Language struct {
	ID   ulid.ULID `json:"id"`
	Name string    `json:"name"`
}
