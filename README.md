Amanar
===

A tool to programmatically insert refreshed HashiCorp Vault credentials into desktop database application configurations.

## Supported Applications
- IntelliJ 2017.2 datasources (e.g. Datagrip)
- IntelliJ 2017.2 run configurations
- Querious 2

## Usage

The usage of this program depends on three environment variables:

- *VAULT_ADDR*, specifying which Vault server to connect to so that credentials can be retrieved
- *GITHUB_TOKEN*, specifying the personal Github token which will allow refreshed Vault credentials to be retrieved
- *CONFIG_FILEPATH*, specifying the location of your configuration file

## Configuration

The configuration file that must be provided is a JSON file conforming to the JSON Schema set forth in `amanar_config_schema.json`. Information on each of the options is given as the `description` attribute in the schema.

Note that IntelliJ-specific paths can be found with [this guide to IntelliJ storage locations](https://www.jetbrains.com/help/idea/directories-used-by-intellij-idea-to-store-settings-caches-plugins-and-logs.html) for global configurations, and usually the `.idea` directory for project-specific storage.

## Building

This is a Mac OS-specific project. It may be possible to make this work with a Linux keychain, but no such attempt has been or will be made.

The project has been successfully built on Go `1.8.3` on Mac OS 10.12.5. The mininum possible Go version required is `1.8.1`.

`cgo` is also used to interface with OSX Foundation and Security libraries as well as for SQLite support for Querious. You may require `CGO_ENABLED=1` to build this project.

## Developing: Regenerating Bindata

We compile the JSON Schema for the Amanar configuration into the Go binary for convenience.

To regenerate this file when the data is updated, run `go-bindata amanar_config_schema.json`.
