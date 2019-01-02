# vipertools

A few viper tools to shortcut things I was repeating in my projects.

# How to get

```
go get github.com/dwburke/vipertools
```

# Examples

```
import (
    "github.com/dwburke/vipertools"
)

var configPattern string

init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&configPattern, "config-pattern", "configs/*.yml", "load list of config files")
}

func initConfig() {
    if err := vipertools.ReadPattern(configPattern); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

## Other Examples:

### Merge in an array of specific config files:
```
config_files := []string{
    "/etc/app/config.yml",
    os.Getenv("HOME") + "/.app-config.yml",
}
err := vipertools.MergeConfigs(config_list)
```

### Merge in all files in a directory:
```
err := vipertools.ReadDir("/home/dwburke/configs")
```

### Merge all config files for a pattern:
```
err := vipertools.ReadPattern("/home/dwburke/configs/myapp/*.yml")
```

### Read one file in your current directory:
```
err := vipertools.ReadFile("config.yml")
```


# Pull Requests Welcome

You know the dril.

