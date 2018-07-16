package plugin

import (
	"errors"

	"github.com/dcos/dcos-cli/pkg/uriresolver"
)

func (m *Manager) installExe(name string, uri *uriresolver.URI) error {
	if res.Name == "" {
		return errors.New("Installing an executable requires a name")
	}
	return nil
}

func (m *Manager) installZip(name string, uri *uriresolver.URI) error {
	return nil
}
