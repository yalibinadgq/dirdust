# dirdust

Small Go tool: declutter ~/Downloads in one command

Side project, maintained when I have time.

## Usage

```bash
./bin/dirdust ~/Downloads --dry-run
./bin/dirdust ~/Downloads
```

## Installation

```bash
go build -o bin/ ./...
```

## What it does

- Single static binary, no runtime deps
- Groups files into folders by extension
- Dry-run prints the plan before moving anything
- Skips hidden files and folders by default

## Project structure

```text
├── .github/
│   └── workflows/
│       └── ci.yml
├── docs/
│   ├── development.md
│   ├── faq.md
│   └── usage.md
├── examples/
│   └── quickstart.md
├── .editorconfig
├── .gitignore
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── SECURITY.md
├── go.mod
└── main.go
```

## Development

```bash
go build ./...
go vet ./...
```

## Acknowledgments

- README structure inspired by popular OSS templates
- Thanks to everyone opening issues with ideas

## License

MIT licensed, see LICENSE.
