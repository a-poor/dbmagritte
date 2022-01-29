package main

const (
	defaultGlobalConfigPath    = ".dbm-conf.yaml"
	defaultMigrationConfigPath = "conf.yaml"
)

// glogabConfig stores the configuration for the full application
// including the database connection information.
type globalConfig struct {
}

func newConfig() *globalConfig {
	return &globalConfig{}
}

func readGlobalConfig(path string) (*globalConfig, error) {
	return &globalConfig{}, nil
}

// write writes the global config to the given path.
func (gc globalConfig) write(path string) error {
	return nil
}

// migrationConfig stores the configuration for a
// single migration.
type migrationConfig struct {
}

func readMigrationConfig(path string) (*migrationConfig, error) {
	return &migrationConfig{}, nil
}

// write writes the migration config to the given path.
func (mc migrationConfig) write(path string) error {
	return nil
}
