# Intro Echo Golang

Echo adalah sebuah framework web yang ditulis dalam bahasa Go dan digunakan untuk membangun aplikasi web dan RESTful API. Echo termasuk salah satu framework yang cepat dan ringan. Echo memanfaatkan fitur-fitur bahasa Go yang modern seperti goroutine dan channel sehingga dapat meningkatkan kinerja dalam memproses HTTP request.

Dalam membuat aplikasi atau RESTful API dengan Echo, terdapat beberapa method HTTP yang umum digunakan yaitu:

1/
    POST ==>
    Method HTTP ini digunakan untuk mengirim data ke server untuk membuat resource baru. Dalam Echo, untuk menangani request POST dapat dilakukan dengan menggunakan e.POST(). Contoh penggunaan method POST dalam Echo:

``
e.POST("/users", CreateUserController)
``

2/
    GET ==>
    Method HTTP ini digunakan untuk mengambil data dari server. Dalam Echo, untuk menangani request GET dapat dilakukan dengan menggunakan e.GET(). Contoh penggunaan method GET dalam Echo:

``
e.GET("/users", GetUsersController)
``

3/
    PUT ==>
    Method HTTP ini digunakan untuk mengirim data ke server untuk memperbarui resource yang sudah ada. Dalam Echo, untuk menangani request PUT dapat dilakukan dengan menggunakan e.PUT(). Contoh penggunaan method PUT dalam Echo:

``
e.PUT("/users/:id", UpdateUserController)
``

4/
    DELETE ==>
    Method HTTP ini digunakan untuk menghapus resource yang sudah ada pada server. Dalam Echo, untuk menangani request DELETE dapat dilakukan dengan menggunakan e.DELETE(). Contoh penggunaan method DELETE dalam Echo:

``
e.DELETE("/users/:id", DeleteUserController)
``

Dalam implementasi method HTTP pada Echo, kita dapat membuat controller atau handler yang berfungsi untuk menangani request dari client. Setelah itu, kita bisa menghubungkan controller tersebut dengan route menggunakan method HTTP yang sesuai.
