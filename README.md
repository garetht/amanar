Amanar
===

A tool to programmatically insert refreshed HashiCorp Vault credentials into IntelliJ configurations.


# Usage

The usage of this program depends on three environment variables:

- *VAULT_ADDR*, specifying which Vault server to connect to so that credentials can be retrieved
- *GITHUB_TOKEN*, specifying the personal Github token which will allow refreshed Vault credentials to be retrieved
- *CONFIG_FILEPATH*, specifying the location of your configuration file


# Building

This is a Mac OS-specific project. It may be possible to make this work with a Linux keychain, but no such attempt has been or will be made.

The project has been successfully built on Go `1.8.3` on Mac OS 10.12.5. The mininum possible Go version required is `1.8.1`.

`cgo` is also used to interface with OSX Foundation and Security libraries. You may require `CGO_ENABLED=1` to build this project.

# Regenerating Bindata

We compile the JSON Schema for the Amanar configuration into the Go binary for convenience.

To regenerate this file when the data is updated, run `go-bindata amanar_config_schema.json`.
