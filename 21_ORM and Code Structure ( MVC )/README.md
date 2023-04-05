# (21) ORM and Code Structure (MVC)

ORM (Object-Relational Mapping) adalah sebuah teknik pemrograman yang digunakan untuk memetakan antara data pada basis data relasional dengan model objek pada sebuah aplikasi. Dalam bahasa Go, terdapat beberapa ORM populer seperti GORM, XORM, dan QBS. Sedangkan MVC (Model-View-Controller) adalah sebuah arsitektur perangkat lunak yang terdiri dari tiga bagian yaitu model, view, dan controller.

Dalam pengembangan aplikasi menggunakan bahasa Go, ORM dapat digunakan untuk memudahkan pengembang dalam mengakses basis data dan mengoptimalkan waktu pengembangan aplikasi. Pemilihan ORM yang tepat dapat membantu pengembang dalam menghindari kelemahan pada akses langsung ke basis data, seperti SQL Injection dan akses yang kurang aman.

Dalam struktur MVC, model berfungsi sebagai representasi data dan logika bisnis, view berfungsi sebagai representasi antarmuka pengguna, dan controller berfungsi sebagai pengatur alur data dan input dari pengguna. Penggunaan struktur MVC dapat membantu pengembang dalam mengatur kode program dengan lebih baik dan memisahkan tugas-tugas yang berbeda untuk memudahkan perawatan kode dan pengembangan aplikasi.

Contoh penggunaan ORM dan struktur MVC pada Go adalah sebagai berikut:

1/  Menggunakan GORM sebagai ORM dan struktur MVC

// Model

``

``type User struct {
    ``

    gorm.Model
    Name     string
    Email    string
    Password string
}

``

// Controller



``func GetUser(c *gin.Context) {
    var user User
    id := c.Param("id")``

    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}



// View

``
router.GET("/users/:id", GetUser)
``

Dalam contoh di atas, terlihat penggunaan ORM GORM pada model User. Kemudian, terdapat controller GetUser yang digunakan untuk mengambil data user dari basis data dan menampilkannya pada view. View di sini diwakili oleh router yang menerima permintaan GET pada URL /users/:id.

Dengan menggunakan ORM dan struktur MVC pada Go, pengembang dapat mempercepat waktu pengembangan aplikasi dan meningkatkan kualitas kode yang dihasilkan.

2/ Penggunaan folder project di direktori praktikum

-masuk ke text editor ( penulis menggunakan VSCodium di parrot OS)

-cek PATH directory dari golang dengan perintah

``
go env GOPATH
``

-setelah gopath di ketahui, pindahkan folder project pada direktori GOPATH yang ada.

-jalankan program dengan dengan cara

``
go run main.go
``

-untuk pengguna baru , jika ingin menggunakan folder yang berbeda , silahkan copy paste program pada direktory yang di buat di GOPATH,  kemudian jalankan perintah

``
go mod init nama_modul
``

``
go mod tidy
``

-setelah itu jalankan perintah

``
go get -u gorm.io/gorm
``

``
go get -u gorm.io/driver/mysql
``

``
go get -u github.com/labstack/echo
``

-terakhir , sesuaikan direktori yang di import dalam code yang tersedia di beberapa folder mengikuti folder yang anda buat sendiri.
contoh :

pada folder config , tepatnya di file config.go
terdapat perintah

``
import (" project/controllers")
``

silahkan ganti menjadi

``
import ("nama_foldermu/controllers")
``
