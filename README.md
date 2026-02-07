# spongecase

A small (slightly pointless) command line utility for converting text to
[Spongecase](https://en.wikipedia.org/wiki/Alternating_caps)

## Usage

```bash
spongecase text to convert

# This will return:
# tExT tO cOnVeRt
```

## Installation

Simply fetch the binary from the GitHub release, make the binary executable,
and add it to your path

E.g.

For Linux on x86:

```bash
curl -L -o spc \
  "https://github.com/aliktb/spongecase/releases/latest/download/spongecase_linux_amd64"
```

Or for ARM-based Macs:

```bash
curl -L -o spc \
  "https://github.com/aliktb/spongecase/releases/latest/download/spongecase_darwin_arm64"
```

> [!TIP]
> Note, the binary is being downloaded as `spc` to make it easier than using
> `spongecase` in the terminal
>
> This would make the usage as:
>
> ```bash
> spc text to convert
> ```

Next, simply make the binary executable and move it to your path. E.g.

```bash
chmod +x ./spc

# Create bin directory within home directory
mkdir -p ~/.local/bin

# Add this to .bashrc or .zshrc
# This will add the local bin directory to the system path
# for extra info, see https://opensource.com/article/17/6/set-path-linux
export PATH="~/.local/bin:$PATH"
```
