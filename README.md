# *KFS-Backend*

### 📌 Proje Amacı ###

KFS-Backend, girişimcilerin projelerini tanıtabilecekleri ve yatırımcıların bu projelere güvenli bir şekilde yatırım yapabilecekleri bir kitle fonlama yönetim yazılımı için geliştirilen bir backend uygulamasıdır. Bu sistem, kullanıcı güvenliği, veri gizliliği ve yasal uyumluluk esas alınarak tasarlanmıştır. Aynı zamanda girişimciler ve yatırımcılar arasında etkin bir iletişim ve süreç yönetimi sağlayacak özellikler sunmaktadır.

Proje, iki ana panelden oluşmaktadır:

Admin Paneli: Kullanıcı yönetimi, doğrulamalar ve genel sistem kontrolü için kullanılır.

Kullanıcı Paneli: Girişimcilerin projelerini oluşturup yönetebileceği, yatırımcıların projelere yatırım yapabileceği kullanıcı arayüzüdür.



### 🚀 Teknoloji Listesi ###

| Teknoloji| Açıklama|
|----------|----------|
| Back-end  |Golang, Fiber  | 
| Front-end  | Next.js, TypeScript, Tailwind CSS, Shadcn UI |
| Database  |PostgreSQL (SupaBase)  |
| State Management |TanStack Query  |
| Authentication |JWT, Cookie  |
| API Documentation |TPostman  |
| ORM |GORM |

### 📖 Özellikler ###

- Güvenli ve yasalara uyumlu bir kitle fonlama platformu

- Girişimciler için proje oluşturma ve yönetim arayüzü

- Yatırımcılar için yatırım süreci ve fon yönetimi

- Admin paneli üzerinden kullanıcı ve proje yönetimi

- Detaylı rol ve yetkilendirme mekanizması

- İleri düzey loglama ve denetim mekanizmaları

- AML ve KYC süreçlerini destekleyen güvenlik önlemleri

- Esnek yatırım ve fonlama yönetimi

### 📂 Proje Kurulumu ###

1️⃣ Depoyu Klonlayın

`git clone https://github.com/kullaniciadi/kfs-backend.git
cd kfs-backend`

2️⃣ Çevresel Değişkenleri Tanımlayın

.env dosyanızı oluşturun ve gerekli değişkenleri ekleyin:

`DATABASE_URL=postgres://username:password@localhost:5432/kfs_db
JWT_SECRET=your_secret_key`

3️⃣ Bağımlılıkları Yükleyin

`go mod tidy`

4️⃣ Veritabanını Başlatın

## PostgreSQL çalıştırın ##
`docker-compose up -d`

5️⃣ Sunucuyu Başlatın

`go run main.go`

### 🛠 API Kullanımı ###

| Yöntem| Endpoint|Açıklama|
|----------|----------|----------|
| POST |/api/auth/register | Kullanıcı kaydı oluşturur|
