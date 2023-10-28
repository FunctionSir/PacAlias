<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2023-10-27 17:43:29
 * @LastEditTime: 2023-10-28 17:03:53
 * @LastEditors: FunctionSir
 * @Description: README
 * @FilePath: /PacAlias/README.md
-->
# PacAlias

Give pacman repo a alias.

## How to use

There is a example:  
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
-1/--quite: No println.  
