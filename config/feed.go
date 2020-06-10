package config

type Feed struct {
	ID              string
	URL             string
	PollingInterval int64
	Type            string
}
