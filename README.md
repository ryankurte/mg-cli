# mg-cli: a command line utility for mailgun

## Status

Usable. Email sending is a bit grim at the moment with the requirement for HTML and TXT versions, this should be improved.

## Building

1. `go get .` to fetch dependencies
2. `go build ./...` to build utility

## Features

- [x] Create lists
- [ ] Delete lists
- [x] Fetch lists
- [x] Synchronize lists
- [x] Add members to lists
- [ ] Remove members from lists
- [ ] Fetch members of lists
- [ ] Synchronise members of lists
- [x] Send emails to a list

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
  sync-lists   Synchronize existing and desired lists
  version      Show version and exit
```

