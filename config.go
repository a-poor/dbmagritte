package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const (
	defaultMigrationsDir       = "migrations"
	defaultGlobalConfigPath    = "dbmconf.yaml"
	defaultMigrationConfigPath = "conf.yaml"
)

// glogabConfig stores the configuration for the full application
// including the database connection information.
type globalConfig struct {
	Database struct {
		Driver string `yaml:"driver,omitempty"`
		DSN    string `yaml:"dsn,omitempty"`
	} `yaml:"database"`
	MigrationDir string `yaml:"migrationDir,omitempty"`
}

func newGlobalConfig() *globalConfig {
	return &globalConfig{}
}

func readGlobalConfig(path string) (*globalConfig, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg *globalConfig
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
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

func (gc globalConfig) createNewMigration(hash string) (*migrationConfig, error) {
	// Create the migration object
	mc := newMigrationConfig(hash)

	// Does it exist?
	p := path.Join(gc.MigrationDir, hash)
	state, err := whatsThatPath(p)
	if err != nil {
		return nil, err
	}
	if state != PathDoesNotExist {
		return nil, fmt.Errorf("migration directory already exists: %s", p)
	}

	// Create the migration directory.
	err = os.MkdirAll(p, 0755)
	if err != nil {
		return nil, err
	}

	// Create the migration config file.
	err = mc.writeConfig(p)
	if err != nil {
		return nil, err
	}

	// Create the SQL files.
	err = mc.initSQL(p)
	if err != nil {
		return nil, err
	}

	// Success!
	return mc, nil
}

func (gc globalConfig) getMigrationConfig(hash string) error {
	return nil
}

func (gc globalConfig) validateMigrationConfig(mc migrationConfig) error {
	return nil
}

// migrationConfig stores the configuration for a
// single migration.
type migrationConfig struct {
	GitHash     string   `yaml:"gitHash"`
	UpSQL       []string `yaml:"up,omitempty"`
	DownSQL     []string `yaml:"down,omitempty"`
	ValidateSQL []string `yaml:"validate,omitempty"`
}

func newMigrationConfig(hash string) *migrationConfig {
	return &migrationConfig{
		GitHash: hash,
		UpSQL:   []string{"up.sql"},
		DownSQL: []string{"down.sql"},
	}
}

func readMigrationConfig(path string) (*migrationConfig, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML.
	var mc migrationConfig
	err = yaml.Unmarshal(b, &mc)
	if err != nil {
		return nil, err
	}

	// Success!
	return &mc, nil
}

// writeConfig writes the migration config to the given path.
func (mc migrationConfig) writeConfig(p string) error {
	// Convert the config data to YAML.
	b, err := yaml.Marshal(mc)
	if err != nil {
		return err
	}

	// Write the YAML to the file.
	cp := path.Join(p, "conf.yaml")
	err = ioutil.WriteFile(cp, b, 0644)
	if err != nil {
		return err
	}

	// Success!
	return nil
}

// initSQL writes the SQL files to the given path.
func (mc migrationConfig) initSQL(dirPath string) error {
	// Create the empty Up SQL file.
	sup := fmt.Sprintf(
		"-- Migration Name: %s\n-- File Name: %s\n\n",
		mc.GitHash,
		"up.sql",
	)
	pup := path.Join(dirPath, "up.sql")
	err := ioutil.WriteFile(pup, []byte(sup), 0644)
	if err != nil {
		return err
	}

	// Create the empty Down SQL file.
	sdown := fmt.Sprintf(
		"-- Migration Name: %s\n-- File Name: %s\n\n",
		mc.GitHash,
		"up.sql",
	)
	pdown := path.Join(dirPath, "down.sql")
	err = ioutil.WriteFile(pdown, []byte(sdown), 0644)
	if err != nil {
		return err
	}

	return nil
}
