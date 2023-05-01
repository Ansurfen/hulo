package server

import (
	"context"
	"errors"
	"fmt"
	pb "hulo/daemon/proto"
	"hulo/shell"
	"net"
	"path"

	"github.com/ansurfen/cushion/utils"
	"google.golang.org/grpc"
)

func init() {
	utils.InitLoggerWithDefault()
}

const (
	LACK_SERVICE_NAME = "lack service name"
	SYS_ERR           = "sys err"
)

type HuloDaemon struct {
	pb.UnimplementedHuloDaemonServer
	services map[string]*Service
	opt      DaemonOpt
}

type HuloDaemonEnv struct{}

func NewHuloDaemon() *HuloDaemon {
	opt.Parse()
	utils.NewEnv(utils.EnvOpt[HuloDaemonEnv]{
		Payload: HuloDaemonEnv{},
		Workdir: ".hulo",
	})
	return &HuloDaemon{
		opt:      opt,
		services: make(map[string]*Service),
	}
}

func (s *HuloDaemon) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

func (s *HuloDaemon) StartService(ctx context.Context, req *pb.StartServiceRequest) (*pb.StartServiceResponse, error) {
	if len(req.ServiceName) == 0 {
		return &pb.StartServiceResponse{}, errors.New(LACK_SERVICE_NAME)
	}
	var service *Service
	if serv, ok := s.services[req.ServiceName]; ok {
		// s.services[req.ServiceName].cnt++
		// ! 应该根据tid决定要不要追加了, tid是个 bitmap
		service = serv
	} else {
		port, err := utils.RandomPort()
		if err != nil {
			return &pb.StartServiceResponse{}, errors.New(SYS_ERR)
		}
		s.services[req.ServiceName] = &Service{
			port: port,
			cnt:  1,
		}
		service = s.services[req.ServiceName]
		sh := shell.NewHuloShell()
		err = sh.StartProcess(path.Join(utils.GetEnv().Workdir(), "bin", shell.ELF(req.ServiceName)), fmt.Sprintf("-p %d", service.port))
		if err != nil {
			return &pb.StartServiceResponse{}, err
		}
	}
	return &pb.StartServiceResponse{
		Port: int32(service.port),
	}, nil
}

func (s *HuloDaemon) CloseService(ctx context.Context, req *pb.CloseServiceRequest) (*pb.CloseServiceResponse, error) {
	if len(req.ServiceName) == 0 {
		return &pb.CloseServiceResponse{}, nil
	}
	if serv, ok := s.services[req.ServiceName]; ok {
		serv.cnt--
		if serv.cnt == 0 {
			ps := shell.PowerShell{}
			process := ps.QueryProcessByName(req.ServiceName)
			for _, p := range process {
				ps.StopProcess(p.Pid())
			}
		}
	}
	return &pb.CloseServiceResponse{}, nil
}

func (s *HuloDaemon) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *s.opt.Port))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterHuloDaemonServer(gsrv, s)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}
