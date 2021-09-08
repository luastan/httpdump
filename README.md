# httpdump

Simple server that dumps HTTP requests to stdout and to the response.

## Usage

By default it listens to 0.0.0.0 on port 8080:

```
httpdump
```

You can change listen address or port as the first argument:

```
httpdump 127.0.0.1:3000
```

## Installation

Using go install:

```
go install github.com/luastan/httpdump@latest
```