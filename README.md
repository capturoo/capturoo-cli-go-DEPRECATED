# Capturoo command-line tool for developers

## Usage

capturoo command is run in conjunction with several commands, accounts,
projects or leads. Each command has subcommands. For full usage run
`capturoo <command> --help`. For example, `capturoo leads --help` will
display the help for the `leads` subcommand.

### Account Information
```
$ capturoo accounts info
```

### Create a new project
```
$ capturoo projects create
```

### List projects
```
$ capturoo projects list
```

### Select a project
```
$ capturoo projects select
```

### Show the last lead added
```
$ capturoo leads last
```

### Show a list of leads for a given project
Example: Retrieve a list of leads and write them to `stdout`. Format defaults to `json`.
```
$ captruoo leads list
```

Example: Retrieve a list of leads for `myproject` in CSV format and write them to `mylead.csv` file.
```
$ capturoo leads list -p myproject -f csv -o mylead.csv
```

Example: Retrieve a list of leads for `anotherproject` in JSON format and write them to `mylist.json`.
```
$ capturoo leads list -p anotherproject -f json -o mylist.json
```


