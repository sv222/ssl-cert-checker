# SSL Certificate Checker

The SSL Certificate Checker is a simple Go program that reads a list of HTTPS addresses from a file, sends a request to each address, and outputs information about the SSL certificate used by the address. The program is designed to be run from the command line, and can be used to quickly check the SSL certificate information for multiple websites.

## Installation

To install the SSL Certificate Checker, you must have Go installed on your system. Once you have Go installed, you can download and build the program using the following commands:

```bash
git clone https://github.com/sv222/ssl-cert-checker.git
cd ssl-cert-checker
docker build -t ssl-cert-checker .
```

This will download the source code for the SSL Certificate Checker, change to the project directory, and build the program.

## Usage

To use the SSL Certificate Checker, create a file that contains a list of HTTPS addresses, one per line. Then, run the program and pass the name of the file as a command-line argument:

```bash
./ssl-cert-checker path-to-file-with-urls.txt
```

This will run the program, read the addresses from the file, and output the SSL certificate information for each address.

```bash
Address: github.com
Common Name: github.com
Issuer: DigiCert TLS Hybrid ECC SHA384 2020 CA1
Valid from: 2023-02-14 00:00:00 +0000 UTC to 2024-03-14 23:59:59 +0000 UTC
Signature algorithm: ECDSA-SHA384
Key algorithm: ECDSA
```

## Docker Usage

You can also use Docker to run the SSL Certificate Checker. First, build the Docker image using the following command:

```bash
docker build -t ssl-cert-checker .
```

Then, run the Docker container and pass the name of the address file as a command-line argument:

```bash
docker run --rm ssl-cert-checker path-to-file-with-urls.txt
```

This will run the program inside a Docker container, mount the address file to the container's /app directory, and pass the file name as a command-line argument to the program.

## Contribution

If you would like to contribute to the SSL Certificate Checker, you can fork the project on GitHub and create a pull request with your changes. Contributions are always welcome, and we appreciate your help in making the program better.

## License

The SSL Certificate Checker is licensed under the MIT License.
