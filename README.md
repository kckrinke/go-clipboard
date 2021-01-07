[![Go Reference](https://pkg.go.dev/badge/github.com/kckrinke/go-clipboard.svg)](https://pkg.go.dev/github.com/kckrinke/go-clipboard)

# Go-Clipboard

Provide copying and pasting to the Clipboard for the Go programming language.

This particular project is a fork of [clipboard](https://github.com/ototto/clipboard)
for the purpose of implementing features necessary for [CDK](https://github.com/kckrinke/go-cdk)
and [CTK](https://github.com/kckrinke/go-ctk).

Any changes made here that are suitable for the upstream project are intended
to be pushed back upstream and eventually not require this project at all. In
the meantime, here we are!

## Build

    > go get github.com/kckrinke/go-clipboard

## Platforms

* OSX
* Windows 7 (probably work on other Windows)
* Linux, Unix (requires 'xclip' or 'xsel' command to be installed)


## Document:

* http://pkg.go.dev/github.com/kckrinke/go-clipboard

## Notes

* Text string only
* UTF-8 text encoding only (no conversion)

## Commands:

paste shell command:

    > go get github.com/kckrinke/go-clipboard/cmd/gopaste
    > # example:
    > gopaste > document.txt

copy shell command:

    > go get github.com/kckrinke/go-clipboard/cmd/gocopy
    > # example:
    > cat document.txt | gocopy

## TODO

* Use [XGB](https://github.com/BurntSushi/xgb) instead of 'xclip' or 'xsel'

## License

This project has a permissive, open source [LICENSE](LICENSE).

## Copyright

  Copyright (c) 2020 Kevin C. Krinke. All rights reserved.

## Upstream Copyright

  Copyright (c) 2013 Ato Araki. All rights reserved.
