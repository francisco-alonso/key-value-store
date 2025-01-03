#!/bin/bash

# Prompt for DockerHub username and password
echo "Please enter your DockerHub username:"
read -r DOCKER_USERNAME

# Securely ask for DockerHub password (hidden input)
echo "Please enter your DockerHub password or access token:"
read -s DOCKER_PASSWORD

# Export the credentials as environment variables
export DOCKER_USERNAME
export DOCKER_PASSWORD

# Run the act deployment with the provided credentials
echo "Running deployment via act..."
act push -j deploy

# Clean up by unsetting the credentials after the deployment
unset DOCKER_USERNAME
unset DOCKER_PASSWORD

echo "Deployment completed."
