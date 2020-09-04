# git-credential-1password

[![license](https://img.shields.io/github/license/develerik/git-credential-1password.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/develerik/git-credential-1password?style=flat-square)](https://goreportcard.com/report/github.com/develerik/git-credential-1password)

Helper to store git credentials inside 1password.

## Table of Contents

- [Install](#install)
  - [Dependencies](#dependencies)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## Install

**Note**: Currently only installation from source is supported, so you must have `go` and `make` installed.

```shell script
git clone https://github.com/develerik/git-credential-1password.git
cd git-credential-1password
make git-credential-1password
```

Move the built binary (inside the `bin` directory) to somewhere in your PATH.

### Dependencies

To use this helper you need to install the 1password cli tool ([download](https://support.1password.com/command-line-getting-started/#set-up-the-command-line-tool))
and of course git.  
You also need to setup the cli tool with your 1password account ([guide](https://support.1password.com/command-line-getting-started/#get-started-with-the-command-line-tool)). 

## Usage

```shell script
git config --global credential.helper '!git-credential-1password'
```

## Maintainers

- **Erik Bender** - *Initial work* - [develerik](https://github.com/develerik)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct.

## License

Distributed under the ISC License. See [LICENSE](LICENSE) for more information.

## Acknowledgements

- [1Password](https://1password.com) for their awesome [cli tool](https://1password.com/downloads/command-line)
- [Steve (acahir)](https://github.com/acahir) for his [python implementation](https://github.com/acahir/git-credential-1password)
of a 1password credential helper which inspired me to create this project
- [Netlify](https://www.netlify.com) for their [netlify credential helper](https://github.com/netlify/netlify-credential-helper)
implemented in Go which helped me a lot on my own implementation
