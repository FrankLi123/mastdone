# mastdone
a mastodon server analytic tool


## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/mastodon-client.git
    cd mastodon-client
    ```

2. Install necessary tools:
    Run the following command to install all necessary tools:
    ```sh
    ./install.sh
    ```
## Usage

1. Run the application:
    ```sh
    go run main.go
    ```

## Makefile

The `Makefile` contains useful commands to help with the development process. 

### Targets

- `format`: This target uses `gofmt` and `goimports` to format the Go code and fix import statements.
- `lint`: This target uses `golint` to analyze the Go code for style mistakes.

### Example Usage

To format the code, run:
```sh
make format