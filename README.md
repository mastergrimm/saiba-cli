<div align="center">
  <img src="image/banner.png" alt="drawing" width="620"/>
</div>

<div align="center"><b>Generate and Assist with building web applications</b></div>

<br />

<h1>Saiba CLI</h1>

> [!WARNING]
> This project is in the early stages of development and is not ready for production use. Many features are currently missing or may not work as expected. Please use with caution and at your own risk.


Saiba CLI is a command-line interface tool designed to streamline the creation and management of Saiba Web Projects, particularly focused on SvelteKit applications.

<img src="image/gif/demo.gif" alt="demo" width="620"/>

## Features

- **Create SvelteKit Projects**: Quickly scaffold a new SvelteKit project with the latest features and best practices.
- **SASS Integration**: Easily add SASS to your projects for more powerful and flexible styling.
- **Lucia Authentication**: Implement Lucia for robust authentication solutions right out of the box.
- **Iconify Icons**: Access a rich library of icons with Iconify integration, enhancing the visual appeal and user interface of your projects.
- **SaibaUI Components**: Utilize SaibaUI to speed up your UI development with pre-built, customizable components tailored for SvelteKit.

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

Then follow the following prompts. "Space" to select options. "Enter" to go the the next field.



## Contributing

Contributions to Saiba CLI are welcome! Feel free to fork the repository, make your changes, and submit a pull request.
