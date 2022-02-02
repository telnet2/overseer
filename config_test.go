package overseer

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	iu "github.com/telnet2/gopkg/ioutil"
)

type TestMe struct {
	suite.Suite
	root string
}

func (tm *TestMe) TestNewConfig() {
	confFile := filepath.Join(tm.root, "test-config.yaml")

	_ = ioutil.WriteFile(confFile, []byte(`
executables:
    - id: "ls"
      command: "ls -al $HOME"
    - id: "cat"
      command: "cat $HOME/.bashrc"
      depends_on: 
        - "ls"
    `), 0644)

	conf, err := NewConfigFromFile(confFile)
	tm.NoError(err)

	tm.Len(conf.Executables, 2)
	tm.Equal(conf.Executables[0].Id, "ls")
	tm.Equal(conf.Executables[1].Id, "cat")
}

func TestMain(t *testing.T) {
	test := new(TestMe)

	var remove func()
	test.root, remove = iu.MustNewTmpDir()
	defer remove()

	suite.Run(t, test)
}
