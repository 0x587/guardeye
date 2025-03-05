package ros

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/samber/lo"
)

var rosEnvInitOnce sync.Once
var rosEnv map[string]string

func (i *impl) rosExec(cmdStr string) (string, error) {
	rosEnvInitOnce.Do(func() {
		rosEnv = make(map[string]string)
		env := []string{
			"/opt/ros/humble/setup.sh",
			"/home/shawn/ros2_ws/install/setup.sh",
		}
		env = lo.Map(env, func(s string, _ int) string { return "source " + s })
		cmd := exec.Command("bash", "-c", strings.Join(append(env, "env"), " && "))
		output := lo.Must(cmd.CombinedOutput())
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				rosEnv[parts[0]] = parts[1]
			}
		}
	})

	cmd := exec.Command("bash", "-c", cmdStr)

	env := os.Environ()
	for k, v := range rosEnv {
		env = append(env, k+"="+v)
	}
	cmd.Env = env
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Join(errors.New("cmd: "+cmd.String()), err, errors.New(string(output)))
	}
	return string(output), nil
}
