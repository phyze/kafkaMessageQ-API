# Aggregate Messaging Controller API


### Index

- [Installing](https://gitlab.com/gamkittisak/amco/blob/master/README.md#installing)
- [Run](https://gitlab.com/gamkittisak/amco/blob/master/README.md#run)
- [Quick start](https://gitlab.com/gamkittisak/amco/blob/master/README.md#quick-start)
- [Document](https://gitlab.com/gamkittisak/amco/blob/master/README.md#document)
    - [Producer](https://gitlab.com/gamkittisak/amco/blob/master/README.md#producer)
        - [aWait: true | false\<boolean>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#await)
        - [message: { JSON [,"cliendID": uuid\<string>] }](https://gitlab.com/gamkittisak/amco/blob/master/README.md#message)
        - [topics: {"sending": string [,"receiving": string] }](https://gitlab.com/gamkittisak/amco/blob/master/README.md#topics)
        - [options](https://gitlab.com/gamkittisak/amco/blob/master/README.md#options)
            - [timeouts](https://gitlab.com/gamkittisak/amco/blob/master/README.md#timeouts)
                -   [timeout: \<int>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#timeout)
                -   [timeoutProduce: \<int>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#timeoutproduce)
                -   [timeoutConsume: \<int>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#timeoutconsume)
            - [clientID: \<string>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#clientid)
        - [examples](https://gitlab.com/gamkittisak/amco/blob/master/README.md#examples)
            - [asynchronous](https://gitlab.com/gamkittisak/amco/blob/master/README.md#asynchronous)
            - [synchronous](https://gitlab.com/gamkittisak/amco/blob/master/README.md#synchronous)
            - [timeouts](https://gitlab.com/gamkittisak/amco/blob/master/README.md#timeouts)
    - [Consumer](https://gitlab.com/gamkittisak/amco/blob/master/README.md#consumer-1)
        - [groupID: \<string>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#groupid)
        - [topics: list\<string | null>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#topics-1)
        - [option](https://gitlab.com/gamkittisak/amco/blob/master/README.md#option)
            - [autoOffsetReset: "earliest" | "latest"\<string>](https://gitlab.com/gamkittisak/amco/blob/master/README.md#autooffsetreset)
        - [examples](https://gitlab.com/gamkittisak/amco/blob/master/README.md#examples-1)
            - [subscribe multi-topics](https://gitlab.com/gamkittisak/amco/blob/master/README.md#subscribe-multi-topics)
            - [subscribe all topics](https://gitlab.com/gamkittisak/amco/blob/master/README.md#subscribe-all-topics)
         

---

### Installing 

---

##### macOS
    brew install librdkafka
    brew install redis
    ln -sfv /usr/local/opt/redis/*.plist ~/Library/LaunchAgents
    launchctl load ~/Library/LaunchAgents/homebrew.mxcl.redis.plist

    
##### Ubuntu
    apt install librdkafka-dev
    apt install redis-server
    
#####  centos7+/redhat7 +
    yum install  http://download-ib01.fedoraproject.org/pub/epel/7/x86_64/Packages/l/librdkafka-0.11.5-1.el7.x86_64.rpm
    yum install  http://download-ib01.fedoraproject.org/pub/epel/7/x86_64/Packages/l/librdkafka-devel-0.11.5-1.el7.x86_64.rpm
    yum install epel-release yum-utils
    yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm
    yum-config-manager --enable remi
    yum install redis
    
#### Environment Variable

    export GOPATH=[path to goDir workspace]
    export AMCO_HOME=$GOPATH/src/AMCO

---

### Run 
    
---
    
#### Redis

##### macOS

    redis-server /usr/local/etc/redis.conf
            
##### Ubuntu 

    systemctl start redis
    systemctl enable redis
        
##### centos/redhat

    systemctl start redis
    systemctl enable redis


#### AMCO Application

    go run main.go


---
## Quick start

##### Producer Asynchronous
Note: Asynchronous in AMCO meaning producer produces the messages to brokers by \
using synchronous send message of the kafka client library, AMCO only using \
synchronous of the kafka client library.


    curl localhost:7890/amco/api/producer -d '{ "aWait":false, "topics":{"sending": "test-req"}, "message":{"greet":"Hello World"} }'

    
##### Consumer 
    
    curl localhost:7890/amco/api/consumer -d '{ "groupID":"test", "autoOffsetReset":"earliest" , "topics":["test-req"]}'
    

---
# Document

### Producer

##### aWait
    
if true AMCO will block process until receive result  message, 
false AMCO will not wait receive message from topic that specified.

##### message
    
that is your any data in json format.


##### topics

for producer topics isn't list but it's json format, below is the fields of topics.

-  ##### sending  
    It's the destination topic that you want to send message to.

-  ##### receiveing 
    note: not allowed redeclare name of sending
    Similar on above but it waits to receive message from another topic.


### Options

##### timeouts 
    
only allowed  producer.

-   ##### timeout
    default 30s

-   ##### timeoutConsume
    default 30s 

-   ##### timeoutProduce 
    default 10s 
    
##### clientID

note: only allowed to producer that aWait set to false

clientID is the UUID used to checking where messages come from  \
it may come from services[i] of services[1],services[2],...,services[n] that subscribe at the same topic.


#### Examples

##### asynchronous

    curl -XPOST localhost:7890/amco/api/producer  \ 
    -d '{"aWait":false, "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"}}'

##### synchronous

    curl -XPOST localhost:7890/amco/api/producer  \ 
    -d '{"aWait":true, "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]","receiving":"[RECEIVE_FROM_TOPIC]"}'
    
##### timeouts
    curl -XPOST localhost:7890/amco/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"},"timeout": 10}'

    curl -XPOST localhost:7890/amco/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"},"timeoutProduce": 30}'
    
    curl -XPOST localhost:7890/amco/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]", \
    "receiving":"[RECEIVE_TOPIC_FROM]"},"timeoutProduce": 10, \
    "timeoutConsume": 30,"timeout": 10}'
    
##### clientID
    
    curl -XPOST localhost:7890/amco/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON][, "clientID":"[UUID]" ]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"}}'
    

---



### Consumer

##### groupID

name of service group which using consume message


 
##### topics 

AMCO  allowed to subscribe multiple topics by default if and only if
you let topics list to empty so that will subscribe all topics 

#### option

##### autoOffsetReset
 
 FLAGS : earliest, latest
 
 earliest: automatically reset the offset to the earliest offset or  oldest 
 messages
    
 latest: automatically reset the offset to the latest offset or newest messages

#### Examples




###### subscribe multi-topics

    curl -XPOST localhost:7890/amco/api/consumer \ 
    -d '{"topics":["topice1","topic2","topic3"], "groupID":"GROUP_NAME","autoOffsetReset":"[FLAGS]"}' 

###### subscribe all topics

    curl -XPOST localhost:7890/amco/api/consumer \ 
    -d '{"topics":[], "groupID":"GROUP_NAME","autoOffsetReset":"[FLAGS]"}' 



