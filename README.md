# mico

>mico: [micho] The person who has just started working in the deck sections of the ships is called miço.

![mascot](img/mascot.png)

`mico` is an AI assisted `kubectl` helper designed to simplify and automate Kubernetes cluster management tasks. It uses the power of OpenAI's GPT-3.5 architecture to understand natural language queries and generate `kubectl` commands that can be used to perform various tasks on your Kubernetes cluster.
## Installation
You can install `mico` using the following command:
```
go install github.com/tahtaciburak/mico@latest
```

## Configure
Before using mico, you need to configure it with your OpenAI API Key. This can be done by running the following command:
```
mico configure
```

## Usage
To use mico, simply provide a natural language query as a parameter to the `mico` command. mico will generate the corresponding `kubectl` command that can be used to perform the desired task on your Kubernetes cluster.

For example, to get the version label of all pods with label `app=cassandra`, you can use the following command:
```bash
mico -p "Get the version label of all pods with label app=cassandra"
```
This will generate the following kubectl command:
```bash
kubectl get pods -l app=cassandra -o jsonpath='{.items[*].metadata.labels.version'
```

You can also pipe your results to directly execute.

```bash
mico -p "Get restart counts for each pod in production namespace" | zsh
```

This will generate the output of `kubectl`:
```bash
plugin-api-5bff86cb56-524hg     0
plugin-api-5bff86cb56-56nf5     0
plugin-api-5bff86cb56-ddnpk     0
plugin-api-5bff86cb56-dll54     0
plugin-api-5bff86cb56-dm2nx     0
plugin-api-5bff86cb56-fmjbh     0
plugin-api-5bff86cb56-ftslf     0
plugin-api-5bff86cb56-sc8g2     0
```


