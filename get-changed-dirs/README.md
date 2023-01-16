# Get changed directories

This action gets all changed directories in a PR an cteate outputs which can be used as matrix input.

## Usage

```yaml
- uses: actions/checkout@v3
  with:
    fetch-depth: 0

- name: Get all changed dirs
  uses: GalaxyFinX/actions/get-changed-dirs@f25b5719bafef917bed6315fe0a482f9c3fc9ea4
  with:
    # Paths patterns to detect changes (using `glob` partern matching).
    paths: |
      dev/*.yaml
      stg/*.yaml
```

## Scenarios

```yaml
name: Echo

on:
  push:

jobs:
  get-changed-dirs:
    name: Get all changed directories
    runs-on: ubuntu-latest
    outputs:
      matrix: ${{ steps.get-changed-dirs.outputs.matrix }}
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0

      - name: Get all changed dirs
        id: get-changed-dirs
        uses: GalaxyFinX/actions/get-changed-dirs@f25b5719bafef917bed6315fe0a482f9c3fc9ea4
        with:
          paths: |
            dev/*.yaml
            stg/*.yaml

  echo:
    name: Echo changed dirs
    needs: get-changed-dirs
    runs-on: ubuntu-latest
    strategy:
      matrix: ${{ fromJSON(needs.get-changed-dirs.outputs.matrix) }}
    steps:
      - uses: actions/checkout@v3

      - name: Echo changed dirs
        run: |
          CURRENT_DIR=${{ matrix.dirs }}
          
          echo $CURRENT_DIR
        shell: bash
```
