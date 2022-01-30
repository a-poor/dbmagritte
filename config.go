package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	defaultGlobalConfigPath    = "dbmconf.yaml"
	defaultMigrationConfigPath = "conf.yaml"
)

// glogabConfig stores the configuration for the full application
// including the database connection information.
type globalConfig struct {
	Database struct {
		Driver *string `yaml:"driver"`
		DSN    *string `yaml:"dsn"`
	} `yaml:"database"`
}

func newGlobalConfig() *globalConfig {
	return &globalConfig{}
}

func readGlobalConfig(path string) (*globalConfig, error) {
	return &globalConfig{}, nil
}

// write writes the global config to the given path.
func (gc globalConfig) write(path string) error {
	// Convert the config data to YAML.
	b, err := yaml.Marshal(gc)
	if err != nil {
		return err
	}

	// Write the YAML to the file.
	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	// Success!
	return nil
}

// migrationConfig stores the configuration for a
// single migration.
type migrationConfig struct {
}

func newMigrationConfig(hash string) *migrationConfig {
	return &migrationConfig{}
}

func readMigrationConfig(path string) (*migrationConfig, error) {
	return &migrationConfig{}, nil
}

// write writes the migration config to the given path.
func (mc migrationConfig) write(path string) error {
	return nil
}
