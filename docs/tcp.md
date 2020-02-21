# TCP 

Stands for transmission control protocol and the proper way for a network connection to be established between two computers on a network.

## TCP Handshake
There are three scenarios open, closed, and filtered port.
### Open Port
If open then a three way hand  shake happens.

client ---syn---->  server
client <--syn-ack-- server
server ---ack-----> server

### Closed Port
If closed then rst packet is received.

client ---syn----> server
client <---rst---- server


### Filtered Port
If request passes through a firewall then there is generally no response to the client.

client ---sync---> FIREWALL --X server

## Bypassing Firewalls with Port Forwarding

Port forwarding := Using an intermediary step to bypass a firewall. 
The client establishes a connection with a non-blocked resource and then hops to the blocked resource from the connection.


client ---requests stackTitan.com ---> FIREWALL ---Requests passes firewall---> stackTitan.com ---traffic proxied to evil.com---> evil.com

### About Firewalls at the Different Networking Layers ^1

**Layer 3 firewalls** - Known as packet filtering firewalls will filter traffic based solely on **source/destination IP**, **port**, and **protocol**.

**Layer 4 firewalls** - Do the same as the above, plus add the ability to track active network connections, and allow/deny traffic based on the state of those sessions known as stateful packet inspection.

**Layer 7 firewalls** - Known as application gateways can do all of the above, plus include the ability to intelligently inspect the contents of those network packets. For instance, a Layer 7 firewall could deny all HTTP POST requests from Chinese IP addresses. This level of granularity comes at a performance cost, though.

#### References
[1. ServerFault](https://serverfault.com/questions/792572/what-does-a-layer-3-4-firewall-do-that-a-layer-7-does-not)


## TCP Proxy

Request goes to an open port at a trusted website and is then forwarded to the untrusted website