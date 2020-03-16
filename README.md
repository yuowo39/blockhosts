# blockhosts
The simple program can makes a block hosts file.
It supports IPv4 and IPv6. It can makes:

* Hosts file for Windows, macOS, Linux.
* Hosts file for Dnsmasq.

## Input file format
Each host on one line, for example:

    p1-dy.bytecdn.cn
    p2-dy.bytecdn.cn
    p3-dy.bytecdn.cn
    p4-dy.bytecdn.cn
    p5-dy.bytecdn.cn

## How to use

    git clone https://github.com/yuowo39/blockhosts.git
    cd blockhosts/
    go build blockhosts.go
    ./blockhosts tiktok-hosts.txt

## Arguments

    -f       Default: IPv4
    -s       IPv6
    -h       Default: Hosts(Windows, macOS, Linux)
    -d       Dnsmasq
    FilePath
