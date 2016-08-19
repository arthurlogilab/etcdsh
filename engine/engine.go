package engine

import (
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"

	"github.com/kamilhark/etcdsh/pathresolver"
)

type Engine struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (e *Engine) ResolvePath(subPath string) string {
	return e.PathResolver.Resolve(subPath)
}

func (e *Engine) Set(key string, value string) {
	key = e.ResolvePath(key)
	_, err := e.KeysApi.Set(context.Background(), key, value, &client.SetOptions{Dir: false})
	if err != nil {
		fmt.Println(err)
	}
}

func (e *Engine) Get(key string, recursive bool) *client.Node {
	key = e.ResolvePath(key)
	response, err := e.KeysApi.Get(context.Background(), key, &client.GetOptions{Recursive: recursive})
	if err != nil {
		fmt.Println(err)
	} else {
		return response.Node
	}
	return nil
}

func (e *Engine) MkDir(key string) {
	key = e.ResolvePath(key)
	_, err := e.KeysApi.Set(context.Background(), key, "", &client.SetOptions{Dir: true})
	if err != nil {
		fmt.Println(err)
	}
}

func (e *Engine) Rm(key string, recursive bool) {
	key = e.ResolvePath(key)
	_, err := e.KeysApi.Delete(context.Background(), key, &client.DeleteOptions{Recursive: recursive})
	if err != nil {
		fmt.Println(err)
	}
}

func (e *Engine) Cp(srcPath string, dstPath string) {
	srcPath = e.ResolvePath(srcPath)
	dstPath = e.ResolvePath(dstPath)

	node := e.Get(srcPath, true)
	if node != nil {
		e.recurseCp(node, srcPath, dstPath)
	}
}

func (e *Engine) Mv(srcPath string, dstPath string) {
	e.Cp(srcPath, dstPath)
	e.Rm(srcPath, true)
}

func (e *Engine) recurseCp(n *client.Node, srcPath string, dstPath string) {
	newKey := strings.Replace(n.Key, srcPath, dstPath, 1)
	if n.Dir {
		e.MkDir(newKey)
		for _, node := range n.Nodes {
			e.recurseCp(node, srcPath, dstPath)
		}
	} else {
		e.Set(newKey, n.Value)
	}
}
