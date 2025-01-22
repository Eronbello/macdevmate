
# MacDevMate

**MacDevMate** is a Go-based command-line tool designed to streamline the installation of common development tools on macOS using [Homebrew](https://brew.sh/).

## Features

- **Interactive Selection**: Choose tools to install interactively using [survey](https://github.com/AlecAivazis/survey).
- **Colored Output**: Enhanced terminal output with colors using [color](https://github.com/fatih/color).
- **Tool Installation**: Quickly install essential development tools like Git, Docker, Go, Node.js, VSCode, Insomnia, and more.

## Prerequisites

- A macOS system with [Homebrew](https://brew.sh/) installed.
  _(If Homebrew is not installed, MacDevMate will automatically guide you through the installation process.)_

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/macdevmate.git
   cd macdevmate
   ```

2. Build the project:
   ```bash
   go build -o macdevmate
   ```

3. Run the tool:
   ```bash
   ./macdevmate
   ```

## Usage

1. Launch the tool in your terminal:
   ```bash
   ./macdevmate
   ```

2. Follow the interactive prompts to select and install your desired tools.

## Tools Included

MacDevMate supports installing a variety of popular tools, including:

- Git
- Docker
- Go
- Node.js
- VSCode
- Insomnia
- And more...

## Contributing

Contributions are welcome! If you'd like to improve the tool or add features, follow these steps:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your message here"
   ```
4. Push your branch and submit a pull request:
   ```bash
   git push origin feature-name
   ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
