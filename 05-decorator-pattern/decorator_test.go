package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {

	// 普通文件读写
	fileDataSource := NewFileDataSource("hello.txt")
	// 增强，添加加密功能
	encryptionDataSource := NewEncryptionDataSource(fileDataSource)
	encryptionDataSource.WriteData("hello world")
	data := encryptionDataSource.ReadData()
	fmt.Println(data)

}

func TestDecorator2(t *testing.T) {

	// 普通文件读写
	fileDataSource := NewFileDataSource("hello.txt")
	// 增强，添加压缩功能
	compressDataSource := NewCompressDataSource(fileDataSource)
	compressDataSource.WriteData("hello world")
	data := compressDataSource.ReadData()
	fmt.Println(data)

}

func TestDecorator3(t *testing.T) {

	// 数据库读写
	sqlDataSource := NewSqlDataSource("mysql;localhost:3306;table=hello")
	// 增强，添加压缩功能
	compressDataSource := NewCompressDataSource(sqlDataSource)
	compressDataSource.WriteData("hello world")
	data := compressDataSource.ReadData()
	fmt.Println(data)

}
