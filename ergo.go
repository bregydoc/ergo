package ergo

import "github.com/oklog/ulid"

// Error define an error basic struct
type Error struct {
	ID                 ulid.ULID   `json:"id"`
	Code               uint64      `json:"code"`
	Err                error       `json:"err"`
	Message            string      `json:"message"`
	DefaultLanguage    *Language   `json:"default_language"`
	AvailableLanguages []*Language `json:"available_languages"`
}
