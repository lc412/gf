// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 配置管理.
// 配置文件格式支持：json, xml, toml, yaml/yml
package gcfg

import (
    "sync"
    "strings"
    "gitee.com/johng/gf/g/os/gfile"
    "gitee.com/johng/gf/g/container/gmap"
    "gitee.com/johng/gf/g/encoding/gjson"
)

const (
    gDEFAULT_CONFIG_FILE = "config.yml" // 默认的配置管理文件名称
)

// 配置管理对象
type Config struct {
    mu    sync.RWMutex             // 并发互斥锁
    path  string                   // 配置文件存放目录，绝对路径
    jsons *gmap.StringInterfaceMap // 配置文件对象
}

// 生成一个配置管理对象
func New(path string) *Config {
    return &Config {
        path  : path,
        jsons : gmap.NewStringInterfaceMap(),
    }
}

// 判断从哪个配置文件中获取内容
func (c *Config) filePath(files []string) string {
    file := gDEFAULT_CONFIG_FILE
    if len(files) > 0 {
        file = files[0]
    }
    c.mu.RLock()
    fpath := c.path + gfile.Separator + file
    c.mu.RUnlock()
    return fpath
}

// 设置配置管理器的配置文件存放目录绝对路径
func (c *Config) SetPath(path string) {
    c.mu.Lock()
    if strings.Compare(c.path, path) != 0 {
        c.path  = path
        c.jsons = gmap.NewStringInterfaceMap()
    }
    c.mu.Unlock()
}

// 添加配置文件到配置管理器中，第二个参数为非必须，如果不输入表示添加进入默认的配置名称中
func (c *Config) getJson(files []string) *gjson.Json {
    fpath := c.filePath(files)
    if r := c.jsons.Get(fpath); r != nil {
        return r.(*gjson.Json)
    }
    if j, err := gjson.Load(fpath); err == nil {
        c.jsons.Set(fpath, j)
        return j
    }
    return nil
}

// 获取配置项，当不存在时返回nil
func (c *Config) Get(pattern string, files...string) interface{} {
    if j := c.getJson(files); j != nil {
        return j.Get(pattern)
    }
    return nil
}

// 获得一个键值对关联数组/哈希表，方便操作，不需要自己做类型转换
// 注意，如果获取的值不存在，或者类型与json类型不匹配，那么将会返回nil
func (c *Config) GetMap(pattern string, files...string)  map[string]interface{} {
    if j := c.getJson(files); j != nil {
        return j.GetMap(pattern)
    }
    return nil
}

// 获得一个数组[]interface{}，方便操作，不需要自己做类型转换
// 注意，如果获取的值不存在，或者类型与json类型不匹配，那么将会返回nil
func (c *Config) GetArray(pattern string, files...string)  []interface{} {
    if j := c.getJson(files); j != nil {
        return j.GetArray(pattern)
    }
    return nil
}

// 返回指定json中的string
func (c *Config) GetString(pattern string, files...string) string {
    if j := c.getJson(files); j != nil {
        return j.GetString(pattern)
    }
    return ""
}

// 返回指定json中的bool
func (c *Config) GetBool(pattern string, files...string) bool {
    if j := c.getJson(files); j != nil {
        return j.GetBool(pattern)
    }
    return false
}

// 返回指定json中的float32
func (c *Config) GetFloat32(pattern string, files...string) float32 {
    if j := c.getJson(files); j != nil {
        return j.GetFloat32(pattern)
    }
    return 0
}

// 返回指定json中的float64
func (c *Config) GetFloat64(pattern string, files...string) float64 {
    if j := c.getJson(files); j != nil {
        return j.GetFloat64(pattern)
    }
    return 0
}

// 返回指定json中的float64->int
func (c *Config) GetInt(pattern string, files...string)  int {
    if j := c.getJson(files); j != nil {
        return j.GetInt(pattern)
    }
    return 0
}

// 返回指定json中的float64->uint
func (c *Config) GetUint(pattern string, files...string)  uint {
    if j := c.getJson(files); j != nil {
        return j.GetUint(pattern)
    }
    return 0
}