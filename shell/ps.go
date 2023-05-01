package shell

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ansurfen/cushion/utils"
)

type PowerShell struct{}

func (PowerShell) StartProcess(file, args string) error {
	cmd := exec.Command("powershell", fmt.Sprintf(`Start-Process -FilePath "%s" -ArgumentList "%s" -WindowStyle Hidden`, file, args))
	_, err := cmd.CombinedOutput()
	return err
}

func (PowerShell) QueryProcessByName(name string) []Process {
	cmd := exec.Command("powershell", fmt.Sprintf(`Get-Process | Where-Object {$_.ProcessName -eq "%s"}`, name))
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	processes := []Process{}
	idx := 0
	utils.ReadLineFromString(string(out), func(s string) string {
		if idx > 2 && len(s) > 0 {
			res := strings.Split(utils.TrimMultiSpace(strings.TrimSpace(s)), " ")
			if len(res) != 8 {
				return ""
			}
			processes = append(processes, PsProcess{
				handles: Str2Float64(res[0]),
				npm:     Str2Float64(res[1]),
				pm:      Str2Float64(res[2]),
				ws:      Str2Float64(res[3]),
				cpu:     Str2Float64(res[4]),
				id:      Str2Int(res[5]),
				si:      Str2Int(res[6]),
				name:    res[7],
			})
		}
		idx++
		return ""
	})
	return processes
}

func (PowerShell) StopProcess(id int) error {
	cmd := exec.Command("powershell", fmt.Sprintf(`Stop-Process -Id %d`, id))
	_, err := cmd.CombinedOutput()
	return err
}

type PsProcess struct {
	handles float64
	npm     float64
	pm      float64
	ws      float64
	cpu     float64
	id      int
	si      int
	name    string
}

func (p PsProcess) Pid() int {
	return p.id
}
