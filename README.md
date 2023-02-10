# opaki-blue
#### Go Hosts File Updater
Tool to add entries to hosts file

This is a simple Go program for updating or adding a line to the Windows hosts file. The program takes two arguments: ipAddress and domainName, which correspond to the IP address and domain name to be added or updated in the hosts file.

## Prerequisites
Before you begin, make sure you have Go installed on your machine. You can download and install Go from the official website: https://golang.org/

## Usage
To run the program, use the following command in the terminal:

`go run script.go <ipAddress> <domainName>`

Where `ipAddress` is the IP address you want to add or update, and `domainName` is the domain name you want to associate with the IP address.

#### Example
To add or update the line **127.0.0.1 example.com** in the hosts file, run the following command:

`go run script.go 127.0.0.1 example.com`

## Contributing
If you want to contribute to this project, feel free to submit a pull request with your changes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.