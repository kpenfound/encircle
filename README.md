# encircle

Run a circleci yaml locally with Dagger

## Using encircle

- `go build` to build the `encircle` binary
- `./encircle` will execute the yaml in `./circleci/config.yml`

Notes:

- Right now it's hardcoded to a specific workflow in the yaml
- Still a lot to go on the config
- Cant handle conditionals

Sample output:

```shell
go build && ./encircle
Loading config at ./.circleci/config.yml
Dont know how to handle command when
Dont know how to handle command when
Dont know how to handle command when
Dont know how to handle command when
Dont know how to handle command when
Dont know how to handle command when
warning: skipping checkout for local dev
Running workflow test
Running job job_one
1 resolve image config for docker.io/library/node:16
1 DONE 1.0s

2 Run Unit Tests
2 DONE 0.0s

3 host.directory /Users/kylepenfound/github.com/encircle
3 transferring /Users/kylepenfound/github.com/encircle: 261.13kB 0.1s
3 ...

4 from node:16
4 resolve docker.io/library/node:16 0.1s done
4 DONE 0.1s

3 host.directory /Users/kylepenfound/github.com/encircle
3 transferring /Users/kylepenfound/github.com/encircle: 9.20MB 0.5s done
3 DONE 0.5s

2 Run Unit Tests
2 DONE 0.0s

5 job_one
5 CACHED

3 host.directory /Users/kylepenfound/github.com/encircle
3 DONE 0.5s

6 Install npm dependencies
6 0.119 npm install
6 DONE 0.2s

2 Run Unit Tests
2 DONE 0.0s
Running job job_two

2 Run Unit Tests
2 0.069 npm test
2 DONE 0.1s

7 resolve image config for docker.io/library/golang:latest
7 DONE 0.6s

8 from golang:latest
8 resolve docker.io/library/golang:latest
8 DONE 0.2s

8 from golang:latest
8 resolve docker.io/library/golang:latest 0.2s done
8 DONE 0.2s

9 job_two
9 CACHED

10 Run Go Tests
0 0.092 go test
10 DONE 0.1s

11 Run Go Build
0 0.076 go build
11 DONE 0.1s
Running job orb_test

12 resolve image config for docker.io/cimg/base:stable
12 DONE 0.6s

13 from cimg/base:stable
13 resolve docker.io/cimg/base:stable
13 resolve docker.io/cimg/base:stable 0.1s done
13 DONE 0.1s

14 orb_test
14 CACHED

15 Install Node.js 16.13
15 0.207   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
15 0.207                                  Dload  Upload   Total   Spent    Left  Speed
100 15916  100 15916    0     0  84211      0 --:--:-- --:--:-- --:--:-- 86972
=> => Downloading nvm from git to '/home/circleci/.nvm'
15 0.512 Cloning into '/home/circleci/.nvm'...
15 2.929 * (HEAD detached at FETCH_HEAD)
15 2.929   master
15 2.973 => Compressing and cleaning up git repository
15 3.033
15 3.112 => Appending nvm source string to /home/circleci/.profile
15 3.116 => bash_completion source string already in /home/circleci/.profile
15 3.336 => Close and reopen your terminal to start using nvm or run the following to use it now:
15 3.336
15 3.336 export NVM_DIR="$HOME/.nvm"
15 3.336 [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"   This loads nvm
15 5.240 Downloading and installing node v16.13.2...
15 6.021 Downloading https://nodejs.org/dist/v16.13.2/node-v16.13.2-linux-x64.tar.xz...
 100.0% 21.9%
15 7.752 Computing checksum with sha256sum
15 7.989 Checksums matched!
15 15.60 Now using node v16.13.2 (npm v8.1.2)
15 16.99 Creating default alias: default -> 16.13 (-> v16.13.2 *)
15 18.06 default -> 16.13 (-> v16.13.2 *)
15 DONE 18.2s

16
16 DONE 0.0s

16
16 3.452 v16.13.2
```
