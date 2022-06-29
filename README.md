# Sysdig Query Translator

## What it is?
This is a tool created in GO that creates a C library to be used mainly by a Pyhton application. With this tool you can translate a sysdig promql query to a vanilla Prometheus query.

## How to use it?
You can import the library in your Python code and use it. The function RemoveSysdigLabelsInExpression expects two string as input:
- The expresion itself
- The extra scope you want to remove

In the folder example you will find an example of how to integrate it in your Pyhton code.

## Build
If you want to build it by yourself just execute the build command of the makefile:

```bash
make build
```
This will create a folder library with the two needed files with the C library.