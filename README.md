## This repository is outdated and no longer maintained.

# ktCloud CLI Client

[![Go Reference](https://pkg.go.dev/badge/github.com/kt-soft-dev/kt-cli#section-directories.svg)](https://pkg.go.dev/github.com/kt-soft-dev/kt-cli#section-directories)

This is a command-line interface (CLI) client for the ktCloud service. It is written in Go and provides peer-to-peer encryption and zero-trust security. The client allows you to interact with the ktCloud service, enabling you to download, upload, and use the API for ktCloud.

## Features

- **P2P Encryption**: All data transferred between the client and ktCloud service is encrypted using peer-to-peer encryption, ensuring the security of your data.
- **Zero-Trust Security**: The client implements a zero-trust security model, meaning it does not inherently trust any entities. This reduces the potential attack surface.
- **Download and Upload**: The client allows you to download and upload files to and from the ktCloud service.
- **API Interaction**: The client provides a way to interact with the ktCloud API, allowing you to perform various operations on the ktCloud service.

## Installation

Download the latest release from the "**Releases**" page and unpack it to a directory in your PATH.
It is recommended to rename the binary to `ktcloud` for easier usage.

## Building from sources

To get the ktCloud CLI client or its libraries, you need to have Go installed on your machine. 

Also you need `task` installed, you can download it from [here](https://taskfile.dev/installation/).

To build the client, run the following commands:

```bash
git clone "https://github.com/bssth/kt-cli"
cd kt-cli
task build-all
```


For ready binaries, see the "**Releases**" page.

## Using as a library for your Go projects

You can use this repository as a library in your Go projects.
To do this,
you need to import the package **github.com/kt-soft-dev/kt-cli/pkg** and use the functions provided by the client.

Code is well documented, see [godoc](https://pkg.go.dev/github.com/kt-soft-dev/kt-cli#section-directories) for details.


## Making API request

To make an API request, you can use -act.method flag to specify the method of the request. For example: 

```bash
ktcloud -act.method=test.test
```

Output will be like this:

```bash
2024/04/01 06:37:17 {"ok":true}
```

Parameters can be passed using **-params** flag.
Value should be a string with space-separated key-value pairs.

For example:

```bash
ktcloud -act.method=test.test -params="param1=value1 param2=value2"'
```

In this example params are just stubs and will be ignored. To get known about parameters for specific method, please read the API documentation.

## Output modes

Output can be displayed in different modes. By default, output is displayed in usual **log.Println** format like this:

```bash
2024/04/01 06:37:17 {"ok":true}
```

**-output** flag can be used to specify output mode. Currently supported modes are:

- **0** - log with timestamp
- **1** - plain log (simple output, no timestamp)
- **2** - just like plain log but without new line at the end

## Flags and environment variables

The client supports the following flags:
- **-debug** - enable debug mode (more verbose output)
- **-config** - path to the configuration file (default: `config.yaml`)
- **-no-interactive** - disable interactive mode. In this mode, the client will not ask for any input from the user. It is useful when running the client in a script or automated environment.
- **-output** - output mode (see above for details)
- **-no-save** - do not save the configuration file after changes by the client. For example, a client usually saves the token after login. This flag disables this behavior.
- **-token** - token for API requests. If this flag is set, the client will use the provided token for API requests instead of the one stored in the configuration file. Client will save the token to the configuration file if the **no-save** flag is not set.
- **-pretty** - pretty print JSON output. It looks better but takes more space and is useless if you want to parse the output.
- **-passwd** - password for encryption and decryption. **It is highly recommended to use environment variable for this purpose instead of passing the password as a flag**.
- **-public** - path to public key file for encryption. Will be downloaded if not set.
- **-private** - path to private key file for decryption. Will be downloaded and decrypted used your provided password if the flag is not set.

Flags for requests and other actions:
- **-params** - parameters for the request. Value should be a string with space-separated key-value pairs. For example: `param1=value1 param2=value2`.
- **-act.ping** - check the connection to the ktCloud.
- **-act.method** - create a request to the API. Value should be a string with the method name.
- **-act.download** - download a file from the ktCloud. Value should be a string with the file ID. You can get it using another flag or from the API.
  - **-act.download.path** - path to save the downloaded file. If not set, the file will be saved in the current directory.
- **-act.upload** - upload a file to the ktCloud. Value should be a string with the path to the file. Also you can upload with **stdin**. In this case, value should be empty.
  - **-act.upload.name** - name of the file on the ktCloud. If not set, the file will be uploaded with its original name. For **stdin** uploads this flag is required.
  - **-act.upload.folder** - folder ID where the file should be uploaded. If not set, the file will be uploaded to the root folder.
  - **-act.upload.disk** - disk ID where the file should be uploaded.
- **-act.files** - get a list of files in the cloud. Value should be a string with the folder ID or "**.**" to fetch user's default disk.
- **-act.keys** - export disks public/private key pairs to files
  - **act.keys.public** - file name for the public key (default is **public_key.pub**)
  - **act.keys.private** - file name for the private key (default is **private_key.asc**)

Environment variables used by the client:
- **KT_CLI_PASSWD** - password for encryption and decryption
- **KT_CLI_TOKEN** - access token for API requests

## Documentation

This readme file is exhaustive enough to get started with the client.
Code is well documented, see [godoc](https://pkg.go.dev/github.com/kt-soft-dev/kt-cli#section-directories) for details.

Click on "Show internal" button to see client-specific documentation, or "pkg" to see library documentation.

## Contributing

Contributions are welcome. Please feel free to submit a pull request or open an issue on the GitHub repository.

## License

See the `LICENSE` file.