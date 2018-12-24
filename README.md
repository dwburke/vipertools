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
    rootCmd.PersistentFlags().StringVar(&configPattern, "config-pattern", "configs/*.tml", "load list of config files")
}

func initConfig() {
    if err := vipertools.ReadPattern(configPattern); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

## Other Examples:

### Read a slice of specific config files:
```
err := MergeConfigs(config_list)
```

### Read all files in a directory:
```
err := ReadDir("/home/dwburke/configs")
```

### Read all config files for a pattern:
```
err := ReadPattern("/home/dwburke/configs/myapp/*.yml")
```

### Read all files in a directory:
```
err := ReadFile("config.yml")
```


# Pull Requests Welcome

You know the dril.

