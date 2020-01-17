Amanar
===

A tool to programmatically insert refreshed HashiCorp Vault credentials into desktop database application configurations, allowing you to use database GUIs and CLIs seamlessly with time-limited access credentials.

**You should create backup copies of all configuration files before using this tool.**

## Supported Output Applications and Formats
- Datagrip (tested with 2017.2)
- Intellij IDEA Databases (in theory)
- IntelliJ Run Configurations (tested with 2017.2)
- Querious 2
- Sequel Pro (tested with 1.1.1)
- Postico (tested with 1.2.2)
- Shell script environment variable exports (tested with Bash)
- JSON
- Golang templates

## Usage

The usage of this program depends on two environment variables:

- *GITHUB_TOKEN*, specifying the personal Github token which will allow refreshed Vault credentials to be retrieved
- *CONFIG_FILEPATH*, specifying the location of your configuration file

The program makes certain assumptions about the state of your keychain and configuration files. It cannot be used to create new keychain or configuration entries, only update them. In addition, there should be only one keychain entry per unique identifier (usually the database UUID) so that the correct keychain item to update can be selected without reference to a particular user account.

Multiple vault addresses may now be specified in your configuration file.

## Configuration

**An Overview of the Schema**
![](https://cl.ly/6a4a1f269ba8/Screen%20Shot%202018-09-13%20at%2012.03.19%20AM.png)

The configuration file that must be provided is a JSON or YAML file conforming to the JSON Schema set forth in `amanar_config_schema.json`. Information on each of the options is given as the `description` attribute in the schema.

Note that IntelliJ-specific paths can be found with [this guide to IntelliJ storage locations](https://www.jetbrains.com/help/idea/directories-used-by-intellij-idea-to-store-settings-caches-plugins-and-logs.html) for global configurations, and usually the `.idea` directory for project-specific storage.

## Miscellaneous Notes

- For best results, close applications before running Amanar. Many applications do not take kindly to their data being modified while they are in use.
- Do not edit numbers in plists using XCode. XCode will conveniently change your data types for you.
- A reiteration: back up your data before using this tool. There are no known cases of data loss, but if formats change over time this may occur.

## Building

### Dependencies

Dependencies are managed by Go modules. Run `make build` to build this project. `cgo` is also used to interface with OSX Foundation and Security libraries as well as for SQLite support for Querious. You may require `CGO_ENABLED=1` to build this project.

This is a Mac OS-specific project. It may be possible to make this work with a Linux keychain, but no such attempt has been or will be made.

The project has been successfully built on Go `1.13.6` on Mac OS 10.13.6. The mininum possible Go version required is `1.13`.


## Developing: Extending

To add support for a new data source, do the following:

1. Create a `struct` that satisfies the `Flower` interface. This will act to parse and change the required information on disk.
2. Modify the JSON Schema in accordance with the configuration `struct` and document the required parameters.
3. Regenerate the binary data (see below) that bundles the schema in the Go binary
4. Regenerate the configuration types (see below) from the JSON schema to allow Go to parse the schema.
5. Add the lines in `ProcessConfigItem` to process the new `Flower` that you have created

## Developing: Regenerating Configuration Struct types

We use quicktype to generate the configuration types from the provided JSON schema. Quicktype can be installed from NPM with `npm install -g quicktype`.

To regenerate the types, run `npx quicktype -s schema amanar_config_schema.json -t AmanarConfiguration -l go | sed -E -e 's/json:"(.+)"/json:"\1" yaml:"\1"/g' > amanar_configuration.go`.

To regenerate this file when the data is updated, run `go generate`.

## Developing: Regenerating Bindata

We compile the JSON Schema for the Amanar configuration into the Go binary for convenience using `go-bindata` (`brew install go-bindata`)

To regenerate this file when the data is updated, run `go generate`.
