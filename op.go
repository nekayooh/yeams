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
	"time"
)

type YeaRPC struct{}

type YeaModuleOne struct {
	//模块版本
	Version string
	//模块验证密钥
	Key []byte
	//模块运行地址
	Address string
	//模块绑定端口
	Port int64
	//模块启动线程数
	Thread int64
	////模块
	//Module  []*proto.SendMsg
	//模块UUID
	Uuid uuid.UUID
	//模块获取客户端
	Chan chan *proto.YeaNoticeClient
	//模块结束标志
	Close chan bool
	//模块连接客户端
	Conn []*grpc.ClientConn
	//模块客户端
	Client []*proto.YeaNoticeClient
	//存活通知时间
	Ping time.Time
}

type YeaModule struct {
	Count   int64
	Total   int64
	Modules []*YeaModuleOne
}

var ModuleMutex sync.Mutex
var YeaModules = make(map[string]*YeaModule)

//本地模块
var YeaLocalModule []string

//注册单个模块
func RegisterOneModule(index int) {
	//加载模块
	go func(module string) {
		cmd := exec.Command("./module")
		cmd.Dir = "module/" + module
		_ = cmd.Run()
		fmt.Printf("%v已关闭\n", module)
	}(YeaLocalModule[index])
}

//注册模块
func RegisterAllModules() {
	//加载模块
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

//取消注册模块
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
			//删除模块
			delete(YeaModules, a)

			lastModule := ""
			for k, _ := range YeaModules {
				lastModule += k + ","
			}
		}
	}
}

//取消注册模块
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

			lastModule := ""
			for k, _ := range YeaModules {
				lastModule += k + ","
			}
		}
		//删除模块
		delete(YeaModules, a)
	}
}

func main() {
	//file, err := os.Open("./module.json")
	//if err != nil {
	//	fmt.Println("模块配置文件加载失败")
	//	os.Exit(0)
	//}
	//defer file.Close()
	//
	//decoder := json.NewDecoder(file)
	//err = decoder.Decode(&YeaLocalModule)
	//if err != nil {
	//	fmt.Println("模块配置文件解析失败")
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
