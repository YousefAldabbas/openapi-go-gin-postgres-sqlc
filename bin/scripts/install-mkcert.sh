#!/bin/bash

VERSION="1.4.3"
sudo apt install libnss3-tools
wget https://github.com/FiloSottile/mkcert/releases/download/v"$VERSION"/mkcert-v"$VERSION"-linux-amd64 -O mkcert
chmod +x mkcert
sudo mv mkcert /usr/bin/
source ~/.bashrc
cd certificates || exit
mkcert --cert-file localhost.pem --key-file localhost-key.pem localhost 127.0.0.1 ::1
cd ..
mkcert -install
