## omnibuf

state: **pre-alpha-unstable**


omnibuf is a serialization format that will be the hearth of omniql


1. it is implemented over flatbuffers
2. support same flatbuffers types + a Resource type [is some sort of table with batteries, intended to represent the resources of your database or backend]
2. it does not support namespaces, instead all the definition are organized by application version. 
3. schemas are written using yaml files.


the initial implementation is very inefficient and bloated (this is intended, we will improve it with time when we have a good benchmark procedure) , it consist in store all the binaries in an associative array, where the key is the `file path/tree path` of such binary,the path allow us to know how to decode such binary due tha the last two element always come in a pair of the type [kind, id], when a binary need to do a reference to other it will use the full path for it,  this add a lot of unnecessary bytes for transport the binaries, but give us a other features as tradeoff

examples


```go


const (Worker=0, Tool=1)

type Worker struct{
    name string
}
type Tool struct{
   creator Worker
   owner Worker 
}

({key: [0,1], value: {name: "arnold"}}, {key: [0,1,1,1], value: {creator: [0,1], owner:[0,1]}}]

```



 
