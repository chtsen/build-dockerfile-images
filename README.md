# Docker Image Build and Push Workflow

This repository contains a GitHub Actions workflow for building and pushing Docker images to a private Docker registry. The workflow supports building images for multiple platforms and can be triggered manually or on pull requests to the `main` branch.

## Workflow Details

### Workflow Dispatch Inputs

- **directory**: The directory to build the Docker image. If not specified, the workflow will search for `Dockerfile` in the entire repository.
- **platforms**: The platforms to build for (e.g., `linux/amd64,linux/arm64`). If not specified, the default platforms are `linux/amd64` and `linux/arm64`.

### Environment Variables

- **DOCKER_REGISTRY**: The Docker registry to push the images to. (e.g., `registry.cn-hangzhou.aliyuncs.com`)
- **DOCKER_ORG**: The Docker organization to use. (e.g., `chtsen-sysnc`)

### Secrets

- **DOCKER_ALIYUN_USER**: Username for the private Docker registry.
- **DOCKER_ALIYUN_PW**: Password for the private Docker registry.

## How It Works

1. **Checkout repository**: The workflow checks out the repository to the runner.
2. **Set up Docker Buildx**: The workflow sets up Docker Buildx for multi-platform builds.
3. **Log in to private Docker registry**: The workflow logs in to the specified Docker registry using the provided secrets.
4. **Get list of Dockerfiles**: The workflow searches for Dockerfiles in the specified directory (or the entire repository if no directory is specified).
5. **Build and Push Docker images**: For each Dockerfile found, the workflow:
   - Extracts the directory name as the image name.
   - Extracts the tag from the `FROM` line in the Dockerfile.
   - Builds and pushes the Docker image to the specified Docker registry with the extracted tag.

## Example Usage

### Manually Trigger the Workflow

You can manually trigger the workflow using the "Run workflow" button on the GitHub Actions page. You can specify the `directory` and `platforms` inputs if needed.

### Pull Request Trigger

The workflow will also be triggered automatically on pull requests to the `main` branch.

## Example `workflow_dispatch` Inputs

- **directory**: `./bitnami/redis`
- **platforms**: `linux/amd64,linux/arm64`

These inputs will cause the workflow to search for Dockerfiles in the `./bitnami/redis` directory and build images for both `linux/amd64` and `linux/arm64` platforms.

## Sample Workflow Configuration

Here is the sample workflow configuration for your reference:

```yaml
name: Build Docker Images

on:
  workflow_dispatch:
    inputs:
      directory:
        description: 'Directory to build Docker image'
        required: false
      platforms:
        description: 'Platforms to build for (e.g., linux/amd64,linux/arm64)'
        required: false
  pull_request:
    branches:
      - main

env:
  DOCKER_REGISTRY: registry.cn-hangzhou.aliyuncs.com
  DOCKER_ORG: chtsen-sysnc

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v2

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Log in to private Docker registry
        uses: docker/login-action@v2
        with:
          registry: ${{ env.DOCKER_REGISTRY }}
          username: ${{ secrets.DOCKER_ALIYUN_USER }}
          password: ${{ secrets.DOCKER_ALIYUN_PW }}

      - name: Get list of Dockerfiles
        id: dockerfiles
        run: |
          if [ -z "${{ github.event.inputs.directory }}" ]; then
            files=$(find . -name 'Dockerfile' -print)
          else
            files=$(find "${{ github.event.inputs.directory }}" -name 'Dockerfile' -print)
          fi
          files=$(echo "$files" | tr '\n' ' ')
          echo "::set-output name=files::$files"

      - name: Build and Push Docker images
        run: |
          IFS=' ' read -r -a files <<< "${{ steps.dockerfiles.outputs.files }}"
          for file in "${files[@]}"; do
            if [ -f "$file" ]; then
              dir=$(dirname "$file")
              image_name=$(basename "$dir")
              
              # Extract the tag from the FROM line in the Dockerfile
              from_tag=$(grep '^FROM' "$file" | awk -F':' '{print $NF}' | head -n 1)

              if [ -z "$from_tag" ]; then
                echo "No FROM tag found in $file, skipping."
                continue
              fi

              tag="${{ env.DOCKER_REGISTRY }}/${{ env.DOCKER_ORG }}/${image_name}:${from_tag}"

              echo "Building Docker image for $file with tag $tag"
              docker buildx build --platform "${{ github.event.inputs.platforms || 'linux/amd64,linux/arm64' }}" --push -t "$tag" "$dir"
            else
              echo "Dockerfile $file not found, skipping."
            fi
          done
