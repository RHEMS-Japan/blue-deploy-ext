package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)
func main() {
	str := `
┏┓╋╋╋╋╋╋╋┏┓
┃┗┳┓┏┳┳━┳┛┣━┳━┳┓┏━┳┳┓
┃╋┃┗┫┃┃┻┫╋┃┻┫╋┃┗┫╋┃┃┃
┗━┻━┻━┻━┻━┻━┫┏┻━┻━╋┓┃
╋╋╋╋╋╋╋╋╋╋╋╋┗┛╋╋╋╋┗━┛
`
	fmt.Println(str)
	out, err := exec.Command("/usr/local/bin/receiver", "-v").Output()
	if err != nil {
		log.Fatalf("receiver version check error : %s", err)
	}
	fmt.Printf("bluedeploy version: \n%s", string(out))

	extv, err := exec.Command("/usr/local/bin/blue-deploy-ext", "-v").Output()
	if err != nil {
	  fmt.Println("Command Exec Error.")
	}
	fmt.Printf("bluedeploy extension version: \n%s", string(extv))
	go Run()

	t := time.NewTimer(60 * time.Second)

	<-t.C

	os.Exit(0)
}

func Run()  {
	cmd := exec.Command("/usr/local/bin/receiver")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("cmd.StdoutPipe() error : %s", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd start error : %s", err)
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Println(in.Text())
	}

	if err := in.Err(); err != nil {
		log.Fatalf("in.Err() : %s", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("cmd.Wait() : %s", err)
	}
}