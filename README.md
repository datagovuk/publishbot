# Automation Alpha

This repository contains the source code the automation alpha for publishing datasets directly from the original data source.  It provides (will provide) a local web interface for interacting with the service, and adapters for obtaining the source data to be uploaded to [publish-data](https://github.com/datagovuk/publish_data_alpha). The initial alpha will provide adapters that watch a directory for changes, and one to periodically trigger and executable (which can generate data however it likes).

## Installation for development

1. Ensure you have a recent version of [Go](https://golang.org/dl/) installed.
2. Clone this repository anywhere you like (does not require GOPATH).
```bash
git clone git@github.com:datagovuk/publishbot.git
```
3. Change into the directory with ```cd publishbot```
4. Setup the environment with ```make setup```
5. Compile the project ```make```
6. Run the output```bin/publishbot```

## Running tests

```bash
make test
```

## Configuration

You should have a config file, called test.yml for now that contains something like:

```yaml
host: 127.0.0.1
port: 2112
adapters:
  - name: spending
    title: Spend data over £25,000
    type: directory
    arguments:
      folder: ./test-folder
```

## Compiling cross-platform
To compile for windows you should use

```bash
GOOS=windows make
```

and your executable will appear in ```bin/windows_amd64/publishbot.exe```.

## Adding dependencies

To add a dependency, instead of ```go get github.com/org/package``` you should use ```bin/gvt fetch github.com/org/package```.  If gvt is not available, you forgot to run make setup.

## Getting stuck

If things appear to stop compiling, or tests running, ```make clean```  may help.
