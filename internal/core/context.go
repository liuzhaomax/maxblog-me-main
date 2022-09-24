package core

import (
	"crypto/rsa"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var ctx *Context
var once sync.Once

func init() {
	once.Do(func() {
		ctx = &Context{}
	})
}

func GetInstanceOfContext() *Context {
	return ctx
}

type Context struct {
	Upstream     Upstream
	Downstream   Downstream
	JWTSecret    string
	PrivateKey   *rsa.PrivateKey
	PublicKey    *rsa.PublicKey
	PublicKeyStr string
}

type Upstream struct {
	MaxblogFEMain Address
}

type Downstream struct {
	MaxblogBEDemo Address
}

type Address struct {
	Host string
	Port int
}

func SetUpstreamAddr(host string, port int) {
	ctx.Upstream.MaxblogFEMain.Host = host
	ctx.Upstream.MaxblogFEMain.Port = port
}

func GetUpstreamAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Upstream.MaxblogFEMain.Host, ctx.Upstream.MaxblogFEMain.Port)
}

func SetDownstreamAddr(host string, port int) {
	ctx.Downstream.MaxblogBEDemo.Host = host
	ctx.Downstream.MaxblogBEDemo.Port = port
}

func GetDownstreamMaxblogBETemplateAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Downstream.MaxblogBEDemo.Host, ctx.Downstream.MaxblogBEDemo.Port)
}

func GetProjectPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	indexWithoutFileName := strings.LastIndex(path, string(os.PathSeparator))
	indexWithoutLastPath := strings.LastIndex(path[:indexWithoutFileName], string(os.PathSeparator))
	return strings.Replace(path[:indexWithoutLastPath], "\\", "/", -1)
}

func GetPublicKey() *rsa.PublicKey {
	return GetInstanceOfContext().PublicKey
}

func GetPublicKeyStr() string {
	return GetInstanceOfContext().PublicKeyStr
}

func GetPrivateKey() *rsa.PrivateKey {
	return GetInstanceOfContext().PrivateKey
}
