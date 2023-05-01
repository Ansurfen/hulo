package terminal

import (
	"context"
	"errors"
	"fmt"
	pb "hulo/daemon/proto"
	hi "hulo/sdk/go"
	"hulo/shell"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SUSPEND = iota
	READY
	RUNNING
	STOP
)

type huloService struct {
	port  int32
	state uint8
	conn  *grpc.ClientConn
	cli   hi.HuloInterfaceClient
}

type HuloDaemon struct {
	conn     *grpc.ClientConn
	cli      pb.HuloDaemonClient
	services map[string]*huloService
	term     *HuloTerminal
}

func NewDaemon(port int) *HuloDaemon {
	daemon := &HuloDaemon{
		services: make(map[string]*huloService),
	}
	var err error
	daemon.conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	daemon.cli = pb.NewHuloDaemonClient(daemon.conn)
	return daemon
}

func (d *HuloDaemon) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := d.cli.Ping(ctx, &pb.PingRequest{})
	return err
}

func (d *HuloDaemon) StartService(name string) uint8 {
	if service, ok := d.services[name]; ok {
		return service.state
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := d.cli.StartService(ctx, &pb.StartServiceRequest{
		ServiceName: name,
	})
	if err != nil {
		d.term.IO.WriteStderr(err.Error())
		return STOP
	}
	var service *huloService
	if s, ok := d.services[name]; ok {
		service = s
	} else {
		service = &huloService{
			port:  res.Port,
			state: SUSPEND,
		}
		d.services[name] = service
	}
	if err := d.connectService(name); err != nil {
		service.state = STOP
		return STOP
	}
	service.state = READY
	return service.state
}

func (d *HuloDaemon) SwitchServiceState(service string, state int) {
	if service, ok := d.services[service]; ok {
		service.state = uint8(state)
	}
}

func (d *HuloDaemon) CloseService(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := d.cli.CloseService(ctx, &pb.CloseServiceRequest{
		ServiceName: name,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (d *HuloDaemon) Call(service, fun, arg string) error {
	c := d.GetService("mysql")
	_, err := c.Call(context.Background(), &hi.CallRequest{
		Func: fun,
		Arg:  arg,
	})
	if err != nil {
		d.term.IO.WriteStderr(err.Error())
	}
	return err
}

func (d *HuloDaemon) Close() {
	for service, _ := range d.services {
		d.CloseService(service)
	}
	d.conn.Close()
	d.kill()
}

func (d *HuloDaemon) kill() {
	sh := shell.NewHuloShell()
	process := sh.QueryProcessByName("hulo_daemon")
	for _, p := range process {
		sh.StopProcess(p.Pid())
	}
}

func (d *HuloDaemon) connectService(name string) error {
	if serive, ok := d.services[name]; ok {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", serive.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			zap.S().Error("did not connect: %v", err)
		}
		d.services[name].conn = conn
		d.services[name].cli = hi.NewHuloInterfaceClient(conn)
		return nil
	}
	return errors.New("err")
}

func (d *HuloDaemon) GetService(name string) hi.HuloInterfaceClient {
	if service, ok := d.services[name]; ok {
		return service.cli
	}
	return nil
}
