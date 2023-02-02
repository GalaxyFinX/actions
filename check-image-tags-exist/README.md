# Check Image Tag Exists

Check whether or not all of the image tags already exist in the registry.

Currently, we only support ECR.

## Usage

```yaml
  check-image-tag-exists:
    name: Check whether or not all of the image tags exist in the registry
    runs-on: ubuntu-22.04
    steps:
      - uses: actions/checkout@v3

      - name: Check
        uses: GalaxyFinX/actions/check-image-tag-exists@706a45ac9fed0bb7f099f4b1fae457e6725d510e
        with:
          # If we specify 2 images: 1 exist and the other does not, the output will be "0"
          image-names: >
            340396142553.dkr.ecr.ap-southeast-1.amazonaws.com/ecr-public/amazoncorretto/amazoncorretto:16 # This image exists in the repo
            340396142553.dkr.ecr.ap-southeast-1.amazonaws.com/ecr-public/amazoncorretto/amazoncorretto:17 # This image doesn't exist in the repo
          registry-type: ecr
          aws-region: ap-southeast-1
          aws-access-key-id: "xxxxx"
          aws-secret-access-key: "xxxxx"
          aws-session-token: "xxxxx"
```

## Scenarios

### Authenticate ECR with AWS credentials

```yaml
name: Build

on:
  pull_request:
    branches:
      - "main"

permissions:
  contents: read

env:
  REGISTRY: test
  IMAGE_NAME: test

jobs:
  docker:
    runs-on: ubuntu-22.04
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Docker meta
        id: meta
        uses: docker/metadata-action@v4
        with:
          # list of Docker images to use as base name for tags
          images: |
            ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          # generate Docker tags based on the following events/attributes
          tags: |
            type=ref,event=branch
            type=ref,event=pr
            type=semver,pattern={{version}}
          flavor: |
            latest=false
    
      - name: Check if tags already exist in the repo
        id: check-tag
        uses: GalaxyFinX/actions/check-image-tag-exists@706a45ac9fed0bb7f099f4b1fae457e6725d510e
        with:
          image-names: >
            ${{ steps.meta.outputs.tags }}
          registry-type: ecr
          aws-region: ap-southeast-1
          aws-access-key-id: "xxxxx"
          aws-secret-access-key: "xxxxx"

      - name: Build and push image
        uses: GalaxyFinX/actions/docker-build@706a45ac9fed0bb7f099f4b1fae457e6725d510e
        with:
          context: "."
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
```

### Authenticate ECR with AWS Web Identity Token

If your self-hosted runners are running inside EKS cluster, you can setup your workflow like this

```yaml
name: Build

on:
  pull_request:
    branches:
      - "main"

permissions:
  contents: read

env:
  REGISTRY: test
  IMAGE_NAME: test

jobs:
  docker:
    runs-on: ubuntu-22.04
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Docker meta
        id: meta
        uses: docker/metadata-action@v4
        with:
          # list of Docker images to use as base name for tags
          images: |
            ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          # generate Docker tags based on the following events/attributes
          tags: |
            type=ref,event=branch
            type=ref,event=pr
            type=semver,pattern={{version}}
          flavor: |
            latest=false

      - name: Set up web identity token
        id: setup-wit
        run: |
          echo "aws-web-identity-token=$(cat $AWS_WEB_IDENTITY_TOKEN_FILE)" >> $GITHUB_OUTPUT
          echo "aws-role-arn=$AWS_ROLE_ARN" >> $GITHUB_OUTPUT
    
      - name: Check if tags already exist in the repo
        id: check-tag
        uses: GalaxyFinX/actions/check-image-tag-exists@706a45ac9fed0bb7f099f4b1fae457e6725d510e
        with:
          image-names: >
            ${{ steps.meta.outputs.tags }}
          registry-type: ecr
          aws-region: ap-southeast-1
          aws-web-identity-token: ${{ steps.setup-wit.outputs.aws-web-identity-token }}
          aws-role-arn: ${{ steps.setup-wit.outputs.aws-role-arn }}

      - name: Build and push image
        uses: GalaxyFinX/actions/docker-build@706a45ac9fed0bb7f099f4b1fae457e6725d510e
        with:
          context: "."
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
```
