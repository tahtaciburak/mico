# mico

>mico: [micho] The person who has just started working in the deck sections of the ships is called miço.


An **AI assisted** kubectl helper.

## Installation
```
go get github.com/tahtaciburak/mico
```

## Configure

First you need to configure mico with your OpenAI API Key
```
mico configure
```

## Usage
Command:
```
mico -p "Get the version label of all pods with label app=cassandra"
```
Result:
```
kubectl get pods -l app=cassandra -o jsonpath='{.items[*].metadata.labels.version'
```



