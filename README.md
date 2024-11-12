# snest
snest offers functionality to load environent variables straight into a struct through the use of `snest` struct tags.

Features:
* Load environment variables into struct fields through JSON unmarshaling.

TODO:
* Load environment variables into struct recursively allowing for sub-struct fields to be populated.
* Support `[]byte` type.

## Installation
`go get github.com/fletcharoo/snest`

## Usage
Example:
```go
type Config struct {
	URL  string `snest:"API_URL"`
	Port int    `snest:"API_PORT"`
}

func LoadConfig() (conf Config, err error) {
	err = snest.Load(&conf)
	return conf, err
}
```

## Contributing
I'm always open to constructive criticism and help, so if you do wish to contribute to this repo, please abide by the following process:
* Create an issue that describes the bug/feature.
* If you wish to fix/implement this yourself, please create a PR and link to the issue you created (ensure that all code you update/add is well tested).

Please remember that I will not always accept new ideas or code, especially since this repo is intended to be minimalistic and only implement on piece of functionality.
