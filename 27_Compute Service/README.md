## COMPUTE SERVICE

### AWS Compute Service - EC2 Instance

AWS EC2 adalah layanan komputasi berbasis cloud yang memungkinkan developer untuk meluncurkan dan mengelola instance virtual di lingkungan cloud AWS. Dengan EC2, developer dapat dengan mudah membuat, mengonfigurasi, dan mengelola instance server sesuai kebutuhan sistem.

#### Fitur Utama

- Elastisitas: developer dapat dengan mudah menyesuaikan kapasitas komputasi sesuai kebutuhan, baik untuk meningkatkan maupun mengurangi jumlah instance.
- Skalabilitas: Layanan ini memungkinkan developer untuk mengelola skalabilitas vertikal (meningkatkan spesifikasi instance) dan skalabilitas horizontal (menambah atau mengurangi jumlah instance).
- Pilihan Sistem Operasi: developer dapat memilih sistem operasi yang sesuai dengan kebutuhan aplikasi , seperti Linux, Windows, dan lain-lain.
- Keamanan: EC2 menyediakan berbagai mekanisme keamanan yang dapat dikonfigurasi, termasuk pengaturan akses, firewall, dan enkripsi data.
- Key Pair (.pem) di AWS : Key Pair (.pem) adalah file kunci yang digunakan untuk mengakses dan mengamankan instance EC2 di AWS. File ini terdiri dari kunci publik dan kunci pribadi.
  1. Kunci Publik: Digunakan untuk mengenkripsi data yang dikirim ke instance EC2.
  2. Kunci Pribadi: Harus dijaga kerahasiaannya dan digunakan untuk mendekripsi data yang diterima oleh instance EC2.
    Penggunaan Key Pair:

     1. Membuat Key Pair: Dapat dilakukan melalui AWS Management Console atau AWS CLI.
     2. Menggunakan Key Pair: Ketika meluncurkan instance EC2, Anda dapat menentukan Key Pair yang digunakan untuk mengaksesnya melalui SSH.
  
### AWS Relational Database Service (RDS)

AWS RDS adalah layanan manajemen database yang memungkinkan developer untuk dengan mudah mengatur, mengelola, dan mengolah database relasional di cloud AWS.
AWS RDS memungkinkan developer untuk menggunakan dan mengelola database relasional yang terdapat di cloud AWS. Layanan ini menawarkan dukungan untuk database populer seperti MySQL, PostgreSQL, Oracle, dan SQL Server.

#### Fitur Utama

- Manajemen Otomatis: AWS RDS mengelola tugas-tugas administrasi yang rumit seperti instalasi, konfigurasi, pemantauan, backup, dan pembaruan.
- Skalabilitas: developer dapat dengan mudah mengatur kapasitas database sesuai dengan kebutuhan aplikasi, baik dalam hal kapasitas penyimpanan maupun kinerja.
- Keamanan: RDS menyediakan fitur keamanan tingkat tinggi, termasuk enkripsi data saat istirahat dan selama perjalanan, pengelolaan akses pengguna, dan tindakan keamanan yang terpadu dengan AWS.
  
### Pointing Domain menggunakan No-IP

No-IP adalah layanan yang memungkinkan developer untuk mengarahkan domain project ke alamat IP dinamis yang berubah-ubah.

##### Langkah-langkah

1. Daftar dan buat akun di No-IP.
2. Pilih dan daftarkan nama domain yang ingin digunakan.
3. Setelah mendaftar, akan diberikan opsi untuk mengatur dynamic DNS.
4. Unduh dan instal aplikasi No-IP di server atau perangkat yang terhubung ke jaringan tempat IP dinamis berada.
5. Masukkan informasi login No-IP ke dalam aplikasi No-IP yang telah diinstal.
6. Konfigurasikan aplikasi No-IP agar mengupdate alamat IP dengan periode tertentu.
7. Di panel kontrol No-IP, atur pengalihan domain ke alamat IP yang diberikan oleh aplikasi No-IP.
