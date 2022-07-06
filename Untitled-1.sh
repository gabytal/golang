#!/bin/bash
# usage: ./create_eth0 eth0_IP eth1_IP 
cat >/etc/default/grub  << EOF
GRUB_DEFAULT=0
GRUB_TIMEOUT=0
GRUB_DISTRIBUTOR=`lsb_release -i -s 2> /dev/null || echo Debian`
GRUB_CMDLINE_LINUX_DEFAULT="quiet"
GRUB_CMDLINE_LINUX="net.ifnames=0 biosdevname=0 net.ifnames=0 biosdevname=0"
EOF

grub-mkconfig -o /boot/grub/grub.cfg


cat >/etc/network/interfaces.d/eth0.cfg  << EOF
auto eth0
iface eth0 inet static
    address 192.168.51.$1
    netmask 255.255.255.0
    network 192.168.51.0
    broadcast 192.168.51.255
    gateway 192.168.51.1
    dns-nameservers 172.31.59.174
EOF

cat >/etc/network/interfaces.d/eth1.cfg  << EOF
auto eth1
iface eth1 inet static
    address 192.168.51.$2
    netmask 255.255.255.0
    network 192.168.51.0
    broadcast 192.168.51.255
    gateway 192.168.51.1
    dns-nameservers 172.31.59.174
EOF
