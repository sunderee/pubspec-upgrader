# pubspec-upgrader

`pubspec-upgrader` is a Go-based CLI tool which automates checking for obsolete versions of Dart dependencies declared in a local `pubspec.yaml` file.

It's kinda unstable at the moment as it won't recognize things like path or Git dependencies and will run a query against `pub.dev`, which will obviously come back as 404.

## Usage

Make sure you have Go installed on your system. Compile the executive and use `--help` flag to get started

```bash
# Compile the executive
go build -o pubspec-upgrader *.go

# Display help message
./pubspec-upgrader --help
```

As you'll see, this will yield

```
Usage of ./pubspec-upgrader:
  -file string
        Path to the pubspec.yaml file
  -ignore-unstable
        Whether unstable versions are ignored (default true)
```

Note on `-ignore-unstable` flag: setting this one to false, you will regard dependencies marked as alpha/beta as appropriate update candidates. Use this if you are living on the edge.

## License

Project is open-sourced under MIT license.
