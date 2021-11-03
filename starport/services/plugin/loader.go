package plugin

// Loader provides plugin control interfaces.
type Loader interface {
	Load(pluginName string) error
}

type loader struct {
}

func (l *loader) Load(name string) error {
	return nil
}

// NewLoader creates plugin loader.
func NewLoader() Loader {
	// TODO: Need default or config parameters to get path configs.
	return nil
}
