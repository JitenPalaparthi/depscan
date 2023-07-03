
# depscan

- depscan is dependency/package manager scanner.It scans dependency files of respective programming languages.
- maven Java    - pom.xml
- gradle Java   - build.gradle
- pip Python    - requirements or requirements.txt
- npm Node      - package-lock.json [it supports lockfileVersion 1 and 2]

## How to install

- The below command works only if go is installed on the machine and GOBIN environment variable is configured. If GOBIN path is set then just call in the terminal. After go install command run depscan.

```
go install github.com/JitenPalaparthi/depscan
```
```
depscan
```



## Build from the source code

```
go build -o depscan main.go
```

## depscan commands and options

- To see help

```
depscan help
```

```
A dependency scanner that scans repositories that are developed using different programming languages.

Usage:
  depscan [flags]
  depscan [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  scan        scan scans a given repository
  version     depscan v0.0.1

Flags:
      --alsologtostderr           Logs are written to standard error as well as to files.
  -h, --help                      help for depscan
      --log_backtrace_at string   When set to a file and line number holding a logging statement,
                                  such as
                                          -log_backtrace_at=gopherflakes.go:234
                                  a stack trace will be written to the Info log whenever execution
                                  hits that statement. (Unlike with -vmodule, the ".go" must be
                                  present.)
      --log_dir string            Log files will be written to this directory instead of the
                                  default temporary directory.
      --logtostderr               Logs are written to standard error instead of to files.
      --stderrthreshold string    Log events at or above this severity are logged to standard
                                  error as well as to files. (default "ERROR")
  -v, --verbosity int             Enable V-leveled logging at the specified level.
      --vmodule string            The syntax of the argument is a comma-separated list of pattern=N,
                                  where pattern is a literal file name (minus the ".go" suffix) or
                                  "glob" pattern and N is a V level. For instance,
                                                -vmodule=gopher*=3
                                  sets the V level to 3 in all Go files whose names begin "gopher".

Use "depscan [command] --help" for more information about a command.
```

- To know current version

```
depscan version
```

- To scan with all default flags

```
depscan scan
```

- depscan scan flags 

```
Flags:
  -d, --depth uint8     the depth of directory recursion for file scans (default 3)
  -f, --format string   output file format. We support two formats json|yaml (default "json")
  -h, --help            help for scan
  -o, --out string      user has to provide output file name (default "output")
  -p, --path string     user has to provide path.Ideally this is a git repository path (default ".")
  ```

  - depscan command example

  ```
  depscan scan --depth=3 --path=/tmp/test_repos/python/eLearning
  ```

  ```
  depscan scan --path=/tmp/test_repos/gradle/zuul --stderrthreshold=INFO --depth=3
  ```

  - depscan command with yaml format and different output file name
  
  ```
  depscan scan --depth=3 --path=/tmp/test_repos/python/eLearning -o pipoutput.json -f yaml
  ```


  - depscan command with json format and different output file name

 ```
 depscan scan --depth=3 --path=/tmp/test_repos/python/eLearning -o pipoutput.json -f json
 ```

  - depscan logs are by default error logs.Can enable Info logs as well

  ```
  depscan scan --depth=3 --path=/tmp/test_repos/python/eLearning --stderrthreshold=INFO
  ```