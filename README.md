# Orionctl

This is a command line tool for calling FIWARE Orion API.

## Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Overview

An orionctl provides:

- Create, Get Retrieve and Delete Orion Subscription
- Use resources defined in JSON and YAML

## Installing

Using Orionctl is easy.

### Download binary

Download the binary from the [release](https://github.com/YujiAzama/orionctl/releases)

```bash
curl -OL https://github.com/YujiAzama/orionctl/releases/download/v0.1.0/orionctl_linux_x86_64.tar.gz
tar zxvf orionctl_linux_x86_64.tar.gz
chmod +x ./orionctl
sudo mv ./orionctl /usr/local/bin/orionctl
```

### Build from source

```bash
git clone https://github.com/YujiAzama/orionctl && cd orionctl
go build main.go -o orionctl
chmod +x ./orionctl
sudo mv ./orionctl /usr/local/bin/orionctl
```

## Configurations

Orionctl configurations are given priority in the order of flags, environment variables, and configuration file(.orionctl.yaml).

### Configurations by flags

```bash
orionctl get subscription -H orion -p 1027
```

### Configurations by environment variables

```bash
HOST=orion PORT=1027 orionctl get subscription
```

### Configurations by `.orionctl.yaml`.

Configuration file `.orionctl.yaml` is placed in `$HOME` by default.

```yaml:.orionctl.yaml
Host: "orion"
Port: 1027
TLS: false
Token: ""
```

## Getting Started

You can get help by running it with the -h option

```bash
$ orionctl -h
This is a command line interface for control FIWARE Orion.

Usage:
  orionctl [command]

Available Commands:
  create      A brief description of your command
  delete      A brief description of your command
  describe    Describe Orion resources
  get         Get Orion resources
  help        Help about any command
  metrics     Get Orion metrics

Flags:
      --config string               config file (default is $HOME/.orionctl.yaml)
  -s, --fiware-service string       FIWARE Service
  -P, --fiware-servicepath string   FIWARE Service Path
  -h, --help                        help for orionctl
  -H, --host string                 Orion hostname or IP address (default "localhost")
  -p, --port int                    Orion port number (default 1026)
  -t, --toggle                      Help message for toggle

Use "orionctl [command] --help" for more information about a command.
```

First, create a subscription definition file. Generally, all Orion resources are represented by JSON.
An example of JSON definition is as follows:

```json
{
  "description": "A subscription to get info about Room1",
  "subject": {
    "entities": [
      {
        "id": "Room1",
        "type": "Room"
      }
    ],
    "condition": {
      "attrs": [
        "pressure"
      ]
    }
  },
  "notification": {
    "http": {
      "url": "http://localhost:1028/accumulate"
    },
    "attrs": [
      "temperature"
    ]
  },
  "expires": "2040-01-01T14:00:00.00Z",
  "throttling": 5
}
```

Orionctl is a very useful tool and can use YAML definitions.
By writing in YAML, you can manage resources very easily and readable.
An example of YAML definition is as follows:

```yaml
---
description: A subscription to get info about Room1
subject:
  entities:
  - id: Room1
    type: Room
  condition:
    attrs:
    - pressure
notification:
  http:
    url: http://localhost:1028/accumulate
  attrs:
  - temperature
expires: '2040-01-01T14:00:00.00Z'
throttling: 5
```

Create a subscription resource as follows:

```bash
$ orionctl create subscription -f sample.yaml 
subscription "5f301631d9d315f846e98fbf" created
```

Get subscription resources as follows:

```bash
$ orionctl get subscriptions
ID                      	Description                                       
5f1da1d8d9d315f846e98fa6	A subscription to get info about Room1
5f1da1f6d9d315f846e98fa7	A subscription to get info about Room2
5f1ee3bdd9d315f846e98fbd	A subscription to get info about Room3
```

Describe subscription resources as follows:

```bash
$ orionctl describe subscription
ID:           	5f301631d9d315f846e98fbf
Description:  	A subscription to get info about Room1
Subject:
    Entities: 	Id: Room1, Type: Room
    Condition:
        Attrs:	pressure
Notification:
    HTTP:
        URL:  	http://localhost:1028/accumulate
    Attrs:    	temperature
Expires:      	2040-01-01T14:00:00.00Z
Throttling:   	5
              
```

Delete subscription resources as follows:

```bash
$ orionctl delete subscription
subscription "5f301631d9d315f846e98fbf" deleted
```

## Contributing

1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/orionctl)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request

## License

orionclient-go is released under the Apache 2.0 license. See [LICENSE](https://github.com/YujiAzama/orionctl/LICENSE)
