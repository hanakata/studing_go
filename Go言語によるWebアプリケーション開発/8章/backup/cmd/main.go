package main

import (
	"errors"
	"flag"
	"log"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()
	// var (
	// 	dbpath = flag.String("db", "./backupdate", "データベースのディレクトリへのパス")
	// )
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("エラー;コマンドを指定してください")
		return
	}
	// db, err := filedb.Dial(*dbpath)
	// if err != nil{
	// 	fatalErr = err
	// 	return
	// }
	// defer db.Close()
	// col,err := db.C("paths")
	// if err != nil{
	// 	fatalErr = err
	// 	return
	// }
}
