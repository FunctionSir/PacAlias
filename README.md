<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2023-10-27 17:43:29
 * @LastEditTime: 2024-01-11 23:38:28
 * @LastEditors: FunctionSir
 * @Description: README
 * @FilePath: /PacAlias/README.md
-->
# PacAlias

Current Version: 0.2.0-beta (KannaKamui)  
Give pacman repo a alias.  

## How to make and/or install

This repo has a bin version of the program inside.  
If you are using an amd64 platform, just:  

```shell
sudo bash install.sh
```

You can build yourself, just:  

```shell
make
```

Make and install is easy too, just:  

```shell
sudo make install
```

## How to use

Here is a example:  
Background: Using Arch Linux.  

```ini
# At pacalis.conf
[manjaro-core]
Dest = https://mirrors.tuna.tsinghua.edu.cn/manjaro/unstable/core/$arch
OName = core
[manjaro-extra]
Dest = https://mirrors.tuna.tsinghua.edu.cn/manjaro/unstable/extra/$arch
OName = extra
```

```ini
# At pacman.conf
[manjaro-core]
Server = http://127.0.0.1:2090/?repo=$repo&arch=$arch&file=
[manjaro-extra]
Server = http://127.0.0.1:2090/?repo=$repo&arch=$arch&file=
```

Then you can use the packages in repo core or extra of manjaro.  

## Args

-c/--conf: Conf file.  
-p/--port: Http port.  
-a/--addr: Addr to listen.  
-n/--notime: Do NOT add time info to the output.  
-q/--quite: No println.  
