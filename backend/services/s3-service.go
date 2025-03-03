package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3 servis yapısı
type S3Service struct {
	Client   *s3.Client
	Uploader *manager.Uploader
	Bucket   string
}

// Yeni S3 servisi başlat
func NewS3Service() (*S3Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Printf("S3 konfigürasyon yüklenemedi: %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	bucketName := os.Getenv("S3_BUCKET_NAME")
	if bucketName == "" {
		return nil, fmt.Errorf("S3_BUCKET_NAME ortam değişkeni tanımlanmalı")
	}

	return &S3Service{
		Client:   client,
		Uploader: uploader,
		Bucket:   bucketName,
	}, nil
}

// S3'ye dosya yükleme
func (s *S3Service) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		log.Printf("S3 Upload Hatası - Dosya açılamadı: %v", err)
		return "", fmt.Errorf("dosya açılamadı: %w", err)
	}
	defer src.Close()

	// Dosya adını belirleme
	objectKey := fmt.Sprintf("%s/%d-%s", folder, time.Now().Unix(), file.Filename)

	// Dosya içeriğini okuyarak S3'ye yükle
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, src)
	if err != nil {
		log.Printf("S3 Upload Hatası - Dosya kopyalama hatası: %v", err)
		return "", fmt.Errorf("dosya kopyalama hatası: %w", err)
	}

	log.Printf("S3'ye yükleme işlemi başlıyor: %s", objectKey)

	_, err = s.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(buffer.Bytes()),
		ACL:    "private", // Dosyanın gizlilik ayarı
	})

	if err != nil {
		log.Printf("S3 Upload Hatası - S3 yükleme başarısız: %v", err)
		return "", fmt.Errorf("S3 yükleme hatası: %w", err)
	}

	log.Printf("S3'ye başarıyla yüklendi: %s", objectKey)

	return objectKey, nil
}

// S3'ten dosya silme
func (s *S3Service) DeleteFile(objectKey string) error {
	_, err := s.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("S3 dosya silme hatası: %w", err)
	}
	return nil
}

// S3'ten Presigned URL alma (Dosyayı güvenli erişim için)
func (s *S3Service) GetPresignedURL(objectKey string, expiry time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(s.Client)
	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", fmt.Errorf("Presigned URL oluşturulamadı: %w", err)
	}
	return req.URL, nil
}
