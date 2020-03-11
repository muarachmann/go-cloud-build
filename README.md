# go-cloud-build

My personal protfolio site (github.com/muarachmann/muarachmann.com). This can be used as a start off point for anyone willing to learn go
and do same!

**NB** This application uses [go]() 1.12.6 as main programming language. Install it if you don't yet have it via the official documentation.

## Setup

Clone the repo

```bash
git clone "github.com/muarachmann/go-cloud-build"
```

Install dependencies locally:

Get into the source's code directory

```bash
cd go-cloud-build
```

We use the [dep](https://golang.github.io/dep/) to managing the various dependencies.
Make sure to download and install it if you don't have. Run the following:

```bash
dep init
```

Configure your app environment variables with the sample `.env-example`

```bash
cp .env-example .env
```

Run the application by typing the following command:

```bash
go run main.go
```

Navigate to the browser and type

```bash
localhost:3000
```

Enjoy!!!


```yaml

name: Run/Deploy website (muarachmann.com)

# Controls when the action will run. Triggers the workflow on push or pull request 
# events but only for the master branch
on:
  push:
    branches: [ master ]
  pull_request:
    branches: [ master ]

jobs:
  go-tests:
    strategy:
      matrix:
        go-version: [1.12]
        platform: [ubuntu-latest]
    runs-on: ${{ matrix.platform }}

    steps:
    - name: Install Go
      uses: actions/setup-go@v1
      with:
        go-version: ${{ matrix.go-version }}

    - name: Debug
      id: debug_id
      run: |
        pwd
        echo ${HOME}
        echo ${GITHUB_WORKSPACE}
        echo ${GOPATH}
        echo ${GOROOT}
      env:
        GOPATH: /home/runner/work/muarachmann.com/muarachmann.com/go

    - name: Check out code into the Go module directory
      uses: actions/checkout@v2
    
    - name: Test & Run go project
      run: go run main.go
```
