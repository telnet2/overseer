package overseer

import (
	"github.com/telnet2/gopkg/confutil"
)

// HttpCheck represents the condition for http response
type HttpCheck struct {
	URL string `yaml:"url" json:"url,omitempty"`
}

// ReadyCheck represents a condition to determine if an executable is ready for a next step.
type ReadyCheck struct {
	TimeoutSecs int64     `yaml:"timeout_secs" json:"timeout_secs,omitempty"`
	Stdout      string    `yaml:"stdout" json:"stdout,omitempty"`
	Stderr      string    `yaml:"stderr" json:"stderr,omitempty"`
	Http        HttpCheck `yaml:"http" json:"http,omitempty"`
}

// Executable represents an executable to be launched by Overseer.
type Executable struct {
	Id         string     `yaml:"id" json:"id"`
	Command    string     `yaml:"command" json:"command"`
	EnvFile    string     `yaml:"env_file" json:"env_file,omitempty"`
	DependsOn  []string   `yaml:"depends_on" json:"depends_on,omitempty"`
	ReadyCheck ReadyCheck `yaml:"ready_check" json:"ready_check,omitempty"`
}

// Config represents a set of executables to launch via Overseer.
type Config struct {
	Executables   []Executable `yaml:"executables" json:"executables"`
	executableMap map[string]*Executable
	orderedList   []*Executable
}

func NewConfigFromFile(f string) (*Config, error) {
	conf := &Config{}
	if err := confutil.TryConfigParsing(conf, f); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) buildExecutableMap() {
	if len(c.Executables) != len(c.executableMap) {
		for _, exe := range c.Executables {
			c.executableMap[exe.Id] = &exe
		}
		c.runTopoSort()
	}
}

// runTopoSort sorts the executable maps in the order of its dependencies.
// [A,B,C,D...] => A->B,C->D,E
func (c *Config) runTopoSort() {

}

func (c *Config) Launch() {

}
