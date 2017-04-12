package config

import (
    "fmt"
    "os"
    "path"
    "runtime"

    "github.com/spf13/viper"
)

var C *viper.Viper

func init() {
    env := os.Getenv("GOENV")
    if env == "" {
        env = "local"
    }

    _, filename, _, _ := runtime.Caller(0)
    dir := path.Dir(filename)
    C = viper.New()

    C.SetConfigName(env)
    C.AddConfigPath(dir)

    err := C.ReadInConfig()
    if err != nil {
        fmt.Printf("%#v", "test")
    }
}
