package etcd

import (
	"time"
	"log"
	"github.com/coreos/etcd/client"
	"context"
	"fmt"
)

var (
	Endpoints []string
)

func Connect() (client.Client, error){
	cfg := client.Config{
		Endpoints:               Endpoints,
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return c, err
}

func Create(c client.Client, key string, value string) error {
	kAPI := client.NewKeysAPI(c)
	result, err := kAPI.Create(context.Background(), key, value)
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func Delete(c client.Client, key string, option *client.DeleteOptions) error {
	kAPI := client.NewKeysAPI(c)
	result, err := kAPI.Delete(context.Background(), key, option)
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func Set(c client.Client, key string, value string, option *client.SetOptions) (string, error) {
	kAPI := client.NewKeysAPI(c)
	result, err := kAPI.Set(context.Background(), key, value, option)
	if err != nil{
		return "", err
	}
	return result.Node.String(), nil
}

func Get(c client.Client, key string, option *client.GetOptions) (string, error) {
	kAPI := client.NewKeysAPI(c)
	result, err := kAPI.Get(context.Background(), key, option)
	fmt.Println("----", result)
	if err != nil{
		return "", err
	}
	return result.Node.Value, nil
}

func Watch(c client.Client, key string, option *client.WatcherOptions) (<-chan string) {
	resultChan := make(chan string)
	go func(){
		kAPI := client.NewKeysAPI(c)
		watcher := kAPI.Watcher(key, option)

		for{
			result, err :=  watcher.Next(context.Background())
			if err != nil{
				fmt.Println(err)
			}
			resultChan <- result.Action
		}
	}()
	return resultChan
}

func Exist(c client.Client, key string) (bool, error) {
	kAPI := client.NewKeysAPI(c)
	result, err := kAPI.Get(context.Background(), key, &client.GetOptions{})
	if err != nil{
		return false, err
	}
	fmt.Println(result)
	return true, nil
}