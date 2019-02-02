package ergo

import (
	"io"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

type gen struct {
	t       time.Time
	entropy io.Reader
}

func (g *gen) New() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(g.t), g.entropy)
}

// UlidGen is a gen instance to create serial ulids
var UlidGen = gen{
	t:       time.Unix(1000000, 0),
	entropy: ulid.Monotonic(rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano())), 0),
}
