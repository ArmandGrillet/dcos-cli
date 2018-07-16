package uriresolver

import (
	"errors"
	"path/filepath"
	"strings"
)

// URI is a URI, containing only a scheme and path.
type URI struct {
	Scheme string
	Path   string
}

// Resolver stores its accepted schemes.
// The order of the schemes is important as we fallback on the first one when using Resolve.
type Resolver struct {
	schemes []string
}

// New creates a new resolver.
func New(schemes []string) *Resolver {
	return &Resolver{
		schemes: schemes,
	}
}

// Resolve takes a raw URI and tries to detect its scheme and path.
func (r *Resolver) Resolve(rawURI string) (*URI, error) {
	uriParts := strings.SplitN(rawURI, "://", 2)
	if len(uriParts) == 1 {
		if r.schemes[0] == "file" {
			var err error
			rawURI, err = filepath.Abs(rawURI)
			if err != nil {
				return nil, err
			}
		}
		return &URI{
			Scheme: r.schemes[0],
			Path:   rawURI,
		}, nil
	}

	for _, scheme := range r.schemes {
		if strings.HasPrefix(rawURI, scheme) {
			return &URI{
				Scheme: scheme,
				Path:   strings.TrimPrefix(rawURI, scheme),
			}, nil
		}
	}

	return nil, errors.New("Scheme does not exist in schemes")
}
