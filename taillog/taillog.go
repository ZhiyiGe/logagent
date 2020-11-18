package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tails *tail.Tail
)

func Init(filename string) (err error) {
	conf := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在报错
		Poll:      true,
	}
	tails, err = tail.TailFile(filename, conf)
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	return
}
func Readlog() chan *tail.Line{
	return tails.Lines
}
