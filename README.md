# depscan

- depscan is dependency/package manager scanner.It scans dependency files of respective programming languages.
- Maven Java    - pom.xml
- Gradle Java   - build.gradle
- Pip Python    - requirements
- npm Node      - package-lock.json

## How to install

- The below command works only if go is installed on the machine and GOBIN environment variable is configured

```go install github.com/JitenPalaparthi/depscan```

## Build from the source code

```go build -o depscan main.go```

## depscan commands and options

- To see help

```depscan help```

```A dependency scanner that scans repositories developed using different programming languages.

Usage:
  depscan [flags]
  depscan [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  scan        scan scans a given repository
  version     depscan v0.0.1

Flags:
  -h, --help   help for depscan

Use "depscan [command] --help" for more information about a command.
```

- To know current version

```depscan version```

- To scan with all default flags

```depscan scan```

- depscan scan flags 

```Flags:
  -d, --depth uint8     the depth of directory recursion for file scans (default 3)
  -f, --format string   output file format. We support two formats json|yaml (default "json")
  -h, --help            help for scan
  -o, --out string      user has to provide output file name (default "output")
  -p, --path string     user has to provide path.Ideally this is a git repository path (default ".")
  ```

  - depscan command example

  ```depscan scan --depth=3 --path=/tmp/test_repos/python/eLearning```