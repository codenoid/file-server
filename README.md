<img align="right" width="100" height="100" src="https://github.com/codenoid/file-server/blob/master/logo.png?raw=true">

# File Server

Multi-source file server built for hyperscale cutting-edge era

## Installation

We provide serveral installation method

### Using TJ's gobinaries.com

```sh
curl -sf https://gobinaries.com/codenoid/file-server/cmd/file-server | sh
```

### Using Go

```sh
go get github.com/codenoid/file-server/cmd/file-server
```

## Config Example

```yml
identity_type: uri-segment
identity_source: last

sources:
  -
    type: fs
    source: /var/app/master
  -
    type: url
    source: https://storage.googleapis.com/awg3igbab/
```

## Usage

take a look at master branch for config.yml file example

```sh
$ file-server --help
  -bind string
    	bind addr (default ":8080")
  -config string
    	path to config file (default "config.yml")
```

after the server are running, you can access localhost:8080/identity-or-file-name.jpg, then file-server will read one by one from source, will return 404 and empty body if file not found

### Assets

Project logo are taken from [http://clipart-library.com/clip-art/books-transparent-background-10.htm](http://clipart-library.com/clip-art/books-transparent-background-10.htm)