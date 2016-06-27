# lookup
Lookup IPs or Hosts

```shell
~/g/g/s/g/k/lookup ❯❯❯ lookup cname google.com
Name: 	google.com
Canonical Name: 	google.com.

~/g/g/s/g/k/lookup ❯❯❯ lookup host google.com
Name: 	google.com
Address: 	172.217.0.46
Address: 	2607:f8b0:4006:807::200e

~/g/g/s/g/k/lookup ❯❯❯ lookup ip 172.217.0.46
Address: 	172.217.0.46
Name: 	google.com

~/g/g/s/g/k/lookup ❯❯❯ lookup ns google.com
Name: 	google.com
Nameserver: 	ns3.google.com.
Nameserver: 	ns2.google.com.
Nameserver: 	ns1.google.com.
Nameserver: 	ns4.google.com.

~/g/g/s/g/k/lookup ❯❯❯ lookup tcpservice telnet
TCP Service: 	telnet
TCP Port: 	23

~/g/g/s/g/k/lookup ❯❯❯ lookup txt google.com
Name: 	google.com
TXT: 	v=spf1 include:_spf.google.com ~all

~/g/g/s/g/k/lookup ❯❯❯ lookup udpservice telnet
UDP Service: 	telnet
UDP Port: 	23
```
