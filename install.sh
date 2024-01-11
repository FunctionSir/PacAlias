#!/bin/bash
###
 # @Author: FunctionSir
 # @License: AGPLv3
 # @Date: 2023-11-04 14:23:27
 # @LastEditTime: 2024-01-11 23:52:04
 # @LastEditors: FunctionSir
 # @Description: -
 # @FilePath: /PacAlias/install.sh
###
echo 'Installing pacalias...'
sudo cp pacalias /usr/bin
sudo chmod 755 /usr/bin/pacalias
sudo cp -n pacalias.conf /etc
sudo cp pacalias.service /etc/systemd/system
echo 'Done! Use "systemctl enable --now pacalias" to start and enable it.'