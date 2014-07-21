package config

import (
    "os"
    "io"
    "strings"
    "github.com/kylelemons/go-gypsy/yaml"
)

var conf yaml.Map

func LoadFile(path string) {
    file, _ := os.Open(path)
    defer file.Close()
    Load(file)
}

func Load(reader io.Reader) {
    parsed, _ := yaml.Parse(reader)
    conf = parsed.(yaml.Map)
}

func Get(path string) string {
    return get(conf, strings.Split(path, ":")).String()
}

func get(step yaml.Map, segments []string) yaml.Scalar {
    val, ok := step[segments[0]]
    if !ok {
        return ""
    }
    if len(segments) == 1 {
        return val.(yaml.Scalar)
    }
    return get(val.(yaml.Map), segments[1:])
}
