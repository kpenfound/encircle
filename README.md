# encircle

Run a circleci yaml locally with Dagger

## Using encircle

- `go build` to build the `encircle` binary
- `./encircle` will execute the yaml in `./circleci/config.yml`

Notes:

- Right now it's hardcoded to a specific workflow in the yaml
- Doesn't understand non-`run` steps like `checkout` or orbs
- Still a lot to go on the config

Sample output:

```
go build && ./encircle
Loading config at ./.circleci/config.yml
Running workflow build_test
#1 resolve image config for docker.io/library/node:16
#1 DONE 2.5s

#2 Install npm dependencies
#2 DONE 0.0s

#3 host.directory /Users/kylepenfound/github.com/encircle
#3 transferring /Users/kylepenfound/github.com/encircle: 29B
#3 ...

#4 from node:16
#4 resolve docker.io/library/node:16 0.3s done
#4 DONE 0.3s

#3 host.directory /Users/kylepenfound/github.com/encircle
#3 transferring /Users/kylepenfound/github.com/encircle: 9.18MB 0.5s done
#3 DONE 0.5s

#2 Install npm dependencies
#2 0.119 npm install
#2 0.119
#2 DONE 0.2s

#5 Run Unit Tests
#0 0.090 npm test
#0 0.090
#5 DONE 0.1s

#6 resolve image config for docker.io/library/golang:latest
#6 DONE 2.2s

#7 from golang:latest
#7 resolve docker.io/library/golang:latest
#7 resolve docker.io/library/golang:latest 0.3s done
#7 DONE 0.3s

#7 from golang:latest
#7 sha256:77fe3ac745a5ff347cfde07c7e36e684d71ada78b8efa592ba32dcd423a2ac32 0B / 155B 0.2s
...
#7 extracting sha256:77fe3ac745a5ff347cfde07c7e36e684d71ada78b8efa592ba32dcd423a2ac32 done
#7 DONE 25.5s

#8 Run Go Tests
#8 0.105 go test
#8 0.105
#8 DONE 0.4s

#9 Run Go Build
#0 0.104 go build
#0 0.104
```
