# gomod-sync

Utility tool to check and update the Go version specified in the `go.mod` files of all exercises.
It works by specifying the desired Go version for all the `go.mod` files to be in. The `check` command
will verify if all `go.mod` files are in the desired version and the `update` command will update all
`go.mod` files to have the desired Go version.

Some exercises must have its `go.mod` specify a Go version that is different from the other exercise's `go.mod`.
This is supported by the `exceptions` key of the configuration file, where an entry must exist for each exercise
that must not have the default version.

## Installing

### Compiling locally

```console
cd gomod-sync
go build
```

This will create an executable `gomod-sync` (`gomod-sync.exe` in windows) in the current directory 
that you can run to execute the program.

### Running without compiling

```console
cd gomod-sync
go run main.go [command] [flags]
```

### Running the tests

```console
cd gomod-sync
go test ./...
```

## Usage

```
  gomod-sync command [flags]

Available Commands:
  check       Checks if all go.mod files are in the target version
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  list        List go.mod files and the Go version they specify
  update      Updates go.mod files to the target version

```

## Commands

- `gomod-sync check -v target_version [-e exercises_path] [-c config_file]`

  checks if all `go.mod` files are in the target version

- `gomod-sync completion`

  generate the autocompletion script for the specified shell
- `gomod-sync help`

  Help about any command
- `gomod-sync list [-e exercises_path]`

  list `go.mod` files and the Go version they specify
- `gomod-sync update -v target_version [-e exercises_path] [-c config_file]`

  updates `go.mod` files to the target version

## Flags

- `-c, --config config_file`
   
  path to the JSON configuration file.  (default `"config.json"`)
  
- `-e, --exercises exercises_path`

  path to the exercises folder. `go.mod` files will be recursively searched inside this directory. (default `"../exercises"`)
- `-v, --goversion target_version`

  target go version that all go.mod files are expected to have. 
  This will be used to check if the `go.mod` files are in the expected version in case of the check command, 
  and to update all `go.mod` files to this version in the case of the update command.
  Using this flag will override the version specified in the config file.

- `-h, --help`

  help for gomod-sync


## Configuration file

Besides the `-v, --goversion` flag, it is also possible to specify the expected go versions for the `go.mod` files in a JSON configuration file.
This file can be given to the program with the `-c, --config file` flag. If the flag is omitted, a file `config.json`
in the current directory will be tried.

With a configuration file, in addition to define a default Go version all exercises' `go.mod` must have, 
it's also possible to configure different versions for different exercises. This can be useful if a particular exercise 
needs a superior version of Go than the default.

This an example of such configuration file:

```json
{
  "default": "1.16",
  "exceptions": [
    {
      "exercise": "strain",
      "version": "1.18"
    }
  ]
}
```

With such configuration, all `go.mod` files will be expected to have the `1.16` version of Go,
except the exercise `strain`, which must have version `1.18` in its `go.mod`.
Specifying the `-v, --goversion` flag overrides the default version specified in this file.

## Examples


* Check if all `go.mod` files of exercises in the `../exercises` folder have the default version 
  specified in the `config.json` file:

  * `gomod-sync check`

* Check if all `go.mod` files of exercises in the `exercises` folder have the `1.16` Go version:

  * `gomod-sync check --goversion 1.16 --exercises ./exercises`

* Update all `go.mod` files of exercises in the `exercises` folder have the `1.16` Go version:

  * `gomod-sync update --goversion 1.16 --exercises ./exercises`

* Update all `go.mod` files, using a config file to specify the versions of exercises:

  * `gomod-sync update --config a_dir/config.json --exercises ./exercises`