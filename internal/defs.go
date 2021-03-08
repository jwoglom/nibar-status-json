package internal

import (
	"context"
	"encoding/json"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, s *Status, f workFunc) {
	defer wg.Done()
	f(s)
}

func RunStatus() *Status {
	s := NewStatus()
	var wg sync.WaitGroup

	wg.Add(len(cmds))
	for _, cmd := range cmds {
		go work(&wg, s, cmd.workFunc)
	}

	wg.Wait()
	return s
}

func run(cmd string, timeout time.Duration) string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := exec.Command("bash", "-c", cmd)

	out, err := c.Output()

	if ctx.Err() == context.DeadlineExceeded {
		return "deadline exceeded"
	}

	if err != nil {
		panic(err)
	} else {
		return strings.TrimRight(string(out), "\n")
	}

}

func runlines(cmd string, timeout time.Duration) []string {
	return strings.Split(run(cmd, timeout), "\n")
}

func unmarshalJson(input string) map[string]interface{} {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		panic(err)
	}
	return raw
}

func unmarshalJsonArray(input string) []interface{} {
	var raw []interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		panic(err)
	}
	return raw
}
