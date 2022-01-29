package main

// config stores the configuration for the application.
type config struct {
}

func readConfig(path string) (*config, error) {
	return &config{}, nil
}
