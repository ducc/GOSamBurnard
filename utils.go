package main

import (
    "io/ioutil"
    "encoding/json"
)

func loadConfig(path string) (*config, error) {
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var conf config
    err = json.Unmarshal(bytes, &conf)
    if err != nil {
        return nil, err
    }
    return &conf, nil
}

func toJSONString(input interface{}) (*string, error) {
    body, err := json.Marshal(&input)
    if err != nil {
        return nil, err
    }
    temp := string(body)
    return &temp, nil
}