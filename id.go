package ergo

import (
	"fmt"
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
	id := ulid.MustNew(ulid.Timestamp(g.t), g.entropy)
	fmt.Println(id.String())
	return id
}

var t = time.Unix(1000000, 0)

// UlidGen is a gen instance to create serial ulids
var UlidGen = gen{
	t:       t,
	entropy: ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0),
}
