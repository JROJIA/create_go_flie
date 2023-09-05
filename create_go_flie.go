package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

func main() {

	input := bufio.NewScanner(os.Stdin)

	read := "命名规则:\n1.小写英文字母\n2.最少三个字符\n3.不能以数字开头\n4.不包含空格除_+外符号\n\n"

	fmt.Print(read)

	fmt.Print("请输入:")
	// go flash()
	for input.Scan() {
		ins := input.Text()

		fmt.Printf("\n您输入的文件名是:%s\n", ins)

		if inerr := islowlittle(ins); inerr != nil {
			fmt.Printf("%s\n\n重新输入文件名:", inerr)
			continue
		}

		if fileIsExt(ins) {
			return
		}

	}

}

func inall() {
	// 获取终端的文件描述符
	fd := int(os.Stdin.Fd())

	// 设置终端为非规范模式
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("无法设置终端为非规范模式:", err)
		return
	}
	defer term.Restore(fd, oldState)

	// 创建一个通道用于接收信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("按下Ctrl+C (SIGINT)或Ctrl+D (SIGTERM)来退出")
	fmt.Print("开始输入字符: ")

	// 创建一个字节切片来保存输入的字符
	input := make([]byte, 1)

	for {
		select {
		case sig := <-sigCh:
			fmt.Printf("\n接收到信号: %v,退出程序\n", sig)
			return
		default:
			_, err := os.Stdin.Read(input)
			if err != nil {
				fmt.Println("\n读取输入时出错:", err)
				return
			}
			fmt.Printf("捕获到字符: %c\n", input[0])
		}
	}
}

func flash() {

	fmt.Println("flash")
	// for {
	// 	fmt.Print(" ")
	// 	time.Sleep(1e3)
	// 	fmt.Print("\b")
	// 	time.Sleep(5e8)

	// }
	for {
		for _, l := range "|\\-/" {

			fmt.Printf("%s", string(l))
			fmt.Print("\b")
			time.Sleep(5e8)
		}
	}
}

func islowlittle(s string) (err error) {

	// var err error
	strbf := strings.Builder{}
	strUplit := strings.Builder{}
	if slen := len(s); slen < 3 {
		// err = errors.Join(err, fmt.Errorf("文件名%d个字符,最少需要3个字符", slen))
		err = fmt.Errorf("文件名%d个字符,最少需要3个字符", slen)
	}

	if len(s) != 0 {
		switch s[0] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			err = errors.Join(err, fmt.Errorf("首字符为%c,数字不允许", s[0]))
		}
	}

	for _, l := range s {

		switch x := string(l); x {
		case "/", "\\", ":", "*", "\"", "<", ">", "|", "?":
			strbf.WriteString(x)
			continue
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "_", "+":
			continue
		}

		switch {
		case l >= 'a' && l <= 'z':

		default:

			strUplit.WriteRune(l)
		}
	}

	if strbf.Len() > 0 {
		err = errors.Join(err, fmt.Errorf("Windows文件名不允许使用字符:%s", strbf.String()))
	}

	if strUplit.Len() > 0 {
		err = errors.Join(err, fmt.Errorf("不符合字符:%s", strUplit.String()))
	}

	return

}

func CrtateGofile(s string) bool {
	err := os.Mkdir(s, os.ModeDir)
	if err != nil {
		fmt.Printf("文件夹:%s 已存在\n", s)
	}

	f, err := os.OpenFile(s+"/"+s+".go", os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	if err != nil {
		// fmt.Println("f crate:", err)
		os.Chdir(s)
		VScodeOpenFlie(s)
		return true
	}
	f.WriteString("package main\n\n\n")
	f.Close()

	os.Chdir(s)

	cmd := exec.Command("go", "mod", "init", s)

	cmd.Run()

	VScodeOpenFlie(s)
	return true

}

func VScodeOpenFlie(s string) {
	fmt.Printf("正在打开 %s.go ...\n", s)
	cmd := exec.Command("code", ".", "-g", s+".go:4")

	cmd.Run()

}

func fileIsExt(s string) bool {
	finfo, err := os.Stat(s)

	if err != nil || finfo.IsDir() {

		CrtateGofile(s)
		return true
	}
	fmt.Printf("文件名:%s 已存在,是文件\n重新输入文件名:", s)
	// if !finfo.IsDir() {
	// 	fmt.Printf("文件名:%s 不是文件夹\n", s)
	// 	return true
	// }

	return false

}
