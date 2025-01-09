// internal/config/config.go
package config

type StubConfig struct {
	AntiDebug   bool
	AntiVM      bool
	SleepDelay  int
	CustomMutex string
	RunOnce     bool
	// Add more configuration options as needed
}
