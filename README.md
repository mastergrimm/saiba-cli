# Saiba CLI

<img src="image/banner.png" alt="drawing" width="720"/>

Saiba CLI is a command-line interface tool designed to streamline the creation and management of Saiba Web Projects, particularly focused on SvelteKit applications.

## Features

-   **Create SvelteKit Projects**: Quickly scaffold a new SvelteKit project with the latest features and best practices.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your system:

-   Git
-   Node.js
-   Go

### Installation

To install the Saiba CLI, you can use the Go command:

```bash
go install
```

Then build the executable. Feel free to change the .exe name.

```bash
go build -o saiba.exe main.go
```

## Adding Saiba CLI to Your PATH

To use the saiba command from anywhere on your system, you need to add the directory containing the saiba.exe to your PATH.

**For Windows:**
Search for Environment Variables in your Start menu and select "Edit the system environment variables".
In the System Properties window, click on the "Environment Variables" button.
Under "System Variables", find and select the Path variable and click "Edit".
Click "New" and add the path to the folder where saiba.exe is located.
Click "OK" to save and close all the dialogs.

**For Unix/Linux/macOS:**
Add the following line to your ~/.bashrc, ~/.zshrc, or the appropriate configuration file for your shell:

```bash
export PATH=$PATH:/path/to/saiba
Replace /path/to/saiba with the actual path to the saiba executable.
```

## Usage

To create a new SvelteKit project, simply run:

```bash
saiba create
```

This will create a new SvelteKit project using the template from the official SvelteKit repository.

## Contributing

Contributions to Saiba CLI are welcome! Feel free to fork the repository, make your changes, and submit a pull request.
