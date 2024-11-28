# Cheetah - Multi-threaded Web Directory Discovery with HTTP Method Tampering

A fast, multi-threaded Go program designed for discovering hidden web directories and testing HTTP methods on each directory. The program sends requests using various HTTP methods (GET, POST, PUT, DELETE, OPTIONS, HEAD) and shows detailed response information, including status codes and headers.

## Features

- **Multi-threaded Directory Discovery**: Uses Go's concurrency model (goroutines) to send HTTP requests to multiple directories concurrently.
- **HTTP Method Tampering**: Tests common HTTP methods (GET, POST, PUT, DELETE, OPTIONS, HEAD) for each URL.
- **Customizable Concurrency**: Configurable limit on the number of concurrent requests.
- **Detailed Output**: Displays status code, headers, and the HTTP method used for each request.
- **Input File Support**: Accepts a file containing a list of directory paths to test.

## Requirements

- **Go**: The program requires Go version 1.16 or above.
- **Linux/macOS/Windows**: The program is designed to work on all major platforms.

## Installation

1. **Clone the repository**:
```bash
git clone https://github.com/0xAminED/Cheetah.git
cd Cheetah
```
2. **Build the program**:
```bash
go build -o Cheetah
```

3. **Prepare an input file**: The input file should contain one directory path per line (e.g., ```admin```, ```login```, ```uploads```).

4. **Run the program**:
```bash
./Cheetah -u https://example.com -i dirs.txt
```

## Usage

### Command-Line Flags:

- ```-u```: Target URL (e.g., ```https://example.com```)
- ```-i```: Path to the input file containing the list of directories to test.


### Example:
```bash
./Cheetah -u https://example.com -i dirs.txt
```
Where ```dirs.txt``` contains directory paths like:

```bash
admin
login
uploads
images
```
### Sample Output:

```bash
Method: GET | URL: https://example.com/admin
Status Code: 200
Headers: map[Content-Type:[text/html; charset=UTF-8] Date:[Mon, 27 Nov 2024 17:15:09 GMT]]

Method: POST | URL: https://example.com/login
Status Code: 405
Headers: map[Allow:[GET, HEAD, OPTIONS]]

Method: OPTIONS | URL: https://example.com/images
Status Code: 200
Headers: map[Allow:[GET, POST, DELETE, OPTIONS]]

...

Directory discovery completed in 5.72 seconds
```

## Notes:

- You can add any custom HTTP methods to the ```methods``` array in the source code if needed.
- Adjust the ```maxConcurrency``` value to control the number of simultaneous requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
