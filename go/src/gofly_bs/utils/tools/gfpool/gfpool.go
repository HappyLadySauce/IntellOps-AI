package gfpool

import (
	"os"
	"time"

	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gpool"
	"gofly/utils/tools/gtype"
)

// Pool pointer pool.
type Pool struct {
	id   *gtype.Int    // Pool id, which is used to mark this pool whether recreated.
	pool *gpool.Pool   // Underlying pool.
	init *gtype.Bool   // Whether initialized, used for marking this file added to fsnotify, and it can only be added just once.
	ttl  time.Duration // Time to live for file pointer items.
}

// File is an item in the pool.
type File struct {
	*os.File             // Underlying file pointer.
	stat     os.FileInfo // State of current file pointer.
	pid      int         // Belonging pool id, which is set when file pointer created. It's used to check whether the pool is recreated.
	pool     *Pool       // Belonging ool.
	flag     int         // Flash for opening file.
	perm     os.FileMode // Permission for opening file.
	path     string      // Absolute path of the file.
}

var (
	// Global file pointer pool.
	pools = gmap.NewStrAnyMap(true)
)
