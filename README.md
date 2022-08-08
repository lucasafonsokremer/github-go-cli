# github-go-cli

- github-go-cli provides an executable called github-go-cli, that can be used to access all of GitHub’s public API functionality from your command-line

- github-go-cli is written on top of [cobra library](https://github.com/spf13/cobra). Cobra is used in many Go projects such as Kubernetes, Hugo, and GitHub CLI to name a few.

# Overview

- github-go-cli works based on that structure:

![](https://miro.medium.com/max/1328/0*T57SgaZDr9yPicpG)

# Instalation

### Manual

Go to the [releases](https://github.com/lucasafonsokremer/github-go-cli/releases) page and download the latest one for your platform.

# Usage

```
github-go-cli [command] [--flags]
```

- Get user informations:

```
github-go-cli describe user brendangregg
```

output:

```
Name              Location
Brendan Gregg     Sydney, Australia
```

- Get user followers:

```
github-go-cli list followers brendangregg
```

output:

```
Name               Github Link
joaovitor          https://github.com/joaovitor
marcinpohl         https://github.com/marcinpohl
dacresni           https://github.com/dacresni
toke               https://github.com/toke
nyxwulf            https://github.com/nyxwulf
sanga              https://github.com/sanga
jacques            https://github.com/jacques
nicoulaj           https://github.com/nicoulaj
adulau             https://github.com/adulau
adragomir          https://github.com/adragomir
mrluanma           https://github.com/mrluanma
MiloCasagrande     https://github.com/MiloCasagrande
JonGretar          https://github.com/JonGretar
mburns             https://github.com/mburns
yrcjaya            https://github.com/yrcjaya
zdk                https://github.com/zdk
jailton            https://github.com/jailton
openweb            https://github.com/openweb
bartman            https://github.com/bartman
narkisr            https://github.com/narkisr
gburd              https://github.com/gburd
evanfarrar         https://github.com/evanfarrar
jf                 https://github.com/jf
jage               https://github.com/jage
huangliang2211     https://github.com/huangliang2211
filipeamoreira     https://github.com/filipeamoreira
igungor            https://github.com/igungor
johnm              https://github.com/johnm
meangrape          https://github.com/meangrape
liwei              https://github.com/liwei
```

# Get Help

```
github-go-cli --help
```

output:

```
Access all of GitHub’s public API functionality from your command-line

Usage:
  github-go-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  describe    Retrieve detailed information about any object
  help        Help about any command
  list        Retrieve user profile information such as name, location, public repos, followers, following

Flags:
      --config string   config file (default is $HOME/.github-go-cli.yaml)
  -h, --help            help for github-go-cli
  -t, --toggle          Help message for toggle

Use "github-go-cli [command] --help" for more information about a command.
```

# ToDo

- Improve tests
- CLI auto-completion
- New functions
