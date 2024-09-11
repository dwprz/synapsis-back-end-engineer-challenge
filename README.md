
# Synopsis Book Management

Synopsis Book Management adalah aplikasi manajemen perpustakaan berbasis microservice yang dirancang untuk mengelola buku, kategori buku, dan pengguna dengan performa yang skalabel dan andal. Aplikasi ini dibangun menggunakan teknologi-teknologi mutakhir untuk memastikan kehandalan dan efisiensi dalam pengelolaan data perpustakaan.


## Features

- Manajemen Buku: Mengelola informasi buku termasuk judul, penulis, ISBN, sinopsis, tahun terbit, stok, dan lokasi.
- Manajemen Kategori Buku: Mengelola berbagai kategori buku yang tersedia.
- Manajemen Pengguna: Mengelola data pengguna, autentikasi, dan otorisasi menggunakan JWT.
- Pencarian dan Rekomendasi: Menyediakan fitur pencarian buku dan rekomendasi berdasarkan popularitas.



## Tech Stack

Aplikasi ini dibangun dengan arsitektur microservice, memisahkan berbagai layanan untuk memastikan skalabilitas dan isolasi layanan. Berikut adalah teknologi dan implementasi utama yang digunakan dalam aplikasi ini:

- Golang: Bahasa pemrograman utama yang digunakan untuk membangun seluruh layanan microservice.

- PostgreSQL: Setiap layanan memiliki database PostgreSQL terpisah untuk mengelola data yang relevan.

- Fiber: Web framework yang digunakan untuk membangun RESTful API dengan performa tinggi.

- gRPC: Digunakan untuk komunikasi cepat dan efisien antara layanan-layanan microservice.

- Redis Caching: Diimplementasikan di User Service untuk meningkatkan kecepatan akses data dan mengurangi beban pada database.

- Circuit Breaker: Diimplementasikan untuk mengurangi error berkepanjangan dengan memutus aliran permintaan ke layanan yang bermasalah.

- Locking dengan SELECT FOR UPDATE: Menggunakan query SELECT FOR UPDATE untuk memastikan konsistensi data saat update stok buku, mencegah race condition dalam pengelolaan stok.

- Unit Testing dengan Testify: Menggunakan framework Testify untuk memastikan kualitas kode melalui pengujian unit yang ekstensif.

- Error Handling dan Recovery: Diimplementasikan untuk menangani error dengan baik dan memastikan server tetap stabil meskipun terjadi panic.

## Thanks For Wathing