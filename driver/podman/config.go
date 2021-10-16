package podman

type PodmanDriverConfig struct {
	DefaultImage string `yaml:"default_image"`
	DetachKeys   string `yaml:"detach_keys"`
}
