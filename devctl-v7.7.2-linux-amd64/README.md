[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/devctl/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/devctl/tree/main)

# devctl

`devctl` is a command-line tool designed to streamline development workflows at Giant Swarm. It provides various commands to help manage repositories, generate files, and bootstrap applications.

## Installation

> **Important**: We recommend downloading the latest release from our [releases page](https://github.com/giantswarm/devctl/releases) rather than using `go install`. This ensures you get a properly built binary with:
> - Correct version information
> - Git commit information for traceability
> - Build timestamps
> - All necessary build flags
> - Generated code and mocks for testing
>
> While `go install` will work, it won't include this important metadata and may miss generated code that helps with debugging and version tracking.

```bash
# Not recommended
go install github.com/giantswarm/devctl/v7@latest
```

**Recommended**: Download the latest release from

https://github.com/giantswarm/devctl/releases

## Features

### App Management (`devctl app`)

The `app` command provides tools for working with Giant Swarm app repositories.

#### Bootstrap Command (`devctl app bootstrap`)

Creates a new app repository from a template with the following features:

- Creates a new repository from the `template-app` template
- Configures sync methods (vendir or kustomize)
- Sets up patch methods (script or kustomize)
- Configures CI/CD automatically
- Creates a pull request with the initial setup

```bash
devctl app bootstrap \
  --name my-app \
  --patch-method script \
  --sync-method vendir \
  --team myteam \
  --upstream-chart helm/upstream \
  --upstream-repo https://github.com/org/repo \
  --github-token-envvar GITHUB_TOKEN
```

Options:
- `--name`: Name of the app (required)
- `--patch-method`: Method to patch upstream changes (script or kustomize)
- `--sync-method`: Method to sync upstream changes (vendir or kustomize)
- `--team`: Team responsible for the app
- `--upstream-chart`: Path to the upstream chart
- `--upstream-repo`: URL of the upstream repository
- `--github-token-envvar`: Name of environment variable containing GitHub token
- `--dry-run`: Only print what would be done

### Repository Management (`devctl repo`)

The `repo` command helps manage GitHub repositories.

#### Setup Command (`devctl repo setup`)

Configures repository settings according to Giant Swarm standards:

- Sets up branch protection rules
- Configures repository settings
- Sets up team permissions
- Configures CI/CD workflows

```bash
devctl repo setup org/repo-name
```

### Code Generation (`devctl gen`)

The `gen` command provides tools for generating various files:

- Workflows
- Makefile
- License
- And more

```bash
devctl gen workflow
devctl gen makefile
devctl gen license
```

### Release Management (`devctl release`)

Helps manage releases in Giant Swarm repositories:

- Creates release branches
- Updates changelog
- Creates GitHub releases

```bash
devctl release create
```

### Replace Command (`devctl replace`)

Helps replace content across multiple files:

```bash
devctl replace old-text new-text
```

### Shell Completion

Provides shell completion for various shells:

```bash
# Bash
devctl completion bash > ~/.bash_completion.d/devctl

# Zsh
devctl completion zsh > "${fpath[1]}/_devctl"

# Fish
devctl completion fish > ~/.config/fish/completions/devctl.fish
```

## Development

### Requirements

- Go 1.21 or later
- Git
- GitHub account with appropriate permissions

### Building from Source

If you want to build from source, use the Makefile which ensures all necessary steps are executed:

```bash
git clone https://github.com/giantswarm/devctl.git
cd devctl
make build     # Build with proper flags and metadata
```

### Running Tests

```bash
go test ./...
```

### Debug Mode

Set `LOG_LEVEL=debug` to see detailed output:

```bash
LOG_LEVEL=debug devctl app bootstrap ...
```

## Contributing

Please check our [contributing guidelines](CONTRIBUTING.md) for details on how to contribute to this project.

## License

devctl is licensed under the [Apache 2.0 License](LICENSE).
