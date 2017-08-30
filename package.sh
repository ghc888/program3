#!/bin/bash
echo "# Getting branch info"
git status -bs
echo "# Press Return or Space to start build, all other keys to quit"
read -s -n 1 key
if [[ $key != "" ]]; then exit; fi
version=$(git describe --tag)
head=$(git rev-parse --short HEAD)
epoch=$(date +%s)
echo "# Building"
./build.sh
echo "# Cleaning up previous builds"
rm -rf build
rm *.tar.gz
rm *.deb
rm *.rpm
mkdir -p build/usr/bin
mkdir -p build/usr/share/program3/dashboard
mkdir -p build/etc/program3
mkdir -p build/etc/systemd/system
mkdir -p build/etc/init.d
mkdir -p build/var/lib/program3
echo "# Copying files to build dir"
cp program3 build/usr/bin/
#cp etc/config.toml.sample build/etc/replication-manager/config.toml.sample
#cp -r dashboard/* build/usr/share/replication-manager/dashboard/
#cp -r share/* build/usr/share/replication-manager/
#cp service/replication-manager.service build/etc/systemd/system
#cp service/replication-manager.init.el6 build/etc/init.d/replication-manager
#cp service/replication-manager-arbitrator.init.el6 build/etc/init.d/replication-manager-arbitrator
echo "# Building packages"
fpm --epoch $epoch --iteration $head -v $version -C build -s dir -t rpm -n program3 .
fpm --package program3-$version-$head.tar -C build -s dir -t tar -n program3 .
gzip program3-$version-$head.tar
#cp service/replication-manager.init.deb7 build/etc/init.d/program3
fpm --epoch $epoch --iteration $head -v $version -C build -s dir -t deb -n program3 .
