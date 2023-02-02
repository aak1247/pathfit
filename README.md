# Pathfit

| English | [简体中文](README.zh-CN.md) |

Cli tool to translate path to fit current environment.

- Supports WSL、Linux、Windows、MacOS etc，Especially can use windows path in wsl.

## Install

### Install by binary

Download binary file from [release page](https://github.com/aak1247/pathfit/releases), and put it into a directory in PATH.

### Install by Golang

```bash
go install github.com/aak1247/pathfit@latest
```

## Usage

```bash
pathfit /d/Users/aak1247/Repos
# will output: D:\\Users\aak1247\Repos in Windows
```