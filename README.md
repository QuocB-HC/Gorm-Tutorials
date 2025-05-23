# GormTutorial

## Mô tả

Dự án này là một tutorial để tìm hiểu về Gorm.

## Công Nghệ Sử Dụng

- Golang

## Đóng Góp

Nếu bạn muốn đóng góp, hãy tạo một pull request hoặc mở issue mới.

## Cài Đặt và Chạy Dự Án

### Yêu Cầu Hệ Thống

- Go 1.18 trở lên
- Trình soạn thảo như VS Code, GoLand hoặc bất kỳ IDE nào bạn thích

### Cài đặt

Clone Repository
```sh
git clone https://github.com/QuocB-HC/Gorm-Tutorials.git
cd Gorm-Tutorials
```

Chọn Dự Án
```sh
cd project-name
```

Chạy Dự Án
```sh
go run main.go
```

Đối với Connect-Database trước khi chạy nên tạo postgres với docker
```sh
docker run --name postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres
```
