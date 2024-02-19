# serviceFoodOrganicByGoland
Xây dựng server hệ thống quản lý thực phẩm organic

Công nghệ
- Ngôn ngữ: Golang
- Kiến trúc: Clean architecture
- Hệ quản trị cơ sở dữ liệu: MySql
- Framework/Other: 
  + GinGonic
  + aws - sdk - go (s3)
  + jwt - go
  + gorm
  + godotenv
Chức năng:
  - admin
    + CRUD thực phẩm, loại thực phẩm, user, profile(user), bài viết(blog), đại chỉ(user), bình luận, thương hiệu, image(thực phẩm/loại thực phẩm/user), giới thiệu
    + Bán hàng
    + Check out đơn hàng
  - user
    + Đăng nhập(JWT)/đăng ký, xóa tài khoản, cập nhật thông tin profile(upload/update avatar), cập nhật nhiều địa chỉ giao hàng
    + Viết bài viết, bình luận bài viết
    + Mua hàng, thanh toán
  - Lưu trữ ảnh trên S3
