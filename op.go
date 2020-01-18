package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/nekayooh/yeams/proto"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"os"
	"os/exec"
	"sync"
)

type YeaRPC struct{}

type YeaModuleOne struct {
	Version string
	Key     []byte
	Address string
	Port    int64
	Thread  int64
	Module  []*proto.SendMsg
	Uuid    uuid.UUID
	Chan    chan *proto.YeaNoticeClient
	Close   chan bool
	Conn    []*grpc.ClientConn
	Client  []*proto.YeaNoticeClient
}

type YeaModule struct {
	Count   int64
	Total   int64
	Modules []*YeaModuleOne
}

var ModuleMutex sync.Mutex
var YeaModules = make(map[string]*YeaModule)
var YeaLocalModule []string
var PublicPem = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCAUVJRyimtN/h7mKKjaF/iXcc0
+if908qb2zGQGvC5tn4q8xAvLkMlQth+iZg1YS85n2Z9OZt8fNLsWHFu8yf6fx7i
4O6NQkKG9OD6R1KpN/mWV95cZpRR7HTu4Wa/qj3CokkKXza5+2VdWjiEDodoZ8je
q6Ptzg4XcrNytPtZiwIDAQAB
-----END PUBLIC KEY-----`
var PrivatePem = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCAUVJRyimtN/h7mKKjaF/iXcc0+if908qb2zGQGvC5tn4q8xAv
LkMlQth+iZg1YS85n2Z9OZt8fNLsWHFu8yf6fx7i4O6NQkKG9OD6R1KpN/mWV95c
ZpRR7HTu4Wa/qj3CokkKXza5+2VdWjiEDodoZ8jeq6Ptzg4XcrNytPtZiwIDAQAB
AoGAAi20m2rhQO2drGDaXLKISOGVYOg2XEWHlRR9nO1i1ORXnrG4aUn8cy+ABCCg
kuzEI9L5pyVJPIMqBD/kJ8wSYP80I1zLbS9i+NfU11CbKJXpmcu6ma1J5PSw5qGg
hy0UQI1z0raaVNXPPTYJO9ywzUQ1yNcHkDW+epog3PSueaECQQCB+tbYZ07O7vrK
tbUH7yVfUigwEByP14JJUaIZoUWgv0KJFudP+Zg286bh2+1K8dhxA6LZ5pqYYdIn
NHUpd5IDAkEA/LntcsvP03Z3bR4fVdPTXDc4KisvJTdPRcJc6u/HlkesV2JDBOAA
wdHNzFB+sywyUcr5XpuaFGIbIKvWE4uH2QJAFYjjk5L6IZrCfldAmQHsJTDNa7kf
ok1ITrFxs+FeUdWeRmw/AqcNqv0PRxhS5jnPbFn33zYvotOCJ/CvAKHI1QJARuGq
3FCXiHqogj05kqvnkuyV3xXfkjOSE0GxJ996fga6KoQPwfVFoRbD/rLw5jXWIySn
jkZcD614aFBpqW+v+QJAFPojoOPSiFgmjFPM/pTim3vtNw0zu3k4hhD3pI86gluy
U+2Hckc91DGYAvPp9D++xOVFkKOpk/aVr4LJ7harcg==
-----END RSA PRIVATE KEY-----`

//注册单个插件
func RegisterOneModule(index int) {
	//加载插件
	go func(module string) {
		cmd := exec.Command("./module")
		cmd.Dir = "module/" + module
		_ = cmd.Run()
		fmt.Printf("%v已关闭\n", module)
	}(YeaLocalModule[index])
}

//注册插件
func RegisterAllModules() {
	//加载插件
	for _, v := range YeaLocalModule {
		go func(module string) {
			cmd := exec.Command("./module")
			cmd.Dir = "module/" + module
			//if runtime.GOOS == "windows" {
			//	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			//}
			_ = cmd.Run()
			fmt.Printf("%v已关闭\n", module)
		}(v)
		//<-ModuleLoad
	}
}

//取消注册插件
func UnRegisterOneModule() {
	for a, _ := range YeaModules {
		for index, _ := range YeaModules[a].Modules {
			_, _ = (*YeaModules[a].Modules[index].Client[0]).UnRegister(context.Background(), &proto.UnRegMsg{
				Uuid: YeaModules[a].Modules[index].Uuid.String(),
			})
			//关闭线程
			for i := 0; i < int(YeaModules[a].Modules[index].Thread); i++ {
				YeaModules[a].Modules[index].Close <- true
			}
			//关闭连接
			for _, v := range YeaModules[a].Modules[index].Conn {
				_ = v.Close()
			}
			//删除插件
			delete(YeaModules, a)

			lastModule := ""
			for k, _ := range YeaModules {
				lastModule += k + ","
			}
		}
	}
}

//取消注册插件
func UnRegisterAllModules() {
	for a, _ := range YeaModules {
		for index, _ := range YeaModules[a].Modules {
			_, _ = (*YeaModules[a].Modules[index].Client[0]).UnRegister(context.Background(), &proto.UnRegMsg{
				Uuid: YeaModules[a].Modules[index].Uuid.String(),
			})
			//关闭线程
			for i := 0; i < int(YeaModules[a].Modules[index].Thread); i++ {
				YeaModules[a].Modules[index].Close <- true
			}
			//关闭连接
			for _, v := range YeaModules[a].Modules[index].Conn {
				_ = v.Close()
			}
			//删除插件
			delete(YeaModules, a)

			lastModule := ""
			for k, _ := range YeaModules {
				lastModule += k + ","
			}
		}
	}
}

func main() {
	//file, err := os.Open("./module.json")
	//if err != nil {
	//	fmt.Println("插件配置文件加载失败")
	//	os.Exit(0)
	//}
	//defer file.Close()
	//
	//decoder := json.NewDecoder(file)
	//err = decoder.Decode(&YeaLocalModule)
	//if err != nil {
	//	fmt.Println("插件配置文件解析失败")
	//	os.Exit(0)
	//}

	GrpcOp()

	input := bufio.NewScanner(os.Stdin)
	for {
		input.Scan()
		switch input.Text() {
		case "a":
			RegisterAllModules()
			fmt.Println("全部已启动")
		case "s":
			UnRegisterAllModules()
			fmt.Println("全部已关闭")
		}
	}
}
