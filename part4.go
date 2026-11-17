// Package part4 embeds shards 10 through 12 of the 12 that make up the
// amazon/chronos-2 checkpoint, for package
// [github.com/pinkhop/chronos2-go-base]. Import that package instead: it
// assembles all four parts into a ready-to-use model.
//
// The checkpoint is split across four modules so that each stays small
// enough for the Go module proxy to serve.
package part4

import _ "embed"

// The shards are embedded as byte slices rather than strings because the Go
// linker currently aligns embedded byte-slice data but not string data (not
// a language guarantee; the alias test in the chronos2-go repository's parity
// module verifies it): chronos2.Load can then view the weights in place
// instead of copying them onto the heap. Nothing may write to these slices;
// the model every goroutine shares reads them for the life of the process.
var (
	// shard10 is weights/model-00010-of-00012.safetensors.
	//
	//go:embed weights/model-00010-of-00012.safetensors
	shard10 []byte

	// shard11 is weights/model-00011-of-00012.safetensors.
	//
	//go:embed weights/model-00011-of-00012.safetensors
	shard11 []byte

	// shard12 is weights/model-00012-of-00012.safetensors.
	//
	//go:embed weights/model-00012-of-00012.safetensors
	shard12 []byte
)

// Shards returns this part's safetensors shards in checkpoint order, for
// package base to pass to chronos2.Load. The shards alias the embedded data,
// so callers must never modify them. Each call returns a new outer slice,
// so appending to it never affects another caller.
func Shards() [][]byte {
	return [][]byte{shard10, shard11, shard12}
}
