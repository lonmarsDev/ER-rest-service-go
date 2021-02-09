# Log package
Package Log to be used by the data sources ingestion services to be used by any service that would like to schedule periodic pulling task to update the data from the data sources.

## How to use
Follow this example to know how to use this package
```go

	// set the log to console
	log.SetLogger(log.CONSOLE)

	// use the log interface to print debug
	log.Debug("Module1", "hello debug")

	// use the log interface to print Info
	log.Info("Module2", "hello info")

	// use the log interface to print warning
	log.Warning("Module3", "hello warning")

	// use the log interface to print Error
	log.Error("Module3", "hello Error")
```

## Log Types
To change the logging output interface please follow the follwing
- ### Console Log
```go
    // set log to console
	log.SetLogger(log.CONSOLE)
```
- ### File Log
```go
    // set log to file
    log.SetLogger(log.FILE)
```

**Note** you should add the following environment variables in order to make this work
```
LOGFILE_PATH
LOGFILE_NAME
```
- ### Elastic Log
```go
    // set log to Elastic search
	log.SetLogger(log.ELASTIC)
```

**Note** you should add the following environment variables in order to make this work
```
ELASTIC_URL
```