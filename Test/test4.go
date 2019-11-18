package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"myProject/security"
	"os"
)

var privateKey, publicKey []byte

func init() {
	var err error
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}

func main() {
	var mingwen = "一二三四五六七"
	md5 := security.MD5([]byte(mingwen))
	//MD5打印为16进制字符串
	log.Println(hex.EncodeToString(md5))

	//RSA的内容使用base64打印

	//privateKey:=[]byte("`-----BEGIN PUBLIC KEY-----" +
	//	"MIICWwIBAAKBgQChMrGBlKz7OmQ0RedDcHjMwIkv48ZBOt1DcG1hbx08rKgAlw8iq164LR2GUIo6yD9pm3oBQUPqpI0WslN7ln13KrvskH9DwwYTqfBDfIngYUtwusgBv07E2Jtxes1H7xlsTAg8D7V422EosAsHGmBy7/eq/1u3+17ZgdGu51TfyQIDAQABAoGAQ8VNTX5VT3YYJMXy2a6SivqzcpffhRMbbTv6MImHkDfCliTsLxY/R01oaUy5IMeJqXu9SoPG6wJtcspcQMxfSKyilIJsJjF+D0rl4GHMCCGAm1uANx0/o4vCpAOXtl/FuOWFzb6pdhyMDpoT3FCAqYb2QeMsYJKnTE7d1RgEdRUCQQDRQq0oerLyYLPmTqDruDgDrRgllDL5GUtxs+IHbzDObRlMpPecVm+6D1EOQFbPUury7UWN/1IzWur4Fngf4AwfAkEAxTPZuTw5M9YZWyHCApXvQwLSe26+ygLiuR7YUboJRRy9x2bhEV2sBeseGri04TB3PWvVlVo+kByLY+fUo0MXFwJAJyMCQZzZFO2zF7LC8/MLPtzDtFuIQQBTFNvgvSU1ipXq8mO0D6A22yR8M18jHTlTycVIiesjk4lAgs+o/cUoXwJAGsPp1iFdZjK16E+RpIYzHjZA2S3zyTlRCm0sURNd9Lps66aD/7ZmBbuer2PIcRQB6x06tPW1rhuhs6KgkrQlawJAKqa51LkwEwADVqQz2F+LGl8u/6s7Mw4Y1MHStQD9AcApOsiVDzO+3Pz/gsoArdGTpAV9fHejlYPlTemeBl7pQQ==" +
	//	"-----END PUBLIC KEY-----")
	//publicKey:=[]byte("MIGJAoGBAKEysYGUrPs6ZDRF50NweMzAiS/jxkE63UNwbWFvHTysqACXDyKrXrgtHYZQijrIP2mbegFBQ+qkjRayU3uWfXcqu+yQf0PDBhOp8EN8ieBhS3C6yAG/TsTYm3F6zUfvGWxMCDwPtXjbYSiwCwcaYHLv96r/W7f7XtmB0a7nVN/JAgMBAAE=")

	//privateKey, publicKey, _ := security.GenRSAKey(1024)
	//
	//pr,_:=x509.ParsePKCS1PrivateKey(privateKey)
	//derStream := x509.MarshalPKCS1PrivateKey(pr)
	//block := &pem.Block{
	//	Type:  "RSA PRIVATE KEY",
	//	Bytes: derStream,
	//}
	//file, err := os.Create("private.pem")
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = pem.Encode(file, block)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//
	//
	//pb,_:=x509.ParsePKCS1PublicKey(publicKey)
	//derStream1 := x509.MarshalPKCS1PublicKey(pb)
	//block1 := &pem.Block{
	//	Type:  "PUBLIC KEY",
	//	Bytes: derStream1,
	//}
	//file1, err := os.Create("public.pem")
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = pem.Encode(file1, block1)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	log.Println("rsa私钥:\t", base64.StdEncoding.EncodeToString(privateKey))
	log.Println("rsa公钥:\t", base64.StdEncoding.EncodeToString(publicKey))

	miwen, err := security.RsaEncrypt([]byte(mingwen), publicKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("加密后：\t", base64.StdEncoding.EncodeToString(miwen))

	jiemi, err := security.RsaDecrypt(miwen, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("解密后：\t", string(jiemi))
}

/**
 * @brief  获取RSA公钥长度
 * @param[in]       PubKey				    RSA公钥
 * @return   成功返回 RSA公钥长度，失败返回error	错误信息
 */
func GetPubKeyLen(PubKey []byte) (int, error) {
	if PubKey == nil {
		return 0, errors.New("input arguments error")
	}

	block, _ := pem.Decode(PubKey)
	if block == nil {
		return 0, errors.New("public rsaKey error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return 0, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return pub.N.BitLen(), nil
}

/**
 * @brief  获取RSA私钥长度
 * @param[in]       PriKey				    RSA私钥
 * @return   成功返回 RSA私钥长度，失败返回error	错误信息
 */
func GetPriKeyLen(PriKey []byte) (int, error) {
	if PriKey == nil {
		return 0, errors.New("input arguments error")
	}

	block, _ := pem.Decode(PriKey)
	if block == nil {
		return 0, errors.New("private rsaKey error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return 0, err
	}

	return priv.N.BitLen(), nil
}
