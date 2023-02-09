# encircle

Run a circleci yaml locally powered by Dagger

## Using encircle

- `go build` to build the `encircle` binary
- `./encircle` will execute all workflows in the yaml `./circleci/config.yml`
- `./encircle workflow <workflow name>` will execute the workflow `<workflow name>` in the yaml `./circleci/config.yml`
- `./encircle job <job name>` will execute the job `<job name>` in the yaml `./circleci/config.yml`
- To execute your own `config.yml`, run encircle from your projects root directory

TODO:

- Still a lot to go on the config
- Handle yaml conditionals
- Probably need some tests

Sample output:

```shell
./encircle workflow test_two
loading config at ./.circleci/config.yml
warning: unhandled command when
warning: unhandled command when
warning: unhandled command when
warning: unhandled command when
warning: unhandled command when
warning: unhandled command when
warning: skipping checkout for local dev
running workflow test_two
running job job_two
1 resolve image config for docker.io/library/golang:latest
1 DONE 0.4s

2 Run Go Tests
2 DONE 0.0s

3 host.directory /Users/kylepenfound/github.com/encircle
3 transferring /Users/kylepenfound/github.com/encircle: 6.84kB 0.0s done
3 DONE 0.1s

2 Run Go Tests
2 DONE 0.0s

4 from golang:latest
4 resolve docker.io/library/golang:latest 0.1s done
4 DONE 0.1s

5 job_two
5 CACHED

2 Run Go Tests
2 0.094 go test
2 DONE 0.1s

6 Run Go Build
0 0.067 go build
```
