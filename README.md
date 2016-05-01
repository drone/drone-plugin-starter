# drone-plugin-starter

Starter project for creating Drone plugins.

### Metadata

Build and Repository metadata are prefixed with `DRONE_` and sent to the plugin at runtime. The full list of available parameters are already included in the `main.go` file as command line flags. You should remove un-used parameters from the list so that one can easily see which parameters are used by which plugins.

Example parameters:

```
cli.IntFlag{
    Name:   "build.number",
    Usage:  "build number",
    EnvVar: "DRONE_BUILD_NUMBER",
},
cli.StringFlag{
    Name:   "build.status",
    Usage:  "build status",
    Value:  "success",
    EnvVar: "DRONE_BUILD_STATUS",
},
```

### Parameters

Plugin parameters are defined in the yaml file:

```
slack:
  channel: dev
  username: drone
```

They are prefixed with `PLUGIN_` and sent to the plugin at runtime:

```
PLUGIN_CHANNEL=dev
PLUGIN_USERNAME=drone
```

These parameters can be retrieved using `cli.Flag` as seen below:

```
cli.StringFlag{
    Usage:  "slack channel",
    EnvVar: "PLUGIN_CHANNEL",
},
cli.StringFlag{
    Usage:  "slack username",
    EnvVar: "PLUGIN_USERNAME",
},
```

### Secrets

Sensitive fields should not be specified in the yaml file. Instead they are passed to your plugin as environment variable. Secrets should use a prefix that corresponds to the plugin name. For example, the Slack plugin prefixes secrets with `SLACK_`:

```
cli.StringFlag{
    Usage:  "slack api token",
    EnvVar: "SLACK_TOKEN",
},
```

### Logos

Please replace the logo.svg file with a meaningful svg logo for your plugin. If you are you building a Slack plugin, for example, please provide the Slack logo. This icon is displayed when your plugin is listed in the official plugin index.

### Docs

Please provide a DOCS.md file in the root of your repository that documents plugin usage. This documentation is displayed when your plugin is listed in the official plugin index. Use the README.md file to describe building and testing the plugin locally.

### Images

Images are distributed as Docker images in the public Docker registry. Please use a minimalist alpine image when possible to limit the image download size. We are also working on supporting multiple plugin architectures, with compatibility guidelines coming soon

### Testing

Please create plugins that are easily runnable from the command line. This makes it much easier to debug and test plugins locally without having to launch actual builds.

### Vendoring

Please vendor dependencies in a manner compatible with `GOVENDOREXPERIMENT`. All official drone plugins should use [govend](https://github.com/govend/govend) with the `--prune` flag.
