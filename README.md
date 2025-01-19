# System Information Script (Go)

A lightweight Go script that provides detailed system information about your computer. This tool helps system administrators and developers quickly gather essential system metrics and specifications.

## Features

- CPU information (model, cores, usage)
- Memory statistics (total, used, free)
- Disk usage information
- Operating system details
- Network interface information
- Process information

## Prerequisites

- Go 1.16 or higher
- Operating System: Linux, macOS, or Windows

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/go-system-info-script.git

# Navigate to the project directory
cd go-system-info-script

# Install dependencies
go mod tidy
```

## Usage

Run the script using:

```bash
go run main.go
```

## Output Example

```
System Information:
------------------
OS: Darwin 21.6.0
Architecture: amd64
CPU Cores: 8
Memory Total: 16 GB
Memory Free: 8.5 GB
...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
