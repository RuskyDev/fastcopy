# FastCopy
A fast copy tool written in Go.

## Features
- Parallel Copying
  - Dynamically adjusts the number of parallel workers according to CPU core availability.
  - Spawns multiple goroutines to optimize file copying performance using concurrency.
- Cross-Platform
  - Works seamlessly across Windows, Linux, and macOS.
- Large File Support
  - Optimized for copying large file using buffered I/O and file pre-allocation.

## Installation
1. Download the latest release (FastCopy.zip) from the Releases tab.
2. Unzip the downloaded file.
3. Navigate to the folder corresponding to your operating system:
   - **Windows**: Go to the `windows` folder and use `FastCopy.exe`.
   - **Linux**: Go to the `linux` folder and use `FastCopy`.
   - **macOS**: Go to the `macos` folder and use `FastCopy`.
4. On **Linux/macOS**, make the binary executable by running the following command:
   ```bash
   chmod +x FastCopy
## Usage
```
fastcopy <source> <destination> [--quiet]
```

```
<source>: Path to the source file.

<destination>: Path to the destination folder.

Flags:
--quiet (optional): Suppresses output messages.
```

## Donate
If you like this project, you can donate me via paypal<br>

<span class="paypal"><a href="https://www.paypal.me/ruskydev"><img src="https://www.paypalobjects.com/webstatic/mktg/Logo/pp-logo-100px.png"/></a></span>
