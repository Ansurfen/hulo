package shell

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ansurfen/cushion/utils"
)

type Bash struct{}

func (Bash) StartProcess(file, args string) error {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf(`%s %s &`, file, args))
	return cmd.Start()
}

func (Bash) QueryProcessByName(name string) []Process {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("ps -ef|grep %s", name))
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	processes := []Process{}
	utils.ReadLineFromString(string(out), func(s string) string {
		res := strings.Split(utils.TrimMultiSpace(strings.TrimSpace(s)), " ")
		processes = append(processes, BashProcess{
			uid:   res[0],
			pid:   Str2Int(res[1]),
			ppid:  Str2Int(res[2]),
			c:     Str2Int(res[3]),
			stime: res[4],
			tty:   res[5],
			time:  res[6],
			cmd:   strings.Join(res[7:], " "),
		})
		return ""
	})
	return processes
}

func (Bash) StopProcess(id int) error {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("kill %d", id))
	_, err := cmd.CombinedOutput()
	return err
}

type BashProcess struct {
	uid   string
	pid   int
	ppid  int
	c     int
	stime string
	tty   string
	time  string
	cmd   string
}

func (p BashProcess) Pid() int {
	return p.pid
}
