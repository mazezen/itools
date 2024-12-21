package itools

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"net/http"
	"os"
)

// QiNiuFileUpload 上传文件到七牛云
func QiNiuFileUpload(r *http.Request, store string) (string, error) {
	var ak = os.Getenv("QiNiuAK")
	var sk = os.Getenv("QiNiuSk")
	var bucket = os.Getenv("QiNiuBucket")
	var url = os.Getenv("QiuNiuUrl")

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	key := store + fileHeader.Filename
	fileSize := fileHeader.Size

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	url2 := url + ret.Key
	return url2, nil
}

// QiNiuUploadChunk 七牛云分片上传
func QiNiuUploadChunk(r *http.Request, store string) (string, string, error) {
	var ak = os.Getenv("QiNiuAK")
	var sk = os.Getenv("QiNiuSk")
	var bucket = os.Getenv("QiNiuBucket")
	var url = os.Getenv("QiuNiuUrl")

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	resumeUploaderV2 := storage.NewResumeUploaderV2(&cfg)
	upHost, err := resumeUploaderV2.UpHost(ak, bucket)
	if err != nil {
		return "", "", err
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	key := store + fileHeader.Filename
	// 初始化分块上传
	initPartsRet := storage.InitPartsRet{}
	err = resumeUploaderV2.InitParts(context.TODO(), upToken, upHost, bucket, key, true, &initPartsRet)
	if err != nil {
		return "", "", err
	}
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return "", "", err
	}
	fileLen := len(fileContent)
	chunkSize2 := 2 * 1024 * 1024

	num := fileLen / chunkSize2
	if fileLen%chunkSize2 > 0 {
		num++
	}

	// 分块上传
	var uploadPartInfos []storage.UploadPartInfo
	for i := 1; i <= num; i++ {
		partNumber := int64(i)
		fmt.Printf("开始上传第%v片数据", partNumber)

		var partContentBytes []byte
		endSize := i * chunkSize2
		if endSize > fileLen {
			endSize = fileLen
		}
		partContentBytes = fileContent[(i-1)*chunkSize2 : endSize]
		partContentMd5 := Md5S(string(partContentBytes))
		uploadPartsRet := storage.UploadPartsRet{}
		err = resumeUploaderV2.UploadParts(context.TODO(), upToken, upHost, bucket, key, true,
			initPartsRet.UploadID, partNumber, partContentMd5, &uploadPartsRet, bytes.NewReader(partContentBytes),
			len(partContentBytes))
		if err != nil {
			return "", "", err
		}
		uploadPartInfos = append(uploadPartInfos, storage.UploadPartInfo{
			Etag:       uploadPartsRet.Etag,
			PartNumber: partNumber,
		})
		fmt.Printf("结束上传第%d片数据\n", partNumber)
	}

	// 完成上传
	rPutExtra := storage.RputV2Extra{Progresses: uploadPartInfos}
	comletePartRet := storage.PutRet{}
	err = resumeUploaderV2.CompleteParts(context.TODO(), upToken, upHost, &comletePartRet, bucket, key,
		true, initPartsRet.UploadID, &rPutExtra)
	if err != nil {
		return "", "", err
	}

	url2 := url + comletePartRet.Key
	return comletePartRet.Hash, url2, nil
}

// QiNiuResumeUploadFile
// 七牛云文件断点续传
func QiNiuResumeUploadFile(r *http.Request, store string) (string, error) {
	var ak = os.Getenv("QiNiuAK")
	var sk = os.Getenv("QiNiuSk")
	var bucket = os.Getenv("QiNiuBucket")
	var url = os.Getenv("QiuNiuUrl")

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileSize := fileHeader.Size
	key := store + fileHeader.Filename

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	resumeUploaderV2 := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		return "", err
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploaderV2.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	url2 := url + ret.Key
	return url2, nil
}
