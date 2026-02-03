# ko

**kevin's swiss army knife** — A collection of command-line utilities built with Go.

Fly related commands:

- fly machine list 
- fly machine destroy <machine_id> --force
- fly deploy

## Overview

`ko` is me, I am `ko`. These are my commands.

## Installation

### From Source

```bash
git clone https://github.com/aflashyrhetoric/ko.git
cd ko
go build -o ko .
```

Then move `ko` to your PATH:

```bash
mv ko /usr/local/bin/
```

## Usage

```bash
ko [command] [options]
```

Get help for any command:

```bash
ko help [command]
ko [command] --help
```

## License

MIT

---

*Built with Go, Cobra.*
