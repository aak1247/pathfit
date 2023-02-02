# Pathfit

| [English](README.md) | 简体中文 |

命令行路径转换工具，将其他环境下的绝对路径转换为当前环境下的绝对路径。

- 支持WSL、Linux、Windows、MacOS等环境，尤其是可以在wsl下使用windows中的路径

## 安装

### 通过二进制文件安装

直接下载对应平台的二进制文件，放到环境变量中的目录下（PATH）即可。

### 通过Golang安装

```bash
go install github.com/aak1247/pathfit@latest
```

## 使用

```bash
pathfit /d/Users/aak1247/Repos
# will output: D:\\Users\aak1247\Repos in Windows
```