# lacuna

`lacuna` is a simple, lightweight, and easy-to-use CLI tool that works as a word-finding query engine designed for developers. This command line tool allows you to find words that match specific constraints and are likely in a given context. You can specify constraints based on meaning, spelling, sound, and vocabulary to refine your word search queries.

## Features

`lacuna` provides the following features:

- **Means-Like**: Find words that have similar meanings to a given word.
- **Sound-Like**: Find words that sound similar to a given word.
- **Spelled-Like**: Find words that are spelled similarly to a given word.

## Installation
To use  lacuna , you need to have Go installed. You can install  lacuna  by cloning the repository and building the binary locally.

### Prerequisites
Make sure you have Go installed on your machine. You can download and install Go from the [official Go website](https://golang.org/dl/).

### Build from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/sxmbaka/lacuna.git
   ```

2. Navigate to the project directory:

   ```bash
   cd lacuna
   ```

3. Build the binary:
    ```bash
    # to build the binary in the current directory
    go build
    ```

### Install the Binary

After building the binary, you can install it to a directory in your system's PATH to make it accessible globally.

```bash
# to install the binary to the Go binary (see go env GOBIN) directory
go install
```

This command will build the binary and install it to the Go binary directory, which should be included in your system's PATH. After installation, you can use `lacuna` from any terminal window by typing `lacuna`.

### Verify Installation

To verify that `lacuna` is installed correctly, you can run the following command:

```bash
lacuna
# or
lacuna --help
```
If this does not work, you may need to add the Go binary directory to your system's PATH.

## Adding `behead` to the System PATH

To make the `behead` command accessible from any terminal or command prompt window, you need to add the directory containing the `behead` executable to your system's PATH environment variable. Here's how to do it for different operating systems:

### Linux and macOS

1. Open a terminal window.

2. Edit your shell configuration file (e.g., `~/.bashrc` for Bash or `~/.zshrc` for Zsh) using a text editor:

   ```bash
   vim ~/.bashrc
   ```

   or

   ```bash
   vim ~/.zshrc
   ```

3. Add the following line at the end of the file to export the Go binary directory to the PATH:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

4. Save the file and exit the text editor.

5. Source the updated configuration file to apply the changes:

   ```bash
   source ~/.bashrc
   ```

   or

   ```bash
   source ~/.zshrc
   ```

Now you can use the `behead` command from any terminal window.

### Windows

1. Open the Start menu and search for "Environment Variables".

2. Click on "Edit the system environment variables".

3. In the System Properties window, click on "Environment Variables".

4. Under "System Variables", select the "Path" variable and click "Edit".

5. Click "New" and add the path to the Go binary directory (e.g., `C:\Go\bin`).

6. Click "OK" to save the changes.

Now you can use the `behead` command from any Command Prompt or PowerShell window.

## Usage

`lacuna` provides a simple and intuitive interface for finding words based on various constraints. You can use the tool to find words that match specific criteria, such as meaning, spelling, sound, and vocabulary. Two ways to use the tool:

### Root command
when you need to use all the features in combination use the root command and use the proper flags for structuring your query. Example: 
```bash
lacuna --means-like="light" -p bun
lacuna -m hero --sounds-like "king"
lacuna -m justice -s "king" -p bun
# and any other combination
```
### Specific commands
Available commands are `ml`, `sl`, and `pl`.

When you need to use only a single feature use the command associated with it. You CANNOT use other features when using a specific feature command. Examples:
```bash
lacuna ml justice
lacuna sl light
lacuna pl bun
```
> Tip: If you face problems when giving arguments with spaces of special characters, try enclosing the argument in double quotes. Infact it is a good practice to always enclose the argument in double quotes.

## Examples

You can run `lacuna examples` to see some examples of how to use the tool.

## Contributing

If you have any suggestions, feature requests, or bug reports, please open an issue on the GitHub repository. We welcome contributions from the community and are open to pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

