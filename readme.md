<p align="center">
  <img src="./docs/logo2.png">
</p>

# GoCrastination Proxy

> A proxy to help you control over your procrastination habits

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)

## Introduction

I got some complains from people about the anti-procrastination apps, that are easily bypassed. So I thought, what about a proxy that blocks the websites you want to avoid?
It's still not perfect, but it's a bit harder to uninstall or just turn off.

## Features

- Configurable focus time periods
- Customizable list of blocked websites
- Simple YAML configuration
- It's written in Go #lol :sunglasses:

## Installation

To install GoCrastination Proxy, you can just clone this repo and run `go build` in the root directory.

```sh
git clone https://github.com/Pliavi/gocrastination.git
cd gocrastination
go build
```

> I promise I'll make a release soon, please have mercy on me :sob:

## Usage

It's simple as running the binary.

```
./gocrastination --port 62222 --config gocrast.yaml
```

- The `--port` flag is optional, it defaults to `62222`.
- The `--config` flag is "required", it will not enforce you to add, but it will not work without it, but I'll add a default config file soon, for now, you can use the `gocrast.yaml` from the `examples/` folder

## Configuration

The proxy is configured using a YAML file. You can create it wherever you want, and just point the `--config` flag to it.

Here's an example configuration:

```yaml
everyday:
  start: 480 # 8h == 8 * 60 = 480
  end: 1320 # 18h == 18*60 = 1080
  sites:
    - twitter.com
    - facebook.com
    - instagram.com
    - reddit.com
```

- `everyday` is the only configuration for now, so the proxy will work for all days in the week
- The time for `start` and `end` is minute based
- You can put any site you want to block in the `sites` list

## Contributing

Contributions are welcome! If you have any ideas for improvement, feature requests, or bug reports, please open an issue on GitHub or submit a pull request.

## Roadmap

- [ ] **Site Blocking Configuration:**

  - [x] Implement the ability to block access to specific sites using a YAML-defined list of URLs.
  - [ ] Allow configuration of blocking schedules for different days of the week.

- [ ] **Redirection Personalization:**

  - [ ] **SSL Certificate Issue Resolution:**
    - [ ] Investigate solutions to bypass the invalid certificate issue when using the Man-in-the-Middle method for redirection.
    - [ ] Explore viable alternatives that ensure connection security without compromising proxy functionality.
  - [ ] Develop the functionality to redirect blocked sites to other designated sites.
  - [ ] Enable the creation of a list of alternative redirects for each blocked site, with the option for random selection.

- [ ] **Configuration Improvements:**

  - [ ] Simplify the setup process for users, eliminating the need to manually point to configuration files unless preferred.

- [ ] **Testing:**

  - Improve the testing of the proxy, including unit tests for the core functionality.

- [ ] **Documentation:**
  - Prepare comprehensive documentation describing the installation, configuration, and usage of the reverse proxy.
