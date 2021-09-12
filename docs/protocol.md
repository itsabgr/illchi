
## Federated Message Broker
---
#### Messaging is part of every app nova days
Most of the current solutions are bloated with features that make maintaining them complicated.
We implemented a simple but powerful peer to peer message broken that is inspired by email architecture.
The email has successful that implemented many years ago but effective still.

**Goals:**
- Cross-platform (and browser support) with no plugin.
- Use standard protocols and do not invent  new ones as is possible
- Fully distributed with no Single Point Of Failure (`SPOF`)
- Equal and same nodes responsibility to decrease maintenance complexity
- Security and authorization at heart
- Flexible authentication method
- Keep It So Simple (`KISS`) and understandable architecture
- Federated architecture like email
- Firewall and (reverse) proxy friendly

In this solution, nodes can send binary messages to each other after the authentication and authorization of each one separately.
Optionally others can listen and receive messages after authentication.

This system addressing Each node with two parameters: 
- the **broker address** (IP:port)
*used for making the net connection to the broker by IP address and port number*
- and a numeric **ID**.
*that is used to identify the target node connection and is a decimal between 1 and 2^31-1*

**Send**

for sending, each node uses target node broker and target node id to make an HTTP `POST` call:

`https://{target broker}/{target ID}`

then returns 204 status without a body if was successful
In case of error, other status codes with a body that is the error description.

**Receive**

for receiving, each node should connect to itself broker and makes a `GET` request

`https://{broker address}/{desired ID}`

At first, the broker authenticates the request and then upgrades that to a Web-Socket connection thence upgrade-related headers are required.
In case of failure, returns a non-successful status code with a body that describes that error.

##### *more features like multi-target send is on the way*.

*notes*

> each node should send the id listening to

> it's clear that the sender must post the message to the broker that the target is listening to why there is no between brokers message routing.
 
> authentication parameters can send as query strings in HTTP calls to pass them to the authorization module.

> handling Cross-Origin Resource Sharing (`CORS`) over `OPTIONS` requests is configurable

> `anycast DNS` breaks the architecture but can use  `DNS` for service discovery purpose

> to get broker statics should make a get call with no path: 
> `GET http://{broker address}/`

> `IPv6` is  supported

###### ABGR (5 Sep 2021)
