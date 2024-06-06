# Packer Plugin for HCI 

This is a [HashiCorp Packer](https://www.packer.io/) plugin for creating  image in a private cloud.

I wrote this plugin for learning.

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    hci = {
      version = ">= 0.0.1"
      source  = "github.com/wangweihong/hci"
    }
  }
}
```

### Install from source

If you prefer to build the plugin from source, clone the GitHub repository
to `$GOPATH/src/github.com/wangweihong/packer-plugin-hci`.

```sh
mkdir -p $GOPATH/src/github.com/wangweihong; cd $GOPATH/src/github.com/wangwiehong
git clone git@github.com:wangweihong/packer-plugin-hci.git
```

Then enter the plugin directory and run `make dev` command to build the plugin.

```sh
cd $GOPATH/src/github.com/wangweihong/packer-plugin-hci
make dev
```

Upon successful compilation, a `packer-plugin-hci` plugin binary file
can be found in the directory. To install the compiled plugin, please follow the
official Packer documentation on [installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).

## Configuration

For more information on how to configure the plugin, please read the
documentation located in the [`docs/`](docs) directory or [`wiki`](https://github.com/huaweicloud/packer-plugin-huaweicloud/wiki).

## [Logging and Debugging](https://developer.hashicorp.com/packer/docs/debugging)

### Debugging Packer in Linux

```shell
$ export HTTPCLI_DEBUG=1
$ export PACKER_LOG=1
$ export PACKER_LOG_PATH="./packer.log"
```

### Debugging Packer in Powershell/Windows

```powershell
$env:HW_DEBUG=1
$env:PACKER_LOG=1
$env:PACKER_LOG_PATH="./packer.log"
```

## Contributing

* If you think you've found a bug in the code or you have a question regarding
  the usage of this software, please reach out to us by opening an issue in
  this GitHub repository.
* Contributions to this project are welcome: if you want to add a feature or a
  fix a bug, please do so by opening a Pull Request in this GitHub repository.
  In case of feature contribution, we kindly ask you to open an issue to
  discuss it beforehand.
  