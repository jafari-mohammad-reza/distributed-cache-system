# Distributed Cache system
try to be a distributed cache system that transfer messages using grpc and a custome message broker which also usess grpc

the cache system implements AOF feature of redis and items will be stored first in leader replica and then the write log will be sent to followers to cache up
a heart beat system will be implemented for caching if nodes are still alive
election happens using an arbiter which vote which node will be assigned as next leader

client sends a request to find out which node is leader to leader topic => only leader node subscribe to that topic and answer the client with its node
for split brains arbiter selects the new leader and inform the client and chosen node so then client sends all write requests to new node and new node subscribe to leader topic
leader topic has a limit which if a new subscriber joins then the previous one gets remove this will prevent split brain

node recovery after crash will happen by AOF and gap recovery happens by sending the gap to leader and fetch log for that period while receiving new logs continuously.