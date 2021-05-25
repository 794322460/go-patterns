package decorator

import (
	"fmt"
	"strings"
)

/*
	装饰模式使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能
	*设计思想
		1.声明基础接口DataSource
		2.定义基础结构体，并实现接口继承。基本文件读写
		3.装饰结构体中, 采用匿名组合的方式将接口Component作为结构体的属性。扩展（加密文件读写，压缩文件读写等）

与代理模式の区别：代理模式客户端其实不知道自己用的对象是被代理过的，而装饰者模式是客户端自己手动包装的，主动索要增强功能。
	装饰者 ==> 穿不同的马甲就有不同的功能。
*/
type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// =========================================================================== //

// 基本文件读写
type FileDataSource struct {
	Filename string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{Filename: filename}
}

func (dataSource *FileDataSource) WriteData(data string) {
	fmt.Println("基本文件读写：写入内容=", data)
}
func (dataSource *FileDataSource) ReadData() string {
	fmt.Println("基本文件读写：读取文件名称=", dataSource.Filename)
	return "读取文件" + dataSource.Filename + "..."
}

// 数据库读写
type SqlDataSource struct {
	Dsn string
}

func NewSqlDataSource(dsn string) *SqlDataSource {
	return &SqlDataSource{Dsn: dsn}
}

func (dataSource *SqlDataSource) WriteData(data string) {
	fmt.Println("数据库读写：写入内容=", data)
}
func (dataSource *SqlDataSource) ReadData() string {
	fmt.Println("数据库读写：读取Dsn=", dataSource.Dsn)
	return "读取文件" + dataSource.Dsn + "..."
}

// =========================================================================== //

// 马甲：数据源读写加密
type EncryptionDataSource struct {
	DataSource
}

func NewEncryptionDataSource(dataSource DataSource) DataSource {
	return &EncryptionDataSource{DataSource: dataSource}
}

func (dataSource *EncryptionDataSource) WriteData(data string) {
	fmt.Println("EncryptionDataSource：加密数据，data=", data)
	encryptionData := dataSource.encryption(data)
	fmt.Println("EncryptionDataSource：加密后，encryptionData=", encryptionData)
	fmt.Println("EncryptionDataSource：调用原始对象WriteData方法...")
	dataSource.DataSource.WriteData(encryptionData)
}
func (dataSource *EncryptionDataSource) ReadData() string {
	fmt.Println("EncryptionDataSource：调用原始对象ReadData方法...")
	data := dataSource.DataSource.ReadData()
	fmt.Println("EncryptionDataSource：解密数据，data=", data)
	decryption := dataSource.decryption(data)
	fmt.Println("EncryptionDataSource：解密后，decryptionData=", decryption)
	return decryption
}

func (dataSource *EncryptionDataSource) encryption(data string) string {
	fmt.Println("EncryptionDataSource：准备加密字符串...", data)
	return "###" + data + "###"
}
func (dataSource *EncryptionDataSource) decryption(data string) string {
	fmt.Println("EncryptionDataSource：准备解密字符串...", data)
	return strings.ReplaceAll(data, "###", "")
}

// 马甲：数据源读写压缩
type CompressDataSource struct {
	DataSource
}

func NewCompressDataSource(datasource DataSource) *CompressDataSource {
	return &CompressDataSource{DataSource: datasource}
}

func (dataSource *CompressDataSource) WriteData(data string) {
	fmt.Println("CompressDataSource：压缩数据，data=", data)
	compressData := dataSource.compress(data)
	fmt.Println("CompressDataSource：压缩后，compressData=", compressData)
	fmt.Println("CompressDataSource：调用原始对象WriteData方法...")
	dataSource.DataSource.WriteData(compressData)
}
func (dataSource *CompressDataSource) ReadData() string {
	fmt.Println("CompressDataSource：调用原始对象ReadData方法...")
	data := dataSource.DataSource.ReadData()
	fmt.Println("CompressDataSource：解压缩数据，data=", data)
	decompressData := dataSource.decompress(data)
	fmt.Println("CompressDataSource：解压缩后，decompressData=", decompressData)
	return decompressData
}

func (dataSource *CompressDataSource) compress(data string) string {
	fmt.Println("CompressDataSource：准备压缩字符串...", data)
	return "XXX" + data + "XXX"
}
func (dataSource *CompressDataSource) decompress(data string) string {
	fmt.Println("CompressDataSource：准备解压缩字符串...", data)
	return strings.ReplaceAll(data, "XXX", "")
}
