## Setup Proyek  

### 1. Inisialisasi Modul  
Jalankan perintah berikut untuk menginisialisasi proyek dan mengunduh dependensi:  
```sh
go mod init MinIO
go get github.com/minio/minio-go/v7
go get github.com/minio/minio-go/v7/pkg/credentials
go mod tidy
```

### 2. Menjalankan Aplikasi  
Gunakan perintah berikut untuk menjalankan aplikasi:  
```sh
go run main.go
```

## Contoh Output  
```
Successfully connected to MinIO  
Files in bucket:  
example.txt  
File uploaded successfully  
File downloaded successfully  
Presigned URL: http://203.194.113.6:9000/smk-telkom/example.txt  
Public URL: http://203.194.113.6:9000/smk-telkom/example.txt  
File deleted successfully  
```

