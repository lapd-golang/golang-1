package main

import (
	"fmt"
	_"io"
	"io/ioutil"
	"path"

)

func fileList(dir string, ftype string) ( *map[int] string,  error) {
	fdir , err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	//var list map[int] string = make(map[int]string)
	list :=  map[int] string{} 
	for n, fi := range fdir {
		fname := fi.Name()
		if ext := path.Ext(fname); ext != ftype{
			continue
		}
		fname_rel := dir + "/"+fname
		list[n] = fname_rel
	}
	fmt.Printf("list addr %p\n", &list)
	
	return &list, nil
}

func main(){
	flist, err := fileList(".", ".go") 
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("main list addr %p\n", flist)
	for i, ff := range *flist {
		fmt.Println(i, ":", ff)
	}
}
