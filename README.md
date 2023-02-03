# encircle

Run a circleci yaml locally with Dagger

## Using encircle

- `go build` to build the `encircle` binary
- `./encircle` will execute the yaml in `./circleci/config.yml`

Notes:

- Right now it's hardcoded to a specific workflow in the yaml
- Doesn't understand non-`run` steps like `checkout` or orbs
- Still a lot to go on the config
