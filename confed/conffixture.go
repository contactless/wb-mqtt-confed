package confed

import (
	"github.com/contactless/wbgo"
	"testing"
)

type ConfFixture struct {
	*wbgo.DataFileFixture
}

func NewConfFixture(t *testing.T) (f *ConfFixture) {
	f = &ConfFixture{wbgo.NewDataFileFixture(t)}
	f.addSampleFiles()
	return
}

func (f *ConfFixture) addSampleFiles() {
	f.CopyDataFilesToTempDir(
		"sample.json",
		"sample.schema.json",
		"sample-comments.json",
		"sample-badsyntax.json",
		"sample-invalid.json",
		"noconfig.schema.json",
		"sample-to-use-after-new-subconf.json",
		"sample_devtypes/msu21.conf",
		"sample_devtypes/wb-mrm2.conf")
}
