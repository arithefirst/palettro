# Palettro

A command-line utility for changing the primary accent color in your desktop rice (customized Linux desktop setup). Palettro allows you to quickly switch between different accent colors across all your configured applications with a single command.

## Features

- **Quick Color Switching**: Change your primary accent color across your entire rice instantly
- **Multi-Application Support**: Update colors in multiple config files simultaneously
- **Service Management**: Automatically restart applications after color changes
- **Predefined Palette**: Ships with a beautiful Catppuccin color selection
- **Template System**: Uses template files to apply colors to your configurations

## Installation

### Prerequisites

- Go 1.24.5 or later
- Linux operating system

### Build from Source

```bash
git clone https://github.com/arithefirst/palettro
cd palettro
go build -o palettro
sudo mv palettro /usr/local/bin/
```

## Usage

### Basic Usage

```bash
# Apply a color theme
palettro -color blue

# Show all available colors
palettro -showcolors

# Show all configured applications
palettro -showconfigs

# Use a custom config file
palettro -config /path/to/custom/config.json -color red
```

### Command Line Options

| Flag | Description | Default |
|------|-------------|---------|
| `-color` | Color name to apply to your desktop theme | Required |
| `-config` | Path to configuration file | `~/.config/palettro/config.json` |
| `-showcolors` | Display all available color names | `false` |
| `-showconfigs` | Show all registered application configs | `false` |

### Configuration File

Palettro uses a JSON configuration file located at `~/.config/palettro/config.json`. If this file doesn't exist, it will be created automatically with default settings.

#### Configuration Structure

```json
{
  "colors": {
    "colorname": {
      "hex": "#ffffff",
      "rgb": "rgb(255, 255, 255)",
      "hsl": "hsl(0deg, 0%, 100%)"
    }
  },
  "configs": [
    {
      "name": "application-name",
      "path": "~/.config/app/",
      "restart": "application-binary"

      // The "restart" option is optional. Only put the name
      // of the binary if it dosen't auto-update config
      // files. Waybar, for example needs to be restarted 
      // for config changes to take affect. Hyprland, for 
      // example auto-updates when changes are detected.
    }
  ]
}
```

#### Default Colors

Palettro ships with the following color palette:

- `rosewater` - #f5e0dc
- `flamingo` - #f2cdcd
- `pink` - #f5c2e7
- `mauve` - #cba6f7
- `red` - #f38ba8
- `maroon` - #eba0ac
- `peach` - #fab387
- `yellow` - #f9e2af
- `green` - #a6e3a1
- `teal` - #94e2d5
- `sky` - #89dceb
- `sapphire` - #74c7ec
- `blue` - #89b4fa
- `lavender` - #b4befe

#### Adding Applications

To add support for a new application:

1. Add an entry to the `configs` array in your configuration file
2. Create template files in `~/.config/palettro/[app-name]/`
3. Use color variables in your templates that palettro can replace

**Example:**

```json
{
  "name": "waybar",
  "path": "~/.config/waybar/",
  "restart": "waybar"
}
```

## Examples

### Switching Themes

```bash
# Apply a blue theme to all configured applications
palettro -color blue

# Apply a red theme with custom config
palettro -config ~/my-themes.json -color red
```

### Listing Available Options

```bash
# See all available colors
palettro -showcolors

# See all configured applications
palettro -showconfigs
```

## Contributing

We welcome contributions to palettro! This project is licensed under the GNU GPL 3.0.

### Development Setup

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/arithefirst/palettro
   cd palettro
   ```
3. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. Make your changes
5. Test your changes:
   ```bash
   go test ./...
   go build -o palettro
   ./palettro -showcolors
   ```
6. Commit and push your changes
7. Create a pull request

### Contribution Guidelines

1. **Code Style**: Follow standard Go formatting (`go fmt`)
2. **Documentation**: Update documentation for new features
3. **Error Handling**: Use appropriate error handling patterns
4. **Commit Messages**: Use descriptive commit messages with the Conventional Commits standard
   
### Reporting Issues

Please report bugs and feature requests through GitHub Issues. Include:

- Operating system and version
- Go version
- Steps to reproduce the issue
- Expected vs actual behavior
- Relevant configuration files

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Default color palette robbed from [Catppuccin](https://github.com/catppuccin/catppuccin) Mocha
- Built with Go's standard library and embedded resources