package components_test

import (
	"github.com/bradleyjkemp/cupaloy"
)

const tConfDir = ".snapshots"

func textSnapshot() *cupaloy.Config {
	return cupaloy.New(
		cupaloy.SnapshotSubdirectory(tConfDir),
		cupaloy.SnapshotFileExtension(".txt"),
	)
}

func htmlSnapshot() *cupaloy.Config {
	return cupaloy.New(
		cupaloy.SnapshotSubdirectory(tConfDir),
		cupaloy.SnapshotFileExtension(".html"),
	)
}
