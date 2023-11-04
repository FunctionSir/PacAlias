#!/bin/bash
###
 # Author: FunctionSir
 # License: AGPLv3
 # Date: 2023-11-04 14:23:27
 # LastEditTime: 2023-11-04 14:23:38
 # LastEditors: FunctionSir
 # Description: Installer for pacalias.
 # FilePath: /PacAlias/install.sh
###
echo 'Installing pacalias...''
sudo cp pacalias /usr/bin
sudo chmod 755 /usr/bin/pacalias
sudo cp pacalias.conf /etc
sudo cp pacalias.service /etc/systemd/system
echo 'Done! Use "systemctl enable --now pacalias" to start and enable it.'