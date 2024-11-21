# git-secret-scanner

## Overview

`git-secret-scanner` is a powerful tool designed to detect secrets and credentials in git repositories across GitHub and GitLab organizations. By combining the strengths of TruffleHog and Gitleaks, this scanner provides comprehensive secret detection with enhanced classification capabilities.

## Build and Installation

### Prerequisites

Ensure you have the following installed:
- Go (version 1.20 or higher)
- Make
- `git`
- `TruffleHog` (>= 3.82.13)
- `Gitleaks` (>= 8.21.1)

### Building the Project

1. Clone the repository:

   ```bash
   git clone https://github.com/AkhilSharma90/Git-Secrets-Scanner.git
   cd git-secret-scanner
   ```

2. Build the project:

   ```bash
   # Build the binary
   make build

   # OR for a specific OS
   # Linux
   make build-linux

   # macOS
   make build-darwin

   # Windows
   make build-windows
   ```

### Build Options

The Makefile provides several build-related commands:

```bash
# Clean previous builds
make clean

# Run tests
make test

# Run linters
make lint

# Build and install
make install

# Uninstall
make uninstall
```

### Verification After Build

After building, verify the installation:

```bash
# Check the version
./git-secret-scanner version

# Verify tool dependencies
./git-secret-scanner check-deps
```

## Configuration

### Environment Setup

Before scanning, set up authentication tokens:

```bash
# GitHub Token (repo scope)
export GITHUB_TOKEN="your_github_token"

# GitLab Token (read_api and read_repository scopes)
export GITLAB_TOKEN="your_gitlab_token"
```

### Configuration File

Create a `config.yaml` in the project directory or `~/.config/git-secret-scanner/`:

```yaml
# Global configuration
log_level: info
output_format: csv

# Platform-specific settings
github:
  max_repos: 100
  timeout: 30m

gitlab:
  max_groups: 50
  timeout: 45m

# Ignore patterns
ignore:
  - pattern: "*test*"
  - pattern: "*.md"
```

## Running Scans

### Basic Scan Commands

```bash
# Scan GitHub Organization
./git-secret-scanner github -o "<organization_name>"

# Scan GitLab Group
./git-secret-scanner gitlab -g "<group_name>"
```

### Advanced Scanning

```bash
# Scan with custom output
./git-secret-scanner github -o "<org>" -o output.csv

# Ignore specific secrets
./git-secret-scanner github -o "<org>" -i ignore_list.txt

# Use baseline comparison
./git-secret-scanner github -o "<org>" -b previous_scan.csv
```

## Development

### Running Tests

```bash
# Run all tests
make test

# Run specific test
make test TEST=github
```

## Troubleshooting

### Common Issues

- **Dependency Errors**: Ensure all required tools are installed
- **Authentication Failures**: Check token scopes and permissions
- **Scan Timeouts**: Adjust timeout settings in config

### Logging

Enable detailed logging:

```bash
# Set log level
export GSS_LOG_LEVEL=debug

# Run scan with logging
./git-secret-scanner github -o "<org>"
```