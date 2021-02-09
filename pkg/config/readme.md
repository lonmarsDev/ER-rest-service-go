# LLSG Project Config package
Package Config to be used by the data sources ingestion services to be used by any service that would like to schedule periodic pulling task to update the data from the data sources.

## How to use
### Configurations based on Environment variables
Follow this example to know how to use this package
```go

    // Get Config value from Environment Variables
    // key: is the name of the variable
    // Default Value will be the Value if the variable not found
    config.GetEnvironment().Get("Key", "Default Value")

    // Set Config value from Environment Variables
    // Key: the name of the variables
    // value: the value of the variable
    config.GetEnvironment().Set("Key", "Value")

```

### Configuration based on file
you have to create environment variables to use this type of config:
- ```CONFIG_FILE_PATH``` the path of the file which will contains this configuration, this file will be in plain text so take care of this
```go
    // Get Config value from File
    // key: is the name of the variable
    // Default Value will be the Value if the variable not found
    config.GetFile().Get("Key", "Default Value")

    // Set Config value from File
    // Key: the name of the variables
    // value: the value of the variable
    config.GetFile().Set("Key", "Value")
```


## Developer Guide
You can add another config which could be used in the same service

For example if you want to add a config file that use the Database then you will implement 
```GetDB()``` inside the ```config``` package and this function should return a new **struct**
that implements the ```Config``` **interface** 

you will follow something such as the **env** package.
