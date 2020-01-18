package main

import (
	"context"
	"fmt"
	"github.com/nekayooh/yeams/proto"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

//接受注册
func (m YeaRPC) Register(ctx context.Context, msg *proto.RegMsg) (*proto.RegRtnMsg, error) {
	ModuleMutex.Lock()
	defer ModuleMutex.Unlock()

	//重新注册前关闭前置插件线程
	if _, ok := YeaModules[msg.Name]; ok {
		for index, _ := range YeaModules[msg.Name].Modules {
			if YeaModules[msg.Name].Modules[index].Address == msg.Address && YeaModules[msg.Name].Modules[index].Port == msg.Port {
				//关闭线程
				for i := 0; i < int(YeaModules[msg.Name].Modules[index].Thread); i++ {
					YeaModules[msg.Name].Modules[index].Close <- true
				}
				//关闭连接
				for _, v := range YeaModules[msg.Name].Modules[index].Conn {
					_ = v.Close()
				}
				//删除插件
				YeaModules[msg.Name].Modules = append(YeaModules[msg.Name].Modules[:index], YeaModules[msg.Name].Modules[index+1:]...)
				YeaModules[msg.Name].Total = int64(len(YeaModules[msg.Name].Modules))
				//delete(YeaModules, msg.Name)
			}
		}
	} else {
		YeaModules[msg.Name] = &YeaModule{
			Count:   0,
			Total:   0,
			Modules: []*YeaModuleOne{},
		}
	}

	moduleIndex := len(YeaModules[msg.Name].Modules)

	//加载插件
	YeaModules[msg.Name].Modules = append(YeaModules[msg.Name].Modules,
		&YeaModuleOne{
			Version: msg.Version,
			Key:     msg.Key,
			Address: msg.Address,
			Port:    msg.Port,
			Thread:  msg.Thread,
			Module:  msg.Module,
			Uuid:    uuid.Must(uuid.NewV5(uuid.NewV1(), msg.Name), nil),
			Chan:    make(chan *proto.YeaNoticeClient, 1),
			Close:   make(chan bool, 1),
			Conn:    []*grpc.ClientConn{},
			Client:  []*proto.YeaNoticeClient{},
		})
	//失败次数
	FailCount := 1
	//生成连接循环
	for i := 0; i < int(msg.Thread); i++ {
		YeaConn, err := grpc.Dial(msg.Address+":"+strconv.FormatInt(msg.Port, 10), grpc.WithInsecure())
		if err != nil {
			i -= 1
			FailCount += 1
			if FailCount > 4 {
				//取消注册
				delete(YeaModules, msg.Name)
				break
			} else {
				continue
			}
		}
		YeaModules[msg.Name].Modules[moduleIndex].Conn = append(YeaModules[msg.Name].Modules[moduleIndex].Conn, YeaConn)
		YeaClient := proto.NewYeaNoticeClient(YeaConn)
		YeaModules[msg.Name].Modules[moduleIndex].Client = append(YeaModules[msg.Name].Modules[moduleIndex].Client, &YeaClient)

		//将通知客户端推送到通道
		go func() {
			for {
				if <-YeaModules[msg.Name].Modules[moduleIndex].Close {
					break
				}
				YeaModules[msg.Name].Modules[moduleIndex].Chan <- &YeaClient
			}
		}()
	}
	if FailCount > 4 {
		fmt.Printf("加载插件%s[次数%v]失败\n", msg.Name, FailCount)

		//如果插件数为0，则取消
		if YeaModules[msg.Name].Total == 0 {
			delete(YeaModules, msg.Name)
		}

		//注册失败
		//ModuleLoad <- 0
		return &proto.RegRtnMsg{
			Status: proto.ReturnCode_Failure,
		}, nil
	} else {
		fmt.Printf("加载插件%s[线程%v]成功\n", msg.Name, msg.Thread)
		YeaModules[msg.Name].Total++
		//注册成功
		for _, v := range msg.Module {
			//通知关联插件
			YeaModules[v.Name].Modules[moduleIndex].Close <- false
			NoticeClient := <-YeaModules[v.Name].Modules[moduleIndex].Chan
			res, err := (*NoticeClient).Msg(ctx, &proto.SendMsg{
				Name: msg.Name,
				Json: v.Json,
			})
			if err == nil && res.Status == proto.ReturnCode_Success {
			}
		}
		//ModuleLoad <- 0
		return &proto.RegRtnMsg{
			Status: proto.ReturnCode_Success,
			Uuid:   YeaModules[msg.Name].Modules[moduleIndex].Uuid.String(),
		}, nil
	}
}

//取消注册
func (m YeaRPC) UnRegister(ctx context.Context, msg *proto.UnRegMsg) (*proto.UnRegRtnMsg, error) {
	if _, ok := YeaModules[msg.Name]; ok {
		for index, _ := range YeaModules[msg.Name].Modules {
			if YeaModules[msg.Name].Modules[index].Uuid.String() == msg.Uuid {
				//通知关联插件取消注册
				for _, v := range msg.Module {
					//通知关联插件
					if _, ok := YeaModules[v.Name]; ok {
						YeaModules[v.Name].Modules[index].Close <- false
						NoticeClient := <-YeaModules[v.Name].Modules[index].Chan
						res, err := (*NoticeClient).Msg(ctx, &proto.SendMsg{
							Name: msg.Name,
							Json: v.Json,
						})
						if err == nil && res.Status == proto.ReturnCode_Success {
						}
					}
				}
				//关闭线程
				for i := 0; i < int(YeaModules[msg.Name].Modules[index].Thread); i++ {
					YeaModules[msg.Name].Modules[index].Close <- true
				}
				//关闭连接
				for _, v := range YeaModules[msg.Name].Modules[index].Conn {
					_ = v.Close()
				}
				//删除插件
				delete(YeaModules, msg.Name)

				lastModule := ""
				for k, _ := range YeaModules {
					lastModule += k + ","
				}
				fmt.Printf("插件%s注销成功,当前剩余插件%s\n", msg.Name, lastModule)

				//删除插件
				YeaModules[msg.Name].Modules = append(YeaModules[msg.Name].Modules[:index], YeaModules[msg.Name].Modules[index+1:]...)
				YeaModules[msg.Name].Total = int64(len(YeaModules[msg.Name].Modules))

				//如果插件数为0，则取消
				if YeaModules[msg.Name].Total == 0 {
					delete(YeaModules, msg.Name)
				}

				return &proto.UnRegRtnMsg{
					Status: proto.ReturnCode_Success,
				}, nil
			} else {
				return &proto.UnRegRtnMsg{
					Status: proto.ReturnCode_Failure,
				}, nil
			}
		}
		return &proto.UnRegRtnMsg{
			Status: proto.ReturnCode_Failure,
		}, nil
	} else {
		return &proto.UnRegRtnMsg{
			Status: proto.ReturnCode_Failure,
		}, nil
	}
}

//通信
func (m YeaRPC) Msg(ctx context.Context, msg *proto.SendMsg) (*proto.SendRtnMsg, error) {

	moduleIndex := int(YeaModules[msg.Name].Count)
	if moduleIndex >= int(YeaModules[msg.Name].Total)-1 {
		YeaModules[msg.Name].Count = int64(0)
	} else {
		YeaModules[msg.Name].Count++
	}
	//if _, ok := YeaModules[msg.Name]; ok {
	//	if YeaModules[msg.Name].Modules[moduleIndex%int(YeaModules[msg.Name].Total)].Uuid.String() == msg.Uuid {
	if _, ok := YeaModules[msg.Name]; ok {
		YeaModules[msg.Name].Modules[moduleIndex%int(YeaModules[msg.Name].Total)].Close <- false
		NoticeClient := <-YeaModules[msg.Name].Modules[moduleIndex%int(YeaModules[msg.Name].Total)].Chan
		return (*NoticeClient).Msg(ctx, msg)
	} else {
		return &proto.SendRtnMsg{
			Status: proto.ReturnCode_Failure,
		}, nil
	}
	//} else {
	//	return &proto.SendRtnMsg{
	//		Status: proto.ReturnCode_Failure,
	//	}, nil
	//}
	//} else {
	//	return &proto.SendRtnMsg{
	//		Status: proto.ReturnCode_Failure,
	//	}, nil
	//}
}

func GrpcOp() {
	lis, err := net.Listen("tcp", ":20001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterYeaModuleServer(s, &YeaRPC{})

	go func() {
		//go httpserver.HttpServerStart()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
