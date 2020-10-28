# **Go** Eureka Client

Minimal Eureka Golang client

### Prerequisites

* Go 1.11 and above 


### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
go get github.com/mattb2401/go-eureka-client
```

### Examples

Create a simple application server 
```Go
import (
    "net/http"
    "github.com/mattb2401/go-eureka-client/eureka"
)

func main() {
    client := eureka.NewClient("http://127.0.0.1:8761/eureka")
    instance := client.NewInstance("foo", "localhost", "127.0.0.1", 7022, false, "", "", "")
    // Register instance to Eureka
    err := client.Register(instance)
    if err != nil {
        panic(err)
    }
}
```
## Contributing

Please be nice and Give me a hand

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/mattb2401/veego/tags). 

## Authors

See also the list of [contributors](https://github.com/mattb2401/go-eureka-client/contributors) who participated in this project.


## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc