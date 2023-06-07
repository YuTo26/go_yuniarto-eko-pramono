# Autentikasi dengan JWT Middleware

Repository ini menyediakan middleware JWT untuk autentikasi dalam aplikasi web. JSON Web Tokens (JWT) adalah metode populer untuk mengirimkan informasi secara aman antara pihak-pihak sebagai objek JSON. Middleware ini dapat diintegrasikan ke dalam aplikasi web Anda untuk menangani autentikasi pengguna menggunakan JWT.

## Cara Kerjanya

1. Pengguna memberikan kredensial mereka (misalnya, nama pengguna dan kata sandi) ke endpoint autentikasi aplikasi Anda.
2. Aplikasi memvalidasi kredensial dan menghasilkan token JWT menggunakan kunci atau rahasia yang telah ditentukan.
3. Token JWT yang dihasilkan dikembalikan kepada pengguna.
4. Untuk permintaan berikutnya, pengguna menyertakan token JWT dalam header `Authorization` permintaan.
5. Middleware JWT aplikasi akan mengintersep permintaan dan memverifikasi keaslian dan integritas token.
6. Jika token valid, pengguna dianggap terautentikasi dan permintaan diproses. Jika tidak, respons error dikembalikan.

## Fitur

- Autentikasi pengguna menggunakan token JWT.
- Generasi dan verifikasi token.
- Penanganan otomatis validasi dan kedaluwarsa token.
- Opsi middleware yang dapat disesuaikan untuk konfigurasi token.
- Integrasi dengan framework web populer (berikan instruksi atau contoh untuk framework spesifik jika ada).

## Instalasi

Untuk menggunakan middleware JWT ini dalam proyek , ikuti langkah-langkah berikut:

1. Pasang paket melalui manajer paket di terminal:

   ```bash
   go get -u github.com/golang-jwt/jwt/v4

2. Import middleware ke dalam project aplikasi :

    ```bash
    import {
    "github.com/golang-jwt/jwt/v4"
    }

3. Buat konstan di dalam mvc :

    ```bash
    const jwtMiddleware = require('jwt-middleware');

4. Gunakan middleware untuk proteksi rute - rute endpoint :

    ```bash
    app.get('/auth', jwtMiddleware.authenticate, (req, res) => {
        //code here
    });

## Opsi Konfigurasi

Middleware JWT menyediakan beberapa opsi konfigurasi untuk menyesuaikan perilakunya. Berikut adalah beberapa opsi yang sering digunakan:

- secret: Kunci rahasia yang digunakan untuk menandatangani dan memverifikasi token JWT.
- algorithm: Algoritma kriptografi yang digunakan untuk menandatangani dan memverifikasi token (misalnya, HS256, RS256).
- ignoreExpiration: Apakah mengabaikan tanggal kedaluwarsa token.
- audience: Audiensi yang dituju untuk token (opsional).
- issuer: Penerbit token (opsional).
- subject: Subjek token (opsional).
