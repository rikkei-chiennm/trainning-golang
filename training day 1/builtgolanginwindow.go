package main

import "fmt"

func main() {
	/*
					cach cai dat golang tren cac nen tang khac nhau
					https://topdev.vn/blog/hoc-golang-tu-con-so-0-phan-1-cai-dat-golang-tren-linux-va-windows/#cai-dat-golang-tren-ubuntu
					còn về viết chương trình đầu tiên thì nó như phần ở dưới đây
					- package
					- import thư viện
					- main
					- code trong main
				- buid golang in linux:
					+ setting Go runtime: ~ sudo apt-get update
										~ sudo apt-get -y upgrade
					+ down go bằng command:  ~ curl -O https://storage.googleapis.com/golang/go1.12.6.linux-amd64.tar.gz
			        + giải nén và cài đặt : ~ sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz
								          ~ export PATH=$PATH:/usr/local/go/bin
		        - config CPU when build GOLang
					+ Go is using a lot of CPU because it's trying to build as fast as possible, like any other compiler.
					You can control the number of processes Go is using by setting the GOMAXPROCS environment
					variable. Such as GOMAXPROCS=1 go get ... to limit Go to only use 1 process (and thus only 1
					CPU core). This however does not affect the number of processes used by external compilers that cgo may invoke.

					If you need further CPU control, on Unix based systems you can use the nice command to change
					the priority of processes such that other programs have higher CPU priority, making your computer
					less sluggish.



	*/

	fmt.Println("hello word")

}
