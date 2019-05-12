# Aggregate Messaging Controller API


### Index

- [Installing](https://github.com/gamkittisak/kafkaMessageQ-API#installing)
- [Run](https://github.com/gamkittisak/kafkaMessageQ-API#run)
- [Quick start](https://github.com/gamkittisak/kafkaMessageQ-API#quick-start)
- [Document](https://github.com/gamkittisak/kafkaMessageQ-API#document)
    - [Producer](https://github.com/gamkittisak/kafkaMessageQ-API#producer)
        - [aWait: true | false\<boolean>](https://github.com/gamkittisak/kafkaMessageQ-API#await)
        - [message: { JSON [,"cliendID": uuid\<string>] }](https://github.com/gamkittisak/kafkaMessageQ-API#message)
        - [topics: {"sending": string [,"receiving": string] }](https://github.com/gamkittisak/kafkaMessageQ-API#topics)
        - [options](https://github.com/gamkittisak/kafkaMessageQ-API#options)
            - [timeouts](https://github.com/gamkittisak/kafkaMessageQ-API#timeouts)
                -   [timeout: \<int>](https://github.com/gamkittisak/kafkaMessageQ-API#timeout)
                -   [timeoutProduce: \<int>](https://github.com/gamkittisak/kafkaMessageQ-API#timeoutproduce)
                -   [timeoutConsume: \<int>](https://github.com/gamkittisak/kafkaMessageQ-API#timeoutconsume)
            - [clientID: \<string>](https://github.com/gamkittisak/kafkaMessageQ-API#clientid)
        - [examples](https://github.com/gamkittisak/kafkaMessageQ-API#examples)
            - [asynchronous](https://github.com/gamkittisak/kafkaMessageQ-API#asynchronous)
            - [synchronous](https://github.com/gamkittisak/kafkaMessageQ-API#synchronous)
            - [timeouts](https://github.com/gamkittisak/kafkaMessageQ-API#timeouts)
    - [Consumer](https://github.com/gamkittisak/kafkaMessageQ-API#consumer-1)
        - [groupID: \<string>](https://github.com/gamkittisak/kafkaMessageQ-API#groupid)
        - [topics: list\<string | null>](https://github.com/gamkittisak/kafkaMessageQ-API#topics-1)
        - [option](https://github.com/gamkittisak/kafkaMessageQ-API#option)
            - [autoOffsetReset: "earliest" | "latest"\<string>](https://github.com/gamkittisak/kafkaMessageQ-API#autooffsetreset)
        - [examples](https://github.com/gamkittisak/kafkaMessageQ-API#examples-1)
            - [subscribe multi-topics](https://github.com/gamkittisak/kafkaMessageQ-API#subscribe-multi-topics)
            - [subscribe all topics](https://github.com/gamkittisak/kafkaMessageQ-API#subscribe-all-topics)
         

---

### Installing 

---

##### macOS
    brew install librdkafka


    
##### Ubuntu
    apt install librdkafka-dev

    
#####  centos7+/redhat7 +
    yum install  http://download-ib01.fedoraproject.org/pub/epel/7/x86_64/Packages/l/librdkafka-0.11.5-1.el7.x86_64.rpm
    yum install  http://download-ib01.fedoraproject.org/pub/epel/7/x86_64/Packages/l/librdkafka-devel-0.11.5-1.el7.x86_64.rpm

    
#### Environment Variable

    export GOPATH=[path to goDir workspace]
    export AMCO_HOME=$GOPATH/src/AMCO

---

### Run   

#### AMCO Application

    go run main.go


---
## Quick start

##### Producer Asynchronous
Note: Asynchronous in AMCO meaning producer produces the messages to brokers by \
using synchronous send message of the kafka client library, AMCO only using \
synchronous of the kafka client library.


    curl localhost:7890/api/producer -d '{ "aWait":false, "topics":{"sending": "test-req"}, "message":{"greet":"Hello World"} }'

    
##### Consumer 
    
    curl localhost:7890/api/consumer -d '{ "groupID":"test", "autoOffsetReset":"earliest" , "topics":["test-req"]}'
    

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

    curl -XPOST localhost:7890/api/producer  \ 
    -d '{"aWait":false, "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"}}'

##### synchronous

    curl -XPOST localhost:7890/api/producer  \ 
    -d '{"aWait":true, "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]","receiving":"[RECEIVE_FROM_TOPIC]"}'
    
##### timeouts
    curl -XPOST localhost:7890/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"},"timeout": 10}'

    curl -XPOST localhost:7890/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]"},"timeoutProduce": 30}'
    
    curl -XPOST localhost:7890/api/producer  \ 
    -d '{"aWait":[Boolean], "message":{[JSON]}, \ 
    "topics":{"sending":"[SEND_TO_TOPIC]", \
    "receiving":"[RECEIVE_TOPIC_FROM]"},"timeoutProduce": 10, \
    "timeoutConsume": 30,"timeout": 10}'
    
##### clientID
    
    curl -XPOST localhost:7890/api/producer  \ 
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
 
 FLAGS : earliest, latest by default
 
 earliest: automatically reset the offset to the earliest offset or  oldest 
 messages
    
 latest: automatically reset the offset to the latest offset or newest messages

#### Examples




###### subscribe multi-topics

    curl -XPOST localhost:7890/api/consumer \ 
    -d '{"topics":["topice1","topic2","topic3"], "groupID":"GROUP_NAME","autoOffsetReset":"[FLAGS]"}' 

###### subscribe all topics

    NOTE THAT : you must create your topics name  as follows  \
        the first letter that can be an upper case or lower case after first letter \
        can be digit or dash character or letters \

    CASE SUPPORT :
        symbol :
            + <== one or more
            * <== empty or more
        
        [upper case|lower case]+[dash character | digit]*
        
          

    curl -XPOST localhost:7890/api/consumer \ 
    -d '{"topics":[], "groupID":"GROUP_NAME","autoOffsetReset":"[FLAGS]"}' 



