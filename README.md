# mg-cli: a command line utility for mailgun

## Building
1. `go get .` to fetch dependencies
2. `go build ./...` to build utility

## Usage
```
./mg-cli --help                                                                                    [12:15]
Usage:
  mg-cli [OPTIONS] <command>

Application Options:
  -d, --domain=  Mailgun domain for sending [$MG_DOMAIN]
  -k, --api-key= Mailgun API key [$MG_APIKEY]
  -v, --verbose  Enable verbose logging

Help Options:
  -h, --help     Show this help message

Available commands:
  add-member   Add a member to a list
  create-list  Create a new mailing list
  get-lists    Fetch existing mailing lists
  send         Send an email to a list or address
  version      Show version and exit
```

