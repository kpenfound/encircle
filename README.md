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
#1 resolve image config for docker.io/circleci/node:12
#1 DONE 0.5s

#2 Install npm dependencies
#2 DONE 0.0s

#3 host.directory /Users/kylepenfound/github.com/encircle
#3 transferring /Users/kylepenfound/github.com/encircle: 29B
#3 ...

#4 from circleci/node:12
#4 resolve docker.io/circleci/node:12 0.3s done
#4 DONE 0.3s

#3 host.directory /Users/kylepenfound/github.com/encircle
#3 transferring /Users/kylepenfound/github.com/encircle: 9.17MB 0.4s done
#3 DONE 0.5s

#2 Install npm dependencies
#2 DONE 0.0s

#4 from circleci/node:12
#4 CACHED

#2 Install npm dependencies
#2 0.152 sudo npm install -g
#2 0.152
#2 DONE 0.2s

#5 Run Unit Tests
```
