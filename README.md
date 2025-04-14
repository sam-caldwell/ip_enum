ip_enum
========

This enumerates a list of ip addresses from a given CIDR block.

## Usage

`make all` - build the program

`./build/ip_enum <cidr block>` - enumerate the list of ip addresses from CIDR.

Example:
```shell
% ./ip_enum 192.168.1.0/29 
192.168.1.1
192.168.1.2
192.168.1.3
192.168.1.4
192.168.1.5
192.168.1.6
```

This program can be used to feed other things.  Small, clean and modular was good enough for Dennis Ritchie, it's
good enough for me.