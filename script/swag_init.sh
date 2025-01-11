#!/bin/bash

# Define the custom directory for the Swagger documentation
CUSTOM_DIR="./docs/api"

# Create the custom directory if it doesn't exist
mkdir -p $CUSTOM_DIR

# Generate the Swagger documentation using swag, specifying main.go as the main file
swag init -g cmd/app/app.go -o $CUSTOM_DIR

echo "Swagger documentation generated in $CUSTOM_DIR"