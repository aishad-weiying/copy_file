package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// 创建日志文件
	logger, err2 := os.OpenFile("./"+time.Now().Format("2006010215")+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err2 != nil {
		panic(err2)
		return
	}
	defer logger.Close()
	log.SetOutput(logger)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// // 打开父目录,读取其中的子目录
	// ppath := "/Users/weiying/Desktop/backup/"
	// pdir, err := os.OpenFile(ppath, os.O_RDONLY, os.ModeDir)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer pdir.Close()
	// 读取父目录中的子目录,保存到结构体中
	// pfilenames, _ := pdir.Readdir(-1)
	// var pnames []string
	// for _, pname := range pfilenames {
	// 	if pname.IsDir() {
	// 		pnames = append(pnames, pname.Name())
	// 	}
	// }
	// //fmt.Printf("%#v",pnames)
	// for _, value := range pnames {

	// 	首先打开,目录读取目录中的数据
	spath := "/home/wechat_duplicate_removal/data/Main/" + time.Now().Format("2006010215") + "/"
	dir, err := os.OpenFile(spath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		log.Println(err)
	}
	defer dir.Close()

	filenames1, _ := dir.Readdir(-1)
	// 定义切片,用来保存文件名称
	var snames []string
	// 将文件名字保存在切片中
	for _, name := range filenames1 {
		//log.Println(name.Name())
		snames = append(snames, name.Name())
	}
	// 创建目标目录
	dpath := "/home/wechat_data_send_to_ziku/"
	// 判断目录是否存在,如果不存在就创建
	if _, err = os.Stat(dpath); os.IsNotExist(err) {
		err = os.Mkdir(dpath, 0755)
		if err != nil {
			panic(err)
		}
	}
	// 读取目标目录中的文件列表,保存在切片中
	dir2, err := os.OpenFile(dpath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		log.Println(err)
	}
	filenames2, _ := dir2.Readdir(-1)
	// 定义切片,用来保存文件名称
	var dnames []string
	// 将文件名字保存在切片中
	for _, name := range filenames2 {
		//log.Println(name.Name())
		dnames = append(dnames, name.Name())
	}
	// 对比两个切片,如果文件不存在就复制
	duibi(snames, dnames, spath, dpath)
	// }
}

// 对比两个切片,如果文件不存在就复制
func duibi(snames []string, dnames []string, spath string, dpath string) {
	for _, sname := range snames {
		if len(dnames) == 0 {
			err := CopyFile(spath+sname, dpath+sname)
			if err != nil {
				log.Println("复制文件失败", err)
				return
			}
		}
		for i, dname := range dnames {
			if dname == sname {
				//log.Println("文件存在",dname)
				break
			}
			if i == len(dnames)-1 {
				// 复制文件
				err := CopyFile(spath+sname, dpath+sname)
				if err != nil {
					log.Println("复制文件失败", err)
					return
				}
			}
		}
	}
}
func CopyFile(src string, des string) error {
	// 打开源文件
	sfile, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// 创建目标文件
	err = ioutil.WriteFile(des, sfile, 0644)
	if err != nil {
		return err
	}
	log.Println("文件复制完成", src)
	return nil
}
