# Regenerating Bindata

We compile the JSON Schema for the Amanar configuration into the Go binary for convenience.

To regenerate this file when the data is updated, run `go-bindata amanar_config_schema.json`.
