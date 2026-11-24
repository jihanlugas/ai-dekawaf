# Software Requirements Specification for Security Firewall: F5 2DC

**Version:** 1.17  
**Prepared by:** Fairuz Rafida  
**Date:** 7 Oktober 2026  

---

## Table of Contents

- 1. Introduction
  - 1.1 Purpose of the Document
  - 1.2 Scope
  - 1.3 Definitions and Acronyms
- 2. Description
  - 2.1 Product Overview
  - 2.2 Product Features
  - 2.3 User Documentation
- 3. Interface Requirement
  - 3.1 System Architecture
  - 3.2 Software Interface and Stack
  - 3.3 iControl REST API Fundamentals
- 4. Pricing & Charge
- 5. Functional Requirement
  - 5.1 Functional Requirement Table
- 6. Non-Functional Requirements
- 7. Appendix
  - 7.1 Mockup
    - 7.1.1 Security Profile - List (Empty State)
    - 7.1.2 Security Profile - Create Profile Wizard
    - 7.1.3 Security Profile - List (Populated)
    - 7.1.4 Security Profile - Overview
    - 7.1.5 Security Profile - Edit (General Settings & WAF Mode)
    - 7.1.6 Security Profile - Upload/Replace SSL Certificate
    - 7.1.7 Access Control - Whitelist IP (Empty State)
    - 7.1.8 Access Control - Whitelist IP (User Admin, Add & Smart Resolve)
  - 7.2 Akun perangkat Virtual Server F5
  - 7.3 API List
  - 7.4 Tes Case API
  - 7.5 Working Instruction in F5
    - 7.5.1 A. Onboarding Baru (dipicu saat User Admin membuat Security Profile baru mockup 7.1.2 Create Profile Wizard)
    - 7.5.2 C. Renewal SSL Certificate
    - 7.5.3 D. Whitelist IP (Learning Period 7 Hari & Support ID)
    - 7.5.4 E. Konfigurasi Bot Defense (Bot Protection)
  - 7.6 Server Technologies Description

---

## Revision History

| Name | Date | Reason For Changes | Version |
| :--- | :--- | :--- | :--- |
| Fairuz Rafida | 21/11/2025 | Draf Awal | 1.0 |
| Fairuz Rafida | 16/02/2026 | Menambahkan Referensi | 1.1 |
| Fairuz Rafida | 11/03/2026 | Updated: Integration of F5 Integration content (System Architecture, API, Workflows, Extended NFR) | 1.2 |
| Fairuz Rafida | 15/06/2026 | Menambahkan Prasyarat untuk Pengembangan menggunakan F5 | 1.3 |
| Fairuz Rafida | 28/07/2026 | Menambahkan bagian Pricing & Charge | 1.4 |
| Fairuz Rafida | 31/08/2026 | Revisi | 1.5 |
| Fairuz Rafida | 16/09/2026 | Sinkronisasi dengan WI Konsolidasi F5 BIG-IP WAF - Deka WAF V2.0 (Bot Defense): revisi FR-DWAF-0008 Configure Bot Protection agar sesuai prosedur F5 Bot Defense (Logging Profile, Profile Template, Virtual Server attachment) dan penambahan API Namespace `/mgmt/tm/security/bot-defense/profile` & `/mgmt/tm/security/log/profile` (diverifikasi terhadap dokumentasi resmi F5 `techdocs.f5.com`/`clouddocs.f5.com`) | 1.6 |
| Fairuz Rafida | 16/09/2026 | Audit menyeluruh kesesuaian SRS vs Mockup (15 layar) vs WI Konsolidasi V2.0: (1) perbaikan duplikasi/kontradiksi FR-DWAF-0008 vs FR-DWAF-0030 terkait Bot Protection Profile Template, (2) penambahan definisi Node/Pool/Bot Defense/GTM-GSLB/Wide IP pada Tabel 1, (3) koreksi klaim breadcrumb 3-level pada FR-DWAF-0012 agar sesuai mockup (2-level). Open item belum diselesaikan (menunggu keputusan desain): field Service Port pada Create Profile/FR-DWAF-0002 & FR-DWAF-0025 (WI A.2), dan halaman Monitoring pada sidebar Superadmin belum punya FR/mockup. | 1.7 |
| Fairuz Rafida | 16/09/2026 | Penambahan acceptance criteria FR-DWAF-0002 (keterkaitan Security Profile dengan 1 Virtual Server sesuai pricing per Virtual Server di Bab 4) dan FR-DWAF-0009/FR-DWAF-0010 (penegasan akses Add/Delete Whitelist & Blacklist hanya untuk role Super Admin, User Admin bersifat read-only) | 1.8 |
| Fairuz Rafida | 18/09/2026 | Selaraskan SRS dengan mockup Deka WAF terkini (`mockup-waf-new.pdf`) dan WI Konsolidasi F5 BIG-IP WAF V2.0, dalam satu revisi gabungan: (1) hapus/descope Managed Rules, Rate Limiting, Blacklist IP, dan IP Conflict Detection - tidak ada lagi di mockup terkini; (2) sederhanakan Bot Protection menjadi satu checkbox Portal-side, dengan klarifikasi bahwa backend memakai 1 Virtual Server bersama + 1 Bot Defense Profile tetap yang di-setup sekali oleh NOC/Cloudeka (bukan per pelanggan), termasuk klarifikasi bahwa VS ini tidak dihitung dalam skema pricing per VS (Bab 4, Nota Dinas No.047/ND/22500/2019); (3) tambah field Port (dropdown 80-HTTP/443-HTTPS) ke FR-DWAF-0002, FR-DWAF-0018, FR-DWAF-0033, dan Appendix 7.1.2, menyusul penambahan field ini di `mockup-waf-new.pdf`; (4) Enforcement Mode kini wajib dipilih eksplisit (tidak ada default Blocking); (5) Overview: panel Protection digabung ke card Profile Configuration 3 kolom, diperbarui di FR-DWAF-0003, FR-DWAF-0014, dan Appendix 7.1.4; (6) tarik panel "Recent Security Events" dari scope (FR-DWAF-0015, Removed - Descoped) karena belum ada kepastian rencana ke depan; (7) tambah FR-DWAF-0034 (Server Technologies, Proposed); (8) ganti 9 screenshot mockup di Appendix 7.1 dengan versi terbaru (sebelumnya berasal dari desain lama yang masih ada Managed Rules/Blacklist/menu Protection-Access Control-Monitoring terpisah); (9) hapus subbagian Appendix 7.1.7 (Protection Landing Page) dan 7.1.11-7.1.13 (Blacklist IP, IP Conflict Detected). (10) Perjelas skema billing di FR-DWAF-0002 dan Bab 4 Pricing & Charge: unit tarif adalah 1 Virtual Server (bukan per domain) sesuai Nota Dinas No.047/ND/22500/2019 — domain hanya pemicu pembuatan VS; platform F5 bersifat multitenant (1 perangkat fisik melayani banyak pelanggan via partisi/VS terisolasi), tarif tetap dihitung per VS. Juga menghapus 1 heading kosong yang nyasar antara bagian Non-Functional Requirements dan Appendix (berpotensi membuat baris kosong di Daftar Isi). (11) Perjelas batas scope Load Balancing: hapus klaim "load balancing" sebagai layanan yang dikelola pelanggan di Section 1.1 (Purpose), pertegas bullet Node/Pool di In-Scope sebagai infrastruktur pendukung WAF saja, dan tambah bullet Out-of-Scope eksplisit untuk L7 Load Balancer & GLB sebagai produk mandiri (beda SKU/harga per Nota Dinas No.047/ND/22500/2019, di luar cakupan dokumen ini). (12) Terjemahkan seluruh isi dokumen ke Bahasa Indonesia (Product Overview/Features, System Architecture, Technology Stack, Parameters, judul 34 Functional Requirement pada Tabel 5-6, label Non-Functional Requirements, dan 22 User Story pada Tabel 8-16) — heading/judul besar (Introduction, Description, Functional Requirement, dsb.), nama modul resmi F5, istilah teknis baku (Node, Pool, Virtual Server, dropdown, checkbox, dll.), serta string UI literal dari mockup tetap dalam Bahasa Inggris sesuai aslinya. (13) Hapus total 7 baris FR yang sebelumnya hanya ditandai "Removed - Descoped" (FR-DWAF-0006, 0007, 0010, 0011, 0015, 0026, 0030) dari Tabel Functional Requirement - tidak lagi disisakan sebagai baris tombstone, mengikuti arahan bahwa fitur yang sudah dihapus tidak perlu tetap dituliskan. Referensi silang yang jadi menggantung akibat penghapusan ini (deskripsi FR-DWAF-0014 dan kolom Functional Requirement pada Tabel 11) sudah dibersihkan. Penomoran FR-DWAF-0002 s/d FR-DWAF-0034 sengaja TIDAK dirapikan ulang (ada celah nomor) untuk menghindari risiko salah-referensi di seluruh dokumen dan file lain (Test Case Excel, dokumen Spesifikasi Perubahan) yang masih memakai nomor asli. (14) Penambahan Working Instruction di 7.5 | 1.9 |
| Fairuz Rafida | 23/09/2026 | Revisi FR-DWAF-0008 Konfigurasi Bot Protection: desain sebelumnya (1 Virtual Server dipakai BERSAMA oleh SEMUA pelanggan SP Cloudeka untuk Bot Defense) tidak dapat diimplementasikan pada F5, karena setiap Virtual Server hanya dapat melayani 1 domain/Node-Pool tertentu (WI bagian A.6) sehingga 1 VS tidak mungkin dipakai lintas banyak pelanggan (multitenant) sekaligus. Diubah menjadi: Bot Defense Profile & Logging Profile (tetap SATU objek standar `template=Balanced`, `mode=Blocking`, dibuat sekali oleh tim NOC/Cloudeka) di-attach/detach oleh Portal ke Virtual Server MILIK MASING-MASING PELANGGAN (VS yang sama dengan WAF Policy pada FR-DWAF-0002) setiap kali checkbox Bot Protection dicentang/dilepas - bukan lagi ke 1 VS bersama. Klarifikasi billing turut diperbarui: mengaktifkan Bot Protection tetap TIDAK menambah unit VS billing terpisah, karena memakai VS pelanggan yang sudah ada. | 1.10 |
| Fairuz Rafida | 28/09/2026 | Menghapus mekanisme approval Super Admin untuk pembuatan/perubahan Security Profile (FR-DWAF-0031 dihapus total dari tabel FR, mengikuti preseden revisi v1.9): setelah User Admin membuat/mengubah Security Profile melalui Service Portal, sistem otomatis meregistrasikan/deploy konfigurasi ke F5 tanpa menunggu approval - menyelaraskan dengan mockup terkini yang tidak memiliki halaman/status approval untuk Security Profile (status profile hanya Active/Inactive/Error, FR-DWAF-0001). Approval workflow oleh Super Admin TETAP berlaku khusus untuk penambahan IP Whitelist (FR-DWAF-0032, tidak berubah). Pembaruan mengikuti: definisi Super Admin/User Admin (Tabel 1), Scope 1.2, System Architecture Layer 2, acceptance criteria FR-DWAF-0002 dan FR-DWAF-0019, serta penghilangan cross-reference ke FR-DWAF-0031 pada FR-DWAF-0032. Selain itu, merapikan gap penomoran Gambar pada Appendix 7.1 (sebelumnya loncat dari Gambar 7 ke Gambar 9 tanpa Gambar 8) - Gambar 9/10/11 diganti menjadi Gambar 8/9/10. | 1.11 |
| Fairuz Rafida | 28/09/2026 | Menambahkan pemetaan detail pada 7.5 Working Instruction in F5: setiap langkah WI Konsolidasi F5 BIG-IP WAF V2.0 (skenario A-E) dipetakan terhadap Service Portal Cloudeka sesuai `mockup-waf-new.pdf` dan tabel FR Bab 5-6, dengan penandaan OTOMATIS (backend/Automation Engine), INPUT User Admin/Super Admin, atau MANUAL oleh NOC di F5 (untuk skenario B, GSLB/FR-DWAF-0027, Server Technologies/FR-DWAF-0034, dan setup awal Bot Defense Profile/FR-DWAF-0008). Turut dicatat 2 gap yang belum punya acceptance criteria eksplisit: A.7 Enable Logging dan Smart Resolve pada FR-DWAF-0009. | 1.12 |
| Fairuz Rafida | 29/09/2026 | Menyelaraskan SRS dengan pembaruan mockup Deka WAF (`mockup-deka-waf-new2.pdf`): field "Server Technologies" (multi-select, kategori Web Server / Language & Framework / Database) kini tersedia pada Step 2 (Configuration) wizard Create Security Profile, menggantikan field "Application Language" (single-select). FR-DWAF-0034 diubah dari status Proposed menjadi sudah diimplementasikan, dengan pembaruan pada FR-DWAF-0002, FR-DWAF-0003, FR-DWAF-0018, FR-DWAF-0034, deskripsi wizard 3-langkah, dan pemetaan Working Instruction 7.5 (A.8 - Server Technologies, dari MANUAL oleh NOC menjadi INPUT User Admin). Catatan: halaman Overview/Detail pada mockup belum menampilkan "Server Technologies" secara eksplisit (header masih "App Languange" dan card Profile Configuration belum punya baris ini) - SRS diselaraskan dengan asumsi pembaruan tersebut akan menyusul. | 1.13 |
| Fairuz Rafida | 29/09/2026 | Konfirmasi lanjutan: `mockup-deka-waf-new2.pdf` sudah diperbarui lebih lanjut - halaman Overview dan Edit Security Profile kini turut menampilkan field "Server Technologies" pada header profile (menggantikan "App Language"), tampil "-" bila belum ada nilai dipilih, meski card "Profile Configuration" belum memiliki baris terpisah untuk field ini. FR-DWAF-0034 diperbarui dengan acceptance criteria tambahan untuk Overview & Edit, deskripsinya diperluas, dan satu paragraf Acceptance Criteria basi peninggalan revisi lama ("Belum diimplementasikan... menunggu keputusan desain") dihapus karena sudah tidak relevan. Catatan gap pada revisi v1.13 mengenai halaman Overview dinyatakan selesai (resolved). | 1.14 |
| Fairuz Rafida | 30/09/2026 | Menghapus penjelasan langkah "A.5 SSL Profile (Server)" dari `WI Konsolidasi F5 BIG-IP WAF - Deka WAF V2.0.docx` (heading beserta instruksinya), karena langkah menambahkan SSL Profile (Server) terpisah terbukti menyebabkan error saat pelanggan mengakses website (sudah dicatat sebagai catatan pada heading tersebut sebelum akhirnya dihapus total). Field "SSL Profile (Server): `serverssl-insecure-compatible`" tetap ada sebagai bagian dari langkah A.6 Membuat Virtual Server. Penomoran WI (A.6 dst) sengaja tidak dirapikan ulang. Referensi pada tabel pemetaan Working Instruction 7.5 untuk baris "A.5 SSL Profile (Server)" diperbarui dari "WI V2.0 bagian A.5" menjadi "WI V2.0 bagian A.6". | 1.15 |
| Fairuz Rafida | 05/10/2026 | Menyelaraskan mekanisme Whitelist IP: (1) Whitelist IP bukan hasil pengajuan pelanggan melainkan hasil learning period 7 hari sejak Security Profile Active — IP yang mengakses domain/SSL profile otomatis tercatat, dan setelah 7 hari hanya IP tersebut yang diizinkan (IP lain otomatis diblok); (2) FR-DWAF-0032 (Approval Workflow Whitelist IP) dihapus total — tidak ada lagi approval/Pending Approval di SRS (nomor FR tidak dipakai ulang); (3) tambah FR-DWAF-0035 (Learning Period & Auto-Whitelist IP), US-059 s.d. US-061, dan istilah Learning Period pada Tabel 1; (4) FR-DWAF-0009 ditulis ulang: User Admin read-only dan mengajukan ticket untuk penghapusan IP yang tidak seharusnya ada; Super Admin menambah/menghapus IP atas ticket dan langsung diterapkan ke F5 tanpa approval; (5) pembaruan definisi Super Admin/User Admin, Scope 1.2, Product Features, System Architecture, FR-DWAF-0017, FR-DWAF-0019, US-016/US-018, deskripsi Appendix 7.1.7-7.1.9, serta pemetaan WI 7.5 skenario D. (6) tambah mekanisme teknis learning period & pemblokiran di F5 - Log Profile `all-request` dan event log ASM untuk pencatatan IP selama 7 hari, lalu Address List + Firewall Policy AFM (allow Whitelist + default deny) yang di-enforce ke Virtual Server; ditambahkan pada FR-DWAF-0035, FR-DWAF-0009 (Smart Resolve), Tabel 4, Tabel 5 (API List), dan pemetaan WI 7.5 skenario D; Whitelist kini bermakna allowlist akses (inspeksi WAF tetap berlaku, tidak lagi bypass); catatan gap mockup dihapus karena penjelasan learning period ada di User Guide. (7) tambah 10 test case baru (US-062 s.d. US-071) pada Test Case xlsx dan perbarui bagian 7.4 Tes Case API; baris US-016/017/018 pada xlsx diberi catatan karena enforcement Whitelist pindah ke AFM Address List. (8) FR-DWAF-0034 dan pemetaan WI 7.5 diperjelas: Application Language kini ditetapkan otomatis ke Auto Detect oleh backend (WI A.8) setelah field tersebut digantikan Server Technologies; ditambahkan Appendix 7.6 Server Technologies Description (57 teknologi dari dokumen F5 yang termasuk 3 kategori mockup: Web Server, Language & Framework, Database; teknologi kategori lain dari dokumen F5 tidak dimasukkan). (9) Batas minimal Profile Name diselaraskan menjadi 3 karakter pada FR-DWAF-0002 dan FR-DWAF-0018 mengikuti mockup (sebelumnya tidak konsisten: 8 vs 3). Dokumentasi F5 tidak menetapkan panjang minimal nama objek, sehingga batas ini adalah aturan Portal. (10) US-019 dan US-020 (IP Policy Private/Public, "mewajibkan minimal satu whitelist IP") dihapus dari 7.1.5 karena konsepnya tidak ada pada FR maupun mockup dan bertentangan dengan learning period auto-whitelist. (11) Bagian 2.1 Product Overview dan 2.2 Product Features ditulis ulang agar mencerminkan produk WAF dengan Security Profile self-service (sebelumnya berfokus ke SSL): ditambahkan alur pembuatan Node/Pool/VS/WAF Policy, Server Technologies dan Application Language Auto Detect, Learning Period & Auto-Whitelist IP, pengelolaan Whitelist berbasis ticket, serta peran User Admin dan Super Admin. (12) Skenario B (Enable SNI pada Default SSL Client Profile) dihapus dari WI Konsolidasi V2.0; skenario pada WI dan pemetaan 7.5 kini A, C, D, E. (13) FR-DWAF-0035: pemblokiran IP di luar Whitelist berlaku pada profile Blocking maupun Transparent karena dilakukan AFM, bukan ASM. | 1.16 |
| Fairuz Rafida | 07/10/2026 | Menjelaskan lebih rinci untuk pricing dan charge untuk tipe proyek yang digunakan | 1.17 |

---

## 1. Introduction

### 1.1 Purpose of the Document

Service Portal Cloudeka menyediakan solusi Security berbasis teknologi F5 yang sepenuhnya terotomasi melalui integrasi API. Portal menjadi *single point of management* bagi pelanggan untuk mengelola layanan Web Application Firewall (WAF), SSL, dan fitur keamanan lainnya tanpa memerlukan akses langsung ke perangkat fisik atau GUI F5. Load Balancing (Node dan Pool) digunakan sebagai infrastruktur pendukung internal untuk mengarahkan *traffic* ke *origin server* pelanggan sebagai bagian dari *provisioning* WAF — bukan fitur Load Balancer mandiri yang dapat dikonfigurasi terpisah oleh pelanggan (lihat Scope bagian 1.2).

Seluruh konfigurasi dan perubahan yang sebelumnya dilakukan secara manual di perangkat F5 kini dijalankan oleh Portal melalui iControl REST API. Dengan pendekatan ini, Portal bukan hanya *interface*, tetapi produk Security yang mengorkestrasi seluruh operasi F5 secara terpadu, aman, dan dapat ditelusuri (*auditable*).

### 1.2 Scope

**In-Scope:**
- Penyediaan form upload SSL certificate (*certificate*, *private key*, *chain*) melalui Service Portal.
- Validasi otomatis certificate (*format*, *key-certificate match*, *expiry check*).
- Mekanisme *learning period* 7 hari sejak Security Profile berstatus Active: IP yang mengakses domain/SSL profile selama periode tersebut otomatis tercatat ke Whitelist IP; setelah periode berakhir hanya IP pada Whitelist yang dapat mengakses dan IP lain otomatis diblok. Penambahan/penghapusan IP setelahnya dilakukan Super Admin berdasarkan ticket pelanggan, tanpa *approval workflow* (Security Profile/SSL juga otomatis terdaftar ke F5 begitu dibuat/diubah).
- Pengiriman certificate dan konfigurasi SSL ke perangkat F5 melalui API automation.
- Pembuatan atau pembaruan SSL Client Profile pada F5, termasuk pengaturan SNI.
- Attachment SSL Client Profile ke Virtual Server.
- *Synchronized deployment* ke DC1 (JAH) dan DC2 (TBS).
- Backend Node dan Load Balancing Pool creation/management via API sebagai infrastruktur pendukung untuk mengarahkan *traffic* ke *origin server* pelanggan (bagian integral dari *provisioning* WAF), BUKAN Load Balancer sebagai produk/layanan tersendiri yang dapat dikonfigurasi mandiri oleh pelanggan.
- Security Profile configuration (termasuk Bot Protection).
- Access Control management (IP Whitelist - *learning period* 7 hari, *auto-whitelist*, dan pengelolaan IP oleh Super Admin berbasis ticket).
- Monitoring & metrics dashboard untuk security profile (*Total Requests*, *Blocked*, *Block Rate*, *Avg Latency*).

**Out-of-Scope:**
- Manajemen akses perangkat F5 secara langsung oleh pelanggan.
- Pembuatan certificate (*CSR generation*) — hanya upload certificate yang sudah jadi.
- Pengaturan kustom lanjutan di luar WAF Policy default/generic yang dikelola backend Cloudeka — termasuk pembuatan custom ASM policy dari nol, penulisan iRules, dan custom attack signature di luar ruleset yang disediakan sistem.
- Manajemen infrastruktur fisik atau jaringan F5.
- Load Balancer (L7 Load Balancer) dan Global Load Balancer (GLB) sebagai produk/layanan mandiri yang dapat dibeli dan dikonfigurasi terpisah oleh pelanggan. Sesuai Katalog Layanan Load Balancer & Web Application Firewall (`ND Deka WAF.pdf`, Nota Dinas No.047/ND/22500/2019), LB dan GLB adalah SKU terpisah dengan skema harga dan unit VS sendiri (1 VS per service LB, 2 VS per service GLB), berbeda dari WAF (1 VS per service WAF) yang menjadi fokus dokumen ini. *Self-service* Load Balancer/GLB oleh pelanggan (di luar dukungan Node/Pool internal untuk WAF pada FR-DWAF-0025) memerlukan SRS terpisah bila suatu saat dikembangkan.

### 1.3 Definitions and Acronyms

**Tabel 1. Definitions and Acronyms**

| Term | Definition |
| :--- | :--- |
| **SSL** | Secure Sockets Layer, protokol keamanan yang digunakan untuk mengenkripsi komunikasi antara client dan server melalui HTTPS. |
| **SNI** | Server Name Indication, ekstensi TLS yang memungkinkan satu IP address melayani banyak domain dengan sertifikat SSL yang berbeda. |
| **TLS** | Transport Layer Security, versi modern pengganti SSL dan standar enkripsi yang digunakan dalam komunikasi HTTPS. |
| **F5** | Perangkat Application Delivery Controller (ADC) yang digunakan untuk load balancing, SSL offloading, dan WAF. |
| **SSL Client Profile** | Konfigurasi SSL pada F5 yang menentukan certificate, key, dan pengaturan TLS lainnya untuk melayani koneksi HTTPS. |
| **Cert** | File digital berisi identitas domain (CN/SAN), issuer, dan masa berlaku yang digunakan server untuk membuktikan keaslian identitasnya kepada client pada saat TLS handshake. |
| **Chain/CA Bundle** | Sekumpulan sertifikat intermediate dari Certificate Authority (CA) yang diperlukan browser untuk memverifikasi keaslian sertifikat utama (cert). Tanpa chain, browser dapat menandai sertifikat sebagai "not trusted". |
| **Key** | Kunci yang dipasangkan dengan certificate dan digunakan server untuk proses enkripsi/dekripsi. Kunci ini sangat sensitif dan tidak boleh dibagikan atau diekspos. |
| **DC1** | Data Center 1 yang digunakan untuk data center JAH (Jatiluhur). |
| **DC2** | Data Center 2 istilah yang digunakan untuk data center TBS (TB Simatupang). |
| **WAF** | Web Application Firewall, sistem keamanan yang memfilter dan memantau traffic HTTP/HTTPS untuk melindungi aplikasi web dari serangan. |
| **iControl REST API** | API berbasis HTTP/JSON yang disediakan oleh F5 BIG-IP untuk melakukan konfigurasi dan manajemen perangkat secara programatik. |
| **BIG-IP** | Platform perangkat lunak/keras dari F5 yang menyediakan layanan ADC, SSL offloading, load balancing, dan keamanan aplikasi. |
| **transid** | Transaction ID yang digunakan dalam iControl REST API untuk mengelompokkan beberapa perubahan konfigurasi menjadi satu operasi atomik. |
| **Node** | Representasi backend/origin server pelanggan pada F5 (`Local Traffic > Node List`), dirujuk oleh Pool sebagai tujuan traffic. |
| **Pool** | Kumpulan Node/backend server pada F5 yang menerima traffic dari Virtual Server, dapat berisi satu atau lebih member (`Local Traffic > Pools`). |
| **Bot Defense / Bot Protection** | Modul F5 BIG-IP (`Security > Bot Defense`) untuk mendeteksi dan memitigasi traffic bot otomatis, web scraping, dan credential stuffing, dikonfigurasi melalui Bot Defense Profile dengan preset Template (`Relaxed`/`Balanced`/`Strict`). |
| **GTM/GSLB** | Global Traffic Manager / Global Server Load Balancing, modul F5 BIG-IP DNS untuk mengarahkan traffic ke Virtual Server yang tepat lintas data center melalui GSLB Pool dan Wide IP. |
| **Wide IP** | Objek pada F5 BIG-IP DNS/GSLB yang memetakan satu FQDN publik (format: `(subdomain).glb.lintasarta.net`) ke satu atau lebih GSLB Pool. |
| **Super Admin** | Role tim internal Cloudeka (NOC/Ops) yang mengelola Whitelist IP: menambah IP (setelah learning period, termasuk via Smart Resolve) dan menghapus IP atas permintaan ticket pelanggan, tanpa approval workflow. Role ini tidak diakses langsung oleh pelanggan. |
| **User Admin** | Role pengguna dari sisi pelanggan yang membuat, mengelola (`create`/`edit`), dan menghapus (`delete`) Security Profile miliknya sendiri secara self-service melalui Service Portal, yang langsung terdaftar dan ter-deploy otomatis ke F5 begitu disimpan (tanpa approval Super Admin). Hanya dapat melihat Whitelist IP (`read-only`); penambahan/penghapusan IP diminta melalui ticket ke Super Admin. |
| **Learning Period** | Masa 7 hari sejak Security Profile berstatus Active (ketentuan tim Operation Cloudeka) di mana IP manapun dapat mengakses domain/SSL profile dan IP tersebut otomatis tercatat ke Whitelist IP. Setelah masa ini berakhir, hanya IP pada Whitelist yang dapat mengakses; IP lain otomatis diblok. |
| **Virtual Server (VS)** | Objek pada F5 (`Local Traffic > Virtual Servers`) yang menerima traffic klien pada pasangan `IP:Port` tertentu dan meneruskannya ke Pool. Setiap Security Profile menghasilkan 1 Virtual Server, yang menjadi unit tarif Deka WAF. |
| **Security Profile** | Konfigurasi perlindungan WAF milik pelanggan pada Service Portal (domain, node, port, Server Technologies, Enforcement Mode, Bot Protection, dan SSL). Satu Security Profile menghasilkan 1 Virtual Server dan 1 WAF Policy di F5. |
| **Whitelist IP** | Daftar IP yang diizinkan mengakses domain/SSL sebuah Security Profile. Terisi otomatis selama Learning Period dan setelahnya dikelola Super Admin berdasarkan ticket pelanggan. |
| **ASM** | Application Security Manager, modul F5 BIG-IP untuk WAF Policy dan event log request (namespace `/mgmt/tm/asm/`). |
| **AFM** | Advanced Firewall Manager, modul F5 BIG-IP untuk network firewall (`Security > Network Firewall`); dipakai untuk memblok IP di luar Whitelist setelah Learning Period. |
| **Address List** | Objek AFM berisi kumpulan alamat IP yang dipakai sebagai sumber pada rule Firewall Policy; pada Deka WAF isinya adalah Whitelist IP suatu Security Profile. |
| **Smart Resolve** | Fitur pada halaman Whitelist IP role Super Admin: Super Admin memasukkan PS Support ID, lalu Portal mengambil client IP dari event log ASM dan menambahkannya ke Whitelist. |

---

## 2. Description

### 2.1 Product Overview

Service Portal Cloudeka menyediakan produk Security (Web Application Firewall) berbasis F5 yang telah terintegrasi penuh melalui API. Pelanggan dapat membuat dan mengelola Security Profile secara mandiri (*self-service*) untuk melindungi aplikasi webnya, termasuk SSL certificate, Bot Protection, dan kontrol akses berbasis Whitelist IP. Service Portal bertindak sebagai:

**Aplikasi yang Menghadap Pelanggan**  
Pelanggan (User Admin) membuat Security Profile melalui wizard 3 langkah — Basic Information, Configuration (Domain, Node, Port, Server Technologies, Enforcement Mode, Bot Protection), dan SSL Certificate — serta melihat Overview, mengedit, dan menghapus profile miliknya langsung dari UI Portal, tanpa perlu pengetahuan teknis tentang perangkat F5 dan tanpa persetujuan Super Admin. Tim internal Cloudeka (Super Admin) mengelola Whitelist IP berdasarkan ticket pelanggan.

**Lapisan Otomasi**  
Service Portal melakukan otomatisasi penuh terhadap F5 melalui iControl REST API, meliputi:
- Pembuatan Backend Node, Load Balancing Pool, dan Virtual Server berdasarkan input Domain, Node, dan Port.
- Pembuatan WAF Policy sesuai Enforcement Mode (`Blocking`/`Transparent`) dan Server Technologies; Application Language ditetapkan otomatis ke `Auto Detect`.
- Upload certificate/key/chain, pembuatan SSL Client Profile, dan update certificate.
- Pemasangan toggle Bot Protection (satu profil tetap, mengikuti WI Konsolidasi V2.0 bagian E).
- Learning Period 7 hari: pencatatan otomatis IP yang mengakses, lalu hanya IP pada Whitelist yang diizinkan dan IP lain otomatis diblok; penambahan/penghapusan IP oleh Super Admin berdasarkan ticket.
- Deployment tersinkronisasi Multi-DC via Transaction API.
- Validasi TLS handshake pasca-deployment.
- Monitoring metrics & statistik traffic per security profile (*Total Requests*, *Blocked*, *Block Rate*, *Avg Latency*).

**Lapisan Kepatuhan & Audit**  
Seluruh aktivitas terekam otomatis dalam log dan dapat diaudit. Service Portal mempermudah operasional dan menghilangkan kebutuhan melakukan konfigurasi manual melalui F5.

### 2.2 Product Features

- **Manajemen Security Profile Self-Service** sehingga pelanggan (User Admin) dapat membuat, melihat, mengedit, dan menghapus Security Profile miliknya melalui wizard 3 langkah, dan konfigurasinya langsung ter-deploy ke F5 tanpa approval Super Admin.
- **Konfigurasi WAF per Profile** sehingga pelanggan dapat memilih Enforcement Mode (`Blocking`/`Transparent`), Server Technologies (opsional), dan mengaktifkan Bot Protection; Application Language ditetapkan otomatis ke `Auto Detect`.
- **Learning Period & Auto-Whitelist IP** sehingga selama 7 hari sejak profile Active semua IP dapat mengakses dan IP yang mengakses tercatat otomatis ke Whitelist, lalu setelah 7 hari hanya IP pada Whitelist yang dapat mengakses dan IP lain otomatis diblok.
- **Pengelolaan Whitelist IP Berbasis Ticket** sehingga Super Admin menambah atau menghapus IP atas permintaan pelanggan (termasuk Smart Resolve via PS Support ID) dan perubahan langsung diterapkan ke F5 tanpa approval, sementara User Admin hanya dapat melihat daftar Whitelist.
- **Import Otomatis SSL Certificate** sehingga pelanggan akan mengunggah cert/key/chain melalui Service Portal dan Service Portal langsung kirim ke F5 melalui API.
- **Manajemen SSL Client Profile** sehingga Service Portal membuat atau memperbarui SSL Client Profile sesuai domain pelanggan.
- **Manajemen Backend Node dan Pool** sehingga Service Portal otomatis membuat node dan load balancing pool untuk routing traffic.
- **Deployment Tersinkronisasi Multi-DC** — Service Portal melakukan deployment konfigurasi ke DC1 dan DC2 melalui API.
- **Konfigurasi SNI Otomatis** — Service Portal otomatis mengaktifkan Default SSL Profile for SNI.
- **Deployment Otomatis ke Virtual Server** sehingga Service Portal meng-attach profile baru ke VS secara langsung melalui API.
- **Validasi Otomatis Pasca-Deployment** sehingga Service Portal menjalankan TLS handshake check dan menampilkan hasilnya.
- **Audit Logging Menyeluruh** sehingga semua perubahan dicatat untuk compliance & troubleshooting.
- **Monitoring & Metrics Dashboard** sehingga pelanggan dapat memantau *Total Requests*, *Blocked*, *Block Rate*, dan *Avg Latency* untuk setiap security profile secara real-time.

### 2.3 User Documentation

Daftar komponen dokumentasi pengguna (seperti manual pengguna, bantuan daring, dan tutorial) yang akan disertakan bersama perangkat lunak yang akan dikirimkan. Dokumentasi untuk produk akan tersedia di `https://docs.cloudeka.ai`.

---

## 3. Interface Requirement

### 3.1 System Architecture

Sistem terdiri dari tiga lapisan utama yang saling berkomunikasi secara terstruktur:

```text
+---------------------------------------------------------------------------------------------------+
| LAYER 1 — Customer / Browser                                                                      |
|   [ Cloudeka Service Portal UI (Web Browser) ]                                                    |
+---------------------------------------------------------------------------------------------------+
                                                 |
                                               HTTPS
                                                 v
+---------------------------------------------------------------------------------------------------+
| LAYER 2 — Portal Backend                                                                          |
|   +-------------------------------------------------------------------------------------------+   |
|   | Automation Engine (Go / Python)                                                           |   |
|   |  [ Validation ]   [ Learning Period ]   [ Audit Logging ]   [ API Orchestration ]         |   |
|   +-------------------------------------------------------------------------------------------+   |
|   [ iControl REST API (HTTPS / JSON) ]                                                            |
+---------------------------------------------------------------------------------------------------+
                             |                                           |
                             v                                           v
+---------------------------------------------------------------------------------------------------+
| LAYER 3 — F5 BIG-IP Devices                                                                       |
|   +---------------------------------------+             +---------------------------------------+ |
|   | F5 BIG-IP — DC1                       |  <--SYNC--> | F5 BIG-IP — DC2                       | |
|   | (JAH - Jatiluhur)                     |             | (TBS - TB Simatupang)                 | |
|   +---------------------------------------+             +---------------------------------------+ |
+---------------------------------------------------------------------------------------------------+

Legenda:
  --->  Unidirectional flow
  <-->  Bidirectional sync
  [Dark box]  F5 BIG-IP Device
  [Light box] Software Component
```

*Gambar 1. System Architecture*

Berikut penjelasan untuk Gambar 1. System Architecture:

- **Layer 1 - Pelanggan / Browser**  
  Pelanggan berinteraksi melalui Cloudeka Service Portal UI menggunakan web browser. Semua permintaan dikirim melalui koneksi HTTPS terenkripsi ke Portal Backend.
- **Layer 2 - Portal Backend**  
  Automation Engine (Go / Python) menerima permintaan dari Frontend dan menjalankan orkestrasi penuh terhadap F5 BIG-IP melalui iControl REST API (HTTPS / JSON). Layer ini mencakup:
  - **Validasi Input** — validasi format certificate, key-cert match, expiry check.
  - **Learning Period & Auto-Whitelist** — pencatatan otomatis IP yang mengakses selama 7 hari sejak profile Active, lalu pemblokiran otomatis IP lain setelah periode berakhir.
  - **Audit Logging** — pencatatan setiap aksi dengan timestamp dan identitas pengguna.
  - **Orkestrasi API** — eksekusi sequence iControl REST API calls ke F5 DC1 & DC2.
- **Layer 3 - Perangkat F5 BIG-IP**  
  Dua perangkat F5 BIG-IP yang beroperasi di dua data center berbeda:
  - F5 BIG-IP DC1 — JAH (Jatiluhur)
  - F5 BIG-IP DC2 — TBS (TB Simatupang)
  
  Kedua data center di-deploy secara sinkron melalui iControl REST Transaction untuk memastikan konsistensi konfigurasi.

### 3.2 Software Interface and Stack

**Platform:**
- Service Portal Cloudeka (Frontend + Backend API)

**Backend:**
- Automation Engine (Go/Python)
- F5 iControl REST API Integration (v16.1)

**Tabel 2. Technology Stack**

| Component / Layer | Technology |
| :--- | :--- |
| Frontend (Portal UI) | Cloudeka Service Portal (Web Browser) |
| Backend Automation | Go / Python (Automation Engine) |
| Protokol API F5 | iControl REST API v16.1 |
| Format Data | JSON over HTTPS |
| Autentikasi | Berbasis token (iControl REST tokens) |
| Perangkat F5 | BIG-IP DC1 (JAH, Jatiluhur) & DC2 (TBS, TB Simatupang) |
| Audit | Internal Portal Logging |
| Dokumentasi | `https://docs.cloudeka.ai` |

### 3.3 iControl REST API Fundamentals

Tabel berikut merangkum parameter teknis fundamental dari F5 iControl REST API yang digunakan dalam integrasi ini:

**Tabel 3. Parameter**

| Parameter | Value |
| :--- | :--- |
| Base URL | `https://<f5-ip>/mgmt/tm/` |
| Protokol | HTTPS + JSON |
| Autentikasi | Berbasis token (iControl REST tokens — header `X-F5-Auth-Token`) |
| Metode HTTP | `GET`, `POST`, `PATCH`, `DELETE` |
| Transaksi | Operasi multi-step atomik all-or-nothing (commit atau full rollback) |

**Tabel 4. iControl REST API Namespaces**

| API Namespace | Module | Purpose |
| :--- | :--- | :--- |
| `/mgmt/tm/ltm/` | Load Traffic Manager | Manajemen SSL Profile, Virtual Server, Node, Pool, dan Stats/Monitoring (`Total Requests`, `Blocked`, `Block Rate`, `Avg Latency`) |
| `/mgmt/tm/sys/crypto/` | Cryptographic Store | Import dan kelola Key & Certificate pada F5 |
| `/mgmt/tm/sys/file/` | File Storage | Upload file SSL certificate (transfer biner) |
| `/mgmt/tm/asm/` | App Security Manager | WAF Policy, Learning Mode, Enforcement Mode, dan Event Log Requests (sumber pencatatan IP selama learning period) |
| `/mgmt/tm/security/firewall/` | Advanced Firewall Manager (AFM) Network Firewall | Whitelist IP sebagai Address List per profile dan Firewall Policy (allow Whitelist + default deny) yang di-enforce ke Virtual Server setelah learning period 7 hari (WI V2.0 D.4-D.6) |
| `/mgmt/shared/` | Shared Services | Token autentikasi, transfer file (multipart upload) |
| `/mgmt/tm/transaction` | Transaction Management | Operasi multi-step atomik untuk deployment sinkron DC1 & DC2 |
| `/mgmt/tm/security/bot-defense/profile`; `/mgmt/tm/security/log/profile` (sub-objek `botDefense`) | Security Bot Defense (modul terpisah dari ASM) | Bot Defense Profile: properti `template=relaxed\|balanced\|strict` (WI V2.0 E.2), `enforcement-mode`, `deviceid-mode`, `whitelist`. Log Profile: publisher `local-bot-defense` untuk logging Bot Defense (WI V2.0 E.1). Attach ke VS via field Security Policy pada `/mgmt/tm/ltm/virtual/{vsName}` (WI V2.0 E.3). Diverifikasi terhadap `techdocs.f5.com`/`clouddocs.f5.com` (BIG-IP 16.x) - path GUI: `Security > Bot Defense > Bot Defense Profiles`. |

---

## 4. Pricing & Charge

Pada bagian ini menjelaskan untuk pricing dan charge yang digunakan pada Deka WAF. Untuk produk Deka WAF dikenakan charge berdasarkan satu Virtual Server digunakan untuk 1 pelanggan. Sesuai dengan data di Nota Dinas No.047/ND/22500/2019 di halaman 10 untuk pricing Deka WAF yang dikenakan adalah **Rp. 2.700.000,00 per bulan** untuk satu Virtual Server digunakan tipe **Postpaid Fixed**. Untuk tipe pelanggan **Postpaid Usage Based** dan **Prepaid** dikenakan **Rp. 3.700,00 per jam**.

---

## 5. Functional Requirement

### 5.1 Functional Requirement Table

| ID | Feature | Description | Acceptance Criteria |
| :--- | :--- | :--- | :--- |
| **FR-DWAF-0001** | Melihat Daftar Security Profile | Pengguna dapat melihat daftar semua security profile yang telah dibuat dengan informasi: Profile Name, Domain, Mode (`Blocking`/`Transparent`), Status (`Active`/`Inactive`/`Error`), Created Date, Last Updated, dan Action buttons. | - Halaman menampilkan tabel dengan semua field yang required<br>- Data dapat di-sort berdasarkan Profile Name, Domain, Status, atau Created Date<br>- Terdapat filter untuk Mode dan Status<br>- Tersedia pagination dengan default 10 items per halaman<br>- Tersedia search box untuk mencari profile berdasarkan nama atau domain<br>- Tampil empty state ketika tidak ada profile |
| **FR-DWAF-0002** | Membuat Security Profile Baru | Pengguna dapat membuat security profile baru melalui form multi-step yang terdiri dari 3 bagian: Basic Information, Configuration, dan SSL Certificate. | - **Step 1 (Basic Information):** Dapat input Profile Name (min 3 karakter, sesuai mockup dan FR-DWAF-0018) dan Description<br>- **Step 2 (Configuration):** Dapat input Domain, Node, Port (dropdown `80 - HTTP`/`443 - HTTPS`, wajib dipilih), Server Technologies, Enforcement Mode (`Blocking`/`Transparent` - wajib dipilih, tidak ada default), dan centang Protection Features (Bot Protection)<br>- **Step 3 (SSL Certificate):** Dapat upload SSL Chain, SSL Cert, dan SSL Key (max 5MB per file)<br>- Summary panel di sisi kanan menampilkan data yang sudah diisi<br>- Validasi untuk setiap field yang required<br>- Tombol Cancel kembali ke halaman list<br>- Tombol Create untuk menyimpan profile baru<br>- Notifikasi sukses setelah profile berhasil dibuat<br>- Setiap Security Profile yang berhasil dibuat menghasilkan tepat 1 (satu) Virtual Server pada F5, yang otomatis diregistrasikan begitu profile disimpan (tanpa proses approval terpisah), dipicu oleh domain yang didaftarkan pada profile tersebut. Virtual Server inilah yang menjadi unit dasar perhitungan biaya — BUKAN domain — sesuai kesepakatan Nota Dinas No.047/ND/22500/2019 dan skema pricing per Virtual Server pada Bab 4 (Pricing & Charge). Platform F5 bersifat multitenant (satu perangkat fisik melayani banyak pelanggan via partisi/VS terisolasi), sehingga tarif tetap dihitung per VS, bukan per pelanggan atau per domain. |
| **FR-DWAF-0003** | Melihat Detail Security Profile | Pengguna dapat melihat detail lengkap dari security profile termasuk overview, konfigurasi, dan monitoring. | - Menampilkan profile header dengan nama, domain, status, node, server technologies, dan last updated<br>- Menampilkan metrics: Total Requests, Blocked, Block Rate, dan Avg Latency<br>- Menampilkan card Profile Configuration (3 kolom) berisi: Domain, Description, Node, Port, Mode, Bot Protection, SSL Certificate, SSL Chain, SSL Key<br>- Link ke halaman detail profile di header breadcrumb<br>- Setting button untuk akses ke edit profile |
| **FR-DWAF-0004** | Mengedit Security Profile - Pengaturan Umum | Pengguna dapat mengedit general settings dari security profile seperti Profile Name, Domain, Description, dan WAF Mode. | - Form menampilkan field: Profile Name, Domain, Description<br>- Dapat mengubah WAF Mode antara Blocking dan Transparent dengan deskripsi masing-masing mode<br>- Alert menunjukkan dampak perubahan WAF Mode<br>- Tombol Cancel untuk batalkan perubahan<br>- Tombol Save Changes untuk simpan perubahan<br>- Notifikasi sukses setelah perubahan disimpan |
| **FR-DWAF-0005** | Mengedit Security Profile - Sertifikat SSL/TLS | Pengguna dapat upload atau update SSL/TLS certificates untuk HTTPS inspection pada security profile. | - Form menampilkan 3 file upload fields: SSL Certificate, SSL Chain, SSL Private Key<br>- Format accepted: `.pem`, `.crt`, `.cer` untuk certificate; `.pem`, `.key` untuk private key<br>- Max file size 5MB per file<br>- Drag & drop support untuk file upload<br>- Menampilkan file name setelah upload<br>- Dapat replace existing certificate<br>- Tombol Cancel untuk batalkan<br>- Tombol Save Changes untuk simpan |
| **FR-DWAF-0008** | Konfigurasi Bot Protection | **RINGKASAN:** Bot Protection adalah checkbox pada Security Profile pelanggan yang mengontrol proteksi bot secara langsung pada Virtual Server milik pelanggan tersebut (VS yang sama dengan yang dibuat untuk WAF Policy pada FR-DWAF-0002) — BUKAN VS terpisah yang dipakai bersama pelanggan lain.<br><br>**Yang dilihat/dilakukan PELANGGAN:** mencentang checkbox "Bot Protection" saat Create/Edit Security Profile, lalu melihat status "Enabled/Disabled" di halaman Overview.<br><br>**Yang terjadi di F5 (WI Konsolidasi V2.0 bagian E, direvisi):** saat checkbox dicentang, Portal meng-attach SATU Bot Defense Profile standar beserta Bot Defense Logging Profile ke Virtual Server milik pelanggan tsb (VS yang sama dengan WAF Policy pelanggan) — BUKAN ke VS bersama. Saat checkbox dilepas, Portal men-detach kedua profile tsb dari VS pelanggan.<br><br>**KENAPA desainnya diubah:** desain sebelumnya (1 VS dipakai bersama SEMUA pelanggan) tidak dapat diimplementasikan pada F5 — setiap Virtual Server terikat ke 1 pasangan `IP:Port` dan 1 Node/Pool milik satu domain tertentu (WI bagian A.6), sehingga tidak mungkin 1 VS menerima dan memproteksi traffic untuk seluruh domain pelanggan yang berbeda-beda sekaligus. Bot Defense Profile dan Logging Profile karena itu di-attach di level VS masing-masing pelanggan, sama seperti komponen WAF Policy lainnya, bukan pada 1 VS terpisah yang dipakai bersama (Bot Defense Profile-nya sendiri tetap SATU objek standar `template=Balanced`, `mode=Blocking` — yang berubah hanya target attachment-nya, kini per VS pelanggan, bukan jumlah profile-nya). | **[UI Portal - per pelanggan, di mockup]**<br>- Checkbox `"Bot Protection - Block malicious bots"` tersedia pada Step 2 (Configuration) form Create Security Profile<br>- Status Bot Protection (`Enabled`/`Disabled`) ditampilkan pada halaman Overview profile<br>- Mencentang checkbox memicu Portal untuk meng-attach Bot Defense Profile & Logging Profile ke Virtual Server milik pelanggan via API; melepas centang memicu detach - status di Portal selalu sinkron dengan konfigurasi aktual di F5<br><br>**[Backend F5 - attach/detach per pelanggan, dipicu checkbox]**<br>- Bot Defense Logging Profile dibuat SEKALI sebagai resource bersama (WI V2.0 E.1) via `POST /mgmt/tm/security/log/profile` (sub-objek `botDefense`, publisher `local-bot-defense`)<br>- SATU Bot Defense Profile standar dibuat SEKALI sebagai resource bersama oleh tim Cloudeka (WI V2.0 E.2) via `/mgmt/tm/security/bot-defense/profile`, properti `template=balanced`, `enforcementMode=blocking`<br>- Saat checkbox Bot Protection dicentang/dilepas, Portal meng-attach/men-detach Bot Defense Profile & Log Profile tsb ke/dari Virtual Server milik pelanggan yang bersangkutan (WI V2.0 E.3, direvisi) via `PATCH /mgmt/tm/ltm/virtual/{vsName}`, dengan `vsName` = VS pelanggan tsb — VS YANG SAMA dengan VS WAF Policy pelanggan (WI bagian A.6, FR-DWAF-0002), BUKAN VS terpisah<br>- **KLARIFIKASI BILLING:** karena Bot Defense di-attach ke Virtual Server pelanggan yang sudah ada dan sudah dihitung dalam skema pricing per VS (Bab 4 Pricing & Charge, mengacu Nota Dinas No.047/ND/22500/2019 — lihat `ND Deka WAF.pdf`), mengaktifkan/menonaktifkan Bot Protection TIDAK menambah maupun mengurangi unit VS billing - hanya menambah/menghapus atribut Security Policy (Bot Defense Profile & Log Profile) pada VS pelanggan yang sama seperti pada FR-DWAF-0002.<br>- Status checkbox pelanggan kini selalu konsisten dengan enforcement nyata di F5 untuk pelanggan tsb, karena setiap perubahan checkbox memicu API call attach/detach langsung ke VS milik pelanggan tsb — menggantikan desain VS bersama sebelumnya yang secara teknis tidak dapat diimplementasikan pada F5. |
| **FR-DWAF-0009** | Mengelola Whitelist IP | Pengguna dapat melihat daftar IP yang di-whitelist, yaitu IP yang diizinkan mengakses domain/SSL profile (allowlist; inspeksi WAF tetap berlaku). Daftar terisi otomatis dari IP yang mengakses selama learning period 7 hari (FR-DWAF-0035). Penambahan dan penghapusan IP hanya dapat dilakukan Super Admin atas permintaan pelanggan melalui ticket; User Admin termasuk yang membuat/mengelola Security Profile (Deka WAF) terkait hanya memiliki akses read-only dan diarahkan untuk "Submit a ticket" bila ada IP yang perlu dihapus atau ditambahkan, sesuai mockup 7.1.8 (User Admin View) dan 7.1.9 (Super Admin View). | - Menampilkan list whitelist IP dengan kolom: IP Address, Added Date, Actions<br>- Daftar terisi otomatis dari IP yang mengakses domain/SSL profile selama learning period 7 hari (FR-DWAF-0035), tanpa input dari User Admin maupun Super Admin dan tanpa approval<br>- Pada role Super Admin, tersedia input `"Enter IP address"` + tombol Add untuk langsung meng-authorize IP (digunakan setelah learning period, atas permintaan ticket pelanggan); keterangan `"Add trusted IP support ID to immediately authorize IP address"` ditampilkan di atas form, sesuai mockup<br>- **Smart Resolve:** Super Admin memasukkan PS Support ID; Portal mencari request tersebut pada event log ASM, mengambil client IP-nya, lalu menambahkannya ke Whitelist (hanya untuk request yang diproses ASM; lihat API List langkah 8)<br>- Menampilkan deskripsi untuk masing-masing IP (optional)<br>- Search box untuk mencari IP<br>- Pagination dengan 10 items per halaman default<br>- Delete icon untuk remove IP dari whitelist (role Super Admin), dilakukan atas permintaan ticket pelanggan<br>- Notifikasi ketika IP di-remove dari whitelist<br>- User Admin dengan whitelist kosong melihat empty state `"Your IP list is empty"` beserta pesan `"Submit a ticket to request IP allowlisting."`, sesuai mockup 7.1.7<br>- Penambahan dan penghapusan IP oleh Super Admin langsung diterapkan otomatis ke F5 dengan memperbarui Address List profile melalui iControl REST API (API List langkah 6-7), tanpa approval workflow<br>- Setiap aksi Add/Delete pada Whitelist tercatat dalam audit log dengan timestamp dan identitas Super Admin<br>- User Admin yang menilai suatu IP tidak seharusnya memiliki akses mengajukan ticket penghapusan; UI User Admin tidak menyediakan aksi Add maupun Delete<br>- Hanya user dengan role Super Admin yang memiliki akses UI untuk menambah (Add) dan menghapus (Delete) IP pada Whitelist; User Admin hanya memiliki akses read-only (view-only), meskipun User Admin adalah pembuat/pengelola Security Profile (Deka WAF) yang bersangkutan |
| **FR-DWAF-0012** | Navigasi Access Control | Pengguna dapat menavigasi ke halaman Whitelist IP melalui sidebar menu Security Profile. | - Sidebar menampilkan menu items: Overview, Whitelist IP<br>- Active menu item ditampilkan dengan highlight/bold<br>- Breadcrumb navigation menampilkan path: `Deka WAF > Whitelist`, sesuai mockup<br>- Click pada menu item navigasi ke halaman yang sesuai<br>- Profile dropdown menampilkan list profile yang dapat di-switch |
| **FR-DWAF-0013** | Pemilihan & Perpindahan Profile | Pengguna dapat memilih dan switch antara berbagai security profiles. | - Dropdown di sidebar menampilkan list semua security profiles<br>- Profile yang sedang active ditampilkan dengan selected indicator<br>- Click pada profile lain untuk switch<br>- Semua data pada page update sesuai profile yang dipilih<br>- Last visited section untuk profile disimpan (optional) |
| **FR-DWAF-0014** | Tampilan Status Bot Protection (Overview Profile) | Sistem menampilkan status Bot Protection pada card "Profile Configuration" di halaman Overview security profile (kolom Port/Mode/Bot Protection). Ini menggantikan konsep halaman landing "Protection Features" dengan 3 kartu terpisah (Managed Rules/Rate Limiting/Bot Protection) yang sudah dihapus dari mockup terkini. | - Card "Profile Configuration" pada halaman Overview menampilkan status Bot Protection: `Enabled`/`Disabled` beserta label singkat (`"Block malicious bots enabled"`)<br>- Status diperbarui otomatis mengikuti perubahan pada form Create/Edit Security Profile |
| **FR-DWAF-0016** | Metrik Dashboard | Menampilkan key metrics untuk security profile dalam bentuk card widgets. | - Menampilkan 4 metrics: Total Requests, Blocked, Block Rate, Avg Latency<br>- Setiap metric menampilkan: nilai, label, trend indicator (up/down dengan percentage)<br>- Metrics diupdate berdasarkan time range yang dipilih (e.g., Last 24h)<br>- Icon yang sesuai untuk setiap metric<br>- Color coding untuk trend (green untuk positive, red untuk negative) |
| **FR-DWAF-0017** | Penanganan HTTP Request | Sistem dapat menghandle incoming HTTP requests dan menerapkan security rules sesuai konfigurasi profile. | - Bot detection dijalankan untuk identify malicious bots (mengikuti profil Bot Defense tetap sesuai WI V2.0 bagian E)<br>- Whitelist rules di-apply: selama learning period semua IP diizinkan, setelah periode berakhir IP di luar Whitelist diblok (FR-DWAF-0035)<br>- Request di-block atau allow sesuai hasil evaluasi dan enforcement mode<br>- Timestamp dan details request dicatat untuk audit trail |
| **FR-DWAF-0018** | Validasi Form | Sistem melakukan validasi pada semua input form untuk memastikan data yang valid. | - **Profile Name:** min 3 characters, tidak boleh kosong (batas minimal adalah aturan validasi Portal sesuai mockup; F5 tidak menetapkan panjang minimal nama objek)<br>- **Domain:** format valid (`domain.com`), required<br>- **Node/IP:** format valid, required<br>- **Port:** wajib dipilih dari dropdown (`80 - HTTP`/`443 - HTTPS`)<br>- **Server Technologies:** optional, tidak wajib dipilih (multi-select, sesuai mockup, tidak ada tanda `*`)<br>- **File upload:** validate file size (max 5MB) dan file format<br>- **IP Address:** validate format IPv4 atau IPv6<br>- Error message ditampilkan di bawah field yang invalid<br>- Submit button disabled sampai semua required fields valid<br>- Visual indicator untuk field yang error (red border) |
| **FR-DWAF-0019** | Hak Akses & Otorisasi Pengguna | Sistem dapat mengontrol akses dan permissions berdasarkan role pengguna. | - Super Admin (Cloudeka) dapat melihat seluruh Security Profile lintas pelanggan untuk keperluan monitoring dan oversight<br>- Super Admin dapat menambah dan menghapus Whitelist IP atas permintaan ticket pelanggan (tanpa approval workflow) dan mengonfigurasi protection features bila diperlukan<br>- User Admin hanya dapat melihat (read-only) Whitelist IP; penambahan/penghapusan IP diminta melalui ticket (FR-DWAF-0009)<br>- User Admin dapat membuat (create) dan mengedit (edit) Security Profile miliknya sendiri secara self-service, yang langsung terdaftar dan ter-deploy otomatis ke F5 (tanpa approval Super Admin)<br>- User Admin hanya dapat melihat dan mengelola Security Profile miliknya sendiri (tidak dapat mengakses profile milik pelanggan lain)<br>- User Admin dapat menghapus (delete) Security Profile miliknya sendiri (self-service, mengikuti alur konfirmasi 2 langkah pada mockup)<br>- Non-authorized user mendapat error message saat mencoba mengakses restricted action |
| **FR-DWAF-0020** | Deployment Otomatis SSL Certificate & Key ke F5 | Sistem (Automation Engine) secara otomatis menginstall certificate, private key, dan chain yang sudah divalidasi ke F5 crypto store melalui iControl REST API tanpa akses manual ke perangkat F5. | - Certificate dan key diinstall otomatis ke F5 (`/mgmt/tm/sys/crypto/`) setelah validasi form SSL Certificate berhasil<br>- SSL Chain bersifat optional; jika tidak diupload, sistem menggunakan default chain `"none"` sesuai perilaku default F5<br>- Proses instalasi dijalankan sepenuhnya melalui API tanpa akses manual ke perangkat F5<br>- Kegagalan instalasi ditampilkan sebagai notifikasi error ke pengguna beserta penyebabnya<br>- Setiap proses instalasi tercatat dalam audit log |
| **FR-DWAF-0021** | Manajemen Otomatis SSL Client Profile | Sistem secara otomatis membuat atau memperbarui SSL Client Profile pada F5 yang mereferensikan certificate dan key yang telah diinstall, termasuk mengaktifkan Default SSL Profile for SNI. | - SSL Client Profile dibuat otomatis setelah certificate dan key berhasil diinstall ke F5<br>- Update certificate pada profile existing memperbarui SSL Client Profile yang sama tanpa membuat profile baru<br>- Default SSL Profile for SNI diaktifkan otomatis pada setiap SSL Client Profile yang dibuat<br>- Satu SSL Client Profile berlaku untuk satu domain pelanggan sesuai batasan sistem<br>- Perubahan pada SSL Client Profile tercatat dalam audit log |
| **FR-DWAF-0022** | Attachment Otomatis SSL Profile ke Virtual Server | Sistem secara otomatis meng-attach SSL Client Profile ke Virtual Server yang sesuai dengan domain pada profile pengguna melalui API, menggantikan profile lama jika sudah ada. | - SSL Client Profile ter-attach otomatis ke Virtual Server yang sesuai dengan Domain dan Node pada profile<br>- Jika Virtual Server sudah memiliki SSL Client Profile sebelumnya, sistem menggantikan referensi profile lama dengan yang baru<br>- Proses attachment tidak menyebabkan downtime pada traffic yang sedang berjalan<br>- Kegagalan attachment ditampilkan sebagai notifikasi error dan tidak mengubah status Virtual Server yang sedang live |
| **FR-DWAF-0023** | Validasi TLS Otomatis Pasca-Deployment | Sistem menjalankan validasi TLS handshake terhadap domain pelanggan setelah proses deployment selesai untuk memastikan certificate aktif dan valid, kemudian menampilkan hasilnya kepada pengguna. | - TLS handshake check dijalankan otomatis setelah SSL Client Profile berhasil di-attach ke Virtual Server<br>- Hasil validasi (`Success`/`Failed`) ditampilkan kepada pengguna pada halaman profile<br>- Detail certificate (issuer, expiry date) ditampilkan jika validasi berhasil<br>- Notifikasi error ditampilkan jika handshake gagal beserta kemungkinan penyebabnya<br>- Validasi ini dijalankan oleh Automation Engine, bukan melalui iControl REST API F5 |
| **FR-DWAF-0024** | Deployment Tersinkronisasi Multi-DC | Sistem melakukan deployment konfigurasi SSL dan security profile secara sinkron ke F5 BIG-IP DC1 (JAH) dan DC2 (TBS) menggunakan iControl REST Transaction API untuk menjaga konsistensi konfigurasi antar data center. | - Setiap perubahan konfigurasi SSL/security profile dikirim ke DC1 dan DC2 dalam satu transaksi (all-or-nothing)<br>- Jika salah satu data center gagal memproses transaksi, seluruh perubahan di-rollback pada kedua DC<br>- Status sinkronisasi kedua DC ditampilkan kepada pengguna setelah proses deployment selesai<br>- Sistem mendukung multiple concurrent deployment tanpa menyebabkan race condition antar transaksi |
| **FR-DWAF-0025** | Provisioning Otomatis Backend Node & Load Balancing Pool | Sistem secara otomatis membuat backend node dan load balancing pool pada F5 berdasarkan Node/IP yang diinput pengguna saat membuat security profile, termasuk dukungan multiple backend member dalam satu pool. | - Node dibuat otomatis di F5 berdasarkan input Node pada form Create Profile<br>- Pool dibuat otomatis dan node yang diinput ditambahkan sebagai pool member<br>- Sistem mendukung lebih dari satu backend node/member dalam satu pool untuk domain yang sama<br>- Keanggotaan pool member dapat diverifikasi melalui subcollection `/members` pada API |
| **FR-DWAF-0027** | Routing Domain Multi-Site via Global Load Balancing/GLB (Diusulkan — Menunggu Keputusan Scope) | Untuk pelanggan yang memiliki aplikasi di lebih dari satu data center, sistem menyediakan opsi mengarahkan satu public domain ke site yang tepat melalui F5 BIG-IP DNS (GTM), sebagai pelengkap arsitektur single-site (Domain + Node) yang berlaku saat ini. Fitur ini bersifat usulan dan menunggu keputusan resmi apakah GLB termasuk dalam scope Service Portal. | - Sistem dapat membuat/membaca GTM Wide IP yang memetakan satu domain publik ke beberapa GTM Pool (satu per site/data center)<br>- Setiap GTM Pool berisi GTM Pool Member yang mengarah ke Virtual Server milik site tersebut, terpisah dari LTM pool/node lokal site itu<br>- Load balancing method antar site (misalnya round-robin) dapat dikonfigurasi oleh pengguna<br>- Fitur ini tidak diaktifkan sampai ada keputusan scope resmi bahwa GLB termasuk dalam cakupan Service Portal; sampai saat itu, Domain + Node pada wizard tetap diasumsikan single-site<br>- Wizard dan dokumentasi secara jelas membedakan GTM Pool (multi-site) dari LTM Pool (backend lokal) agar tidak membingungkan pengguna maupun pengembang |
| **FR-DWAF-0028** | Mode Automatic Learning WAF | Pengguna dapat mengaktifkan automatic learning suggestion pada security profile sehingga F5 memberikan rekomendasi penyesuaian rule berdasarkan traffic yang diamati, sebagai pengaturan terpisah dari WAF Mode (`Blocking`/`Transparent`). | - Tersedia toggle `"Automatic Learning"` yang independen dari WAF Mode (`Blocking`/`Transparent`) dan dapat dikombinasikan dengan keduanya (misal Blocking + Automatic Learning adalah kombinasi valid)<br>- Ketika diaktifkan (`"automatic"`), sistem mengirim update ke policy-builder profile pada F5 sehingga rekomendasi rule dibuat otomatis dari observed traffic<br>- Ketika dinonaktifkan, learning mode diset ke `"manual"` sehingga tidak ada rekomendasi otomatis<br>- Status learning mode (`Automatic`/`Manual`) ditampilkan pada halaman Protection/General Settings profile<br>- Perubahan learning mode tercatat dalam audit log |
| **FR-DWAF-0029** | Granularitas Blocking Per-Violation (Advanced Mode) | Untuk kebutuhan lanjutan, sistem menyediakan mode konfigurasi advanced/expert yang memungkinkan Super Admin melihat dan menyesuaikan perilaku alarm/block/learn per tipe violation individual (misal SQL Injection, missing parameter, cookie tampering), alih-alih hanya satu mode global. Fitur ini bersifat opsional dan disembunyikan secara default agar wizard utama tetap sederhana. | - Secara default, blocking behavior per violation type mengikuti template bawaan F5 (inherited default) dan tidak ditampilkan di wizard utama<br>- Opsi `"Advanced Mode"` tersedia bagi Super Admin untuk membuka daftar violation type beserta status alarm/block/learn masing-masing<br>- Perubahan pada Advanced Mode hanya memengaruhi violation type yang dipilih, tanpa mengubah violation type lain<br>- Advanced Mode tidak wajib digunakan; profile tetap berfungsi normal hanya dengan WAF Mode dan Automatic Learning saja<br>- Perubahan pada level violation tercatat dalam audit log |
| **FR-DWAF-0033** | Service Port pada Create Security Profile | Sudah diimplementasikan pada `mockup-waf-new.pdf`: field `"Port *"` ditambahkan pada Step 2 (Configuration), tepat di bawah field Node, berupa dropdown wajib diisi dengan 2 pilihan tetap: `"80 - HTTP"` dan `"443 - HTTPS"`. Ini menutup sebagian gap WI Konsolidasi V2.0 (bagian 3 Prasyarat dan A.2) yang mensyaratkan Service Port sesuai request pelanggan. | - Field `"Port *"` tersedia pada Step 2 (Configuration), tepat di bawah field Node, wajib diisi<br>- Berupa dropdown dengan pilihan tetap: `"80 - HTTP"`, `"443 - HTTPS"` (bukan input bebas)<br>- Panel Summary menampilkan `"Port — Select a port"` sebagai item wajib (indikator X merah/centang hijau)<br>- Card Profile Configuration pada Overview menampilkan nilai Port yang dipilih (mis. `"80 - HTTP"`)<br>- Nilai dipetakan ke Service Port pada Pool Member & Health Monitor (WI V2.0 A.2) dan Virtual Server (WI V2.0 A.6)<br>- **CATATAN GAP TERSISA:** dropdown hanya mencakup 2 port umum ($80/443$). WI V2.0 menyebut Service Port "sesuai port request pelanggan" yang bisa bervariasi (mis. 8080, 8443, port custom lain) — pelanggan dengan kebutuhan port di luar $80/443$ belum terakomodasi. Perlu klarifikasi produk: apakah dropdown akan ditambah opsi "Custom", atau memang scope awal dibatasi ke 2 port ini saja. |
| **FR-DWAF-0034** | Server Technologies pada Create Security Profile | Sudah diimplementasikan pada `mockup-deka-waf-new2.pdf`: field `"Server Technologies"` menggantikan field `"Application Language"` (single-select) pada Step 2 (Configuration), berupa multi-select/tag input yang dikelompokkan per kategori (Web Server, Language & Framework, Database), bersifat optional. Field yang sama juga sudah tampil pada header halaman Overview dan Edit (menggantikan "App Language"), meski card "Profile Configuration" belum menambahkan baris terpisah untuk field ini. Ini menutup gap WI Konsolidasi V2.0 (bagian A.8) yang meminta informasi teknologi server (web server, database, framework) untuk mengisi "Server Technologies" pada WAF Policy. Detail desain lihat `"Spesifikasi Perubahan Form Create Security Profile - Deka WAF.docx"`. Karena field Application Language tidak lagi dimasukkan pelanggan, Automation Engine menetapkan `Application Language = Auto Detect` pada WAF Policy secara otomatis (WI V2.0 A.8). Daftar teknologi yang dapat dipilih dan penjelasan tiap teknologi tercantum pada Appendix 7.6 (Server Technologies Description). | - Field `"Server Technologies"` (multi-select, kategori Web Server / Language & Framework / Database) tersedia pada Step 2 (Configuration), menggantikan field `"Application Language"`, tetap optional<br>- Nilai dipetakan ke Server Technologies pada WAF Policy (WI V2.0 A.8) untuk membantu akurasi rule/learning<br>- Overview & Edit header juga menampilkan `"Server Technologies"` (menggantikan `"App Language"`), tampil `"-"` bila belum ada nilai dipilih<br>- Application Language tidak ditampilkan dan tidak dapat diubah di Portal; Automation Engine selalu menetapkannya ke `Auto Detect` saat membuat WAF Policy (WI V2.0 A.8), sehingga encoding aplikasi dikenali otomatis oleh F5<br>- Opsi Server Technologies mengacu pada daftar teknologi resmi F5 (Appendix 7.6) dan dikelompokkan pada 3 kategori mockup: Web Server, Language & Framework, Database.<br>- Teknologi yang memiliki Implied Technologies (mis. Apache Tomcat $\rightarrow$ Java Servlets/JSP, ASP.NET $\rightarrow$ IIS dan Microsoft Windows) otomatis ikut tercakup oleh F5 sehingga pelanggan tidak perlu memilihnya satu per satu |
| **FR-DWAF-0035** | Learning Period & Auto-Whitelist IP | Setelah Security Profile berstatus Active, sistem membuka learning period 7 hari (sesuai ketentuan tim Delivery Cloudeka) di mana IP manapun dapat mengakses domain/SSL profile dan setiap IP yang mengakses otomatis tercatat ke Whitelist IP profile tersebut. Setelah 7 hari, hanya IP yang tercatat pada Whitelist yang dapat mengakses; IP lain otomatis diblok. Secara teknis di F5: selama learning period semua request dicatat lewat Log Profile (event log ASM), dan pada hari ke-7 Whitelist diterapkan sebagai Address List pada Firewall Policy AFM (allow Whitelist + default deny) yang di-enforce ke Virtual Server (WI V2.0 skenario D; API List langkah 1-9). | - Learning period berlangsung 7 (tujuh) hari, dihitung sejak status Security Profile menjadi Active; durasi mengikuti ketentuan tim Operation Cloudeka dan tidak dapat diubah oleh User Admin<br>- Selama learning period, seluruh IP dapat mengakses domain/SSL profile (tidak ada pemblokiran berbasis Whitelist)<br>- Saat status profile menjadi Active, Automation Engine melampirkan Log Profile `"LogProfile-Learning Period"` (Application Security, Local Storage, `Request Type = All requests`) ke Virtual Server profile tersebut tanpa melepas log profile lain yang sudah terpasang (WI V2.0 D.2; API List langkah 1)<br>- Automation Engine membaca event log request secara berkala pada F5 DC1 dan DC2, membuang duplikat, dan mencatat setiap IP baru ke Whitelist profile tersebut (IP Address dan Added Date), tanpa input pengguna dan tanpa approval (WI V2.0 D.3; API List langkah 2)<br>- Pada akhir hari ke-7 ($7 \times 24$ jam sejak Active), Automation Engine membuat Address List dari Whitelist, membuat Firewall Policy (rule allow Whitelist + rule default deny), meng-enforce policy tersebut ke Virtual Server, lalu melepas Log Profile learning period; sejak itu IP di luar Whitelist otomatis diblok tanpa aksi manual (WI V2.0 D.4-D.6; API List langkah 3-5)<br>- Perubahan Whitelist setelah learning period (tambah/hapus IP oleh Super Admin atas ticket) memperbarui Address List dengan mengirim daftar lengkap alamat (`PATCH` menimpa seluruh isi), sehingga IP lain tidak ikut terhapus (WI V2.0 D.7; API List langkah 6-7)<br>- Penerapan ke F5 mengikuti pola deployment sinkron Multi-DC (DC1 & DC2) via Transaction API<br>- Pemblokiran IP di luar Whitelist dilakukan oleh Firewall Policy AFM (network firewall) yang terpisah dari WAF Policy ASM, sehingga berlaku sama pada profile berstatus Blocking maupun Transparent. Enforcement Mode Transparent hanya membuat pelanggaran WAF dicatat tanpa diblok dan tidak memengaruhi kontrol akses IP (sesuai dokumentasi F5: pada mode transparent traffic tidak diblok walau terjadi violation; AFM diproses sebelum ASM). Perilaku ini dijelaskan pada User Guide<br>- Whitelist bersifat per Security Profile; Address List dan Firewall Policy di F5 dibuat per profile (`AL-{profile}`, `FW-{profile}`) dan dihapus bersama profile saat Security Profile dihapus (API List langkah 9)<br>- IP pada Whitelist hanya dapat dihapus oleh Super Admin atas permintaan ticket pelanggan, dan IP baru setelah learning period hanya dapat ditambahkan oleh Super Admin atas ticket (FR-DWAF-0009)<br>- Pencatatan IP otomatis, enforcement di akhir learning period, dan perubahan Whitelist tercatat dalam audit log dengan timestamp<br>- Penjelasan learning period bagi pelanggan (masa 7 hari dan perilaku setelahnya) dijelaskan pada User Guide; tidak ada penambahan mockup<br>- **CATATAN RISIKO:** IP berbahaya yang mengakses selama learning period ikut tercatat ke Whitelist; mitigasinya adalah pelanggan open ticket agar Super Admin menghapus IP tersebut (FR-DWAF-0009)<br>- **CATATAN GAP TEKNIS:** (a) modul AFM (Network Firewall) harus ter-provision dan berlisensi di F5 DC1 dan DC2 — endpoint `address-list` terbukti berfungsi pada perangkat dev (Test Case US-021), status perangkat produksi perlu dikonfirmasi NOC; (b) bila tidak ada IP tercatat pada hari ke-7, policy tidak boleh di-enforce (akan memblok seluruh traffic) - perilaku fallback perlu keputusan produk; (c) log semua request selama 7 hari menambah volume event log pada perangkat multitenant — perlu kapasitas dan retensi; (d) langkah API List bertanda `"Needs Test"` belum diuji pada perangkat |

---

## 6. Non-Functional Requirements

| ID | Requirements | Needs |
| :--- | :--- | :--- |
| **NF-001** | Ketersediaan | 24 Jam/Hari dan 7 Hari/Minggu |
| **NF-002** | Bahasa | Seluruh teks harus menggunakan Bahasa Inggris |
| **NF-003** | Jejak Audit | Setiap perubahan harus dicatat dengan timestamp, identitas pengguna, dan detail aksi. Log harus bersifat immutable dan dapat diakses untuk keperluan compliance review. |
| **NF-004** | Keamanan | Pelanggan tidak memiliki akses langsung ke perangkat F5. Seluruh operasi harus dilakukan secara eksklusif melalui Service Portal. Kredensial perangkat F5 tidak boleh diekspos kepada end user. |
| **NF-005** | Skalabilitas | Satu SSL Client Profile untuk setiap domain profile, diberlakukan melalui validasi sistem. Sistem harus mendukung deployment SSL secara bersamaan (concurrent) di DC1 dan DC2 tanpa race condition. |

---

## 7. Appendix

### 7.1. Mockup

Berikut link untuk mockup yang digunakan sebagai acuan untuk di develop ke Service Portal Cloudeka sebagai berikut:
- Link mockup di Figma
- Link Mockup dalam bentuk PDF

#### 7.1.1 Security Profile - List (Empty State)

Halaman awal Security Profile pada Deka WAF ketika belum ada profile yang dibuat oleh pelanggan, lengkap dengan tombol "Create Profile".

*[Gambar 2. Security Profile - List (Empty State) — Antarmuka Cloudeka Service Portal menu Deka WAF > Security Profile menampilkan empty state "No security profile available - Create your first security profile for your web applications." dengan tombol "+ Create Profile".]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0001 - Melihat Daftar Security Profile | US-003 | Sebagai User Admin, saya melihat daftar semua security profile, sehingga saya dapat mengelola proteksi WAF aplikasi. |

#### 7.1.2 Security Profile - Create Profile Wizard

Tersedia Wizard 3 langkah untuk membuat security profile baru:
1. **Basic Information** (`Profile Name`, `Description`)
2. **Configuration** (`Domain`, `Node`, `Port` dropdown `80-HTTP`/`443-HTTPS`, `Server Technologies` — untuk penjelasan lebih lanjut bisa dilihat pada 7.6, `Enforcement Mode`, `Protection Features`)
3. **SSL Certificate** (upload `SSL Cert`, `SSL Key`, `SSL Chain`)

Dilengkapi panel Summary di sisi kanan yang melacak 7 item wajib (`Profile Name`, `Domain`, `Node`, `Port`, `Mode`, `SSL Cert`, `SSL Key`).

*[Gambar 3. Security Profile - Create Profile Wizard — Antarmuka form multi-step 3 langkah (Basic Information, Configuration, SSL Certificate) dengan panel Summary di sisi kanan dan tombol Cancel serta Create.]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0002 - Membuat Security Profile Baru; FR-DWAF-0018 - Validasi Form | US-002 | Sebagai User Admin, saya membuat WAF Security Profile baru dengan nama dan deskripsi, sehingga saya dapat mulai melindungi aplikasi baru. |
| | US-006 | Sebagai User Admin, saya menentukan Node backend saat membuat profile, sehingga traffic diarahkan ke origin server saya. |
| | US-007 | Sebagai User Admin, Portal membuat load-balancing pool untuk backend profile baru, sehingga traffic dapat didistribusikan ke Node. |
| | US-008 | Sebagai User Admin, saya menentukan Domain publik saat membuat profile, sehingga Portal melakukan provisioning Virtual Server untuknya. |
| | US-009 | Sebagai User Admin, saya mengunggah file SSL certificate saat membuat profile, sehingga traffic HTTPS dapat diinspeksi. |
| | US-010 | Sebagai User Admin, saya mengunggah file SSL private key yang berpasangan dengan certificate. |

#### 7.1.3 Security Profile - List (Populated)

Tabel daftar seluruh security profile milik pelanggan dengan kolom `Profile Name`, `Domain`, `Mode`, `Status`, `Created Date`, `Last Updated`, dan `Actions`, dilengkapi search, filter (Type/Mode/Status), dan pagination.

*[Gambar 4. Security Profile - List (Populated) — Tabel daftar security profile (production-web s.d. production-web-4) dengan indikator Mode (Blocking/Transparent), Status (Active/Inactive/Error), serta kontrol pencarian, filter, dan pagination.]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0001 - Melihat Daftar Security Profile | US-003 | Sebagai User Admin, saya melihat daftar semua security profile, sehingga saya dapat mengelola proteksi WAF aplikasi saya. |

#### 7.1.4 Security Profile - Overview

Halaman detail/overview satu security profile: header status profile, 4 stat tile (`Total Requests`, `Blocked`, `Block Rate`, `Avg Latency`), dan SATU card "Profile Configuration" 3 kolom berisi `Domain`/`Description`/`Node`, `Port`/`Mode`/`Bot Protection`, dan `SSL Certificate`/`Chain`/`Key`.

*[Gambar 5. Security Profile - Overview — Halaman detail profile production-web menampilkan banner informasi profile, 4 kartu metrik utama (37,000 Total Requests, 2,315 Blocked, 6.3% Block Rate, 18 ms Avg Latency), serta kartu Profile Configuration 3 kolom.]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0003 - Melihat Detail Security Profile; FR-DWAF-0014 - Tampilan Status Bot Protection (Overview Profile); FR-DWAF-0016 - Metrik Dashboard | US-033 | Sebagai User Admin, saya melihat statistik traffic dan blocking untuk profile saya di halaman Overview. |

#### 7.1.5 Security Profile - Edit (General Settings & WAF Mode)

Halaman edit profile berisi General Settings (`Profile Name`, `Domain`, `Description`), toggle WAF Mode (`Blocking`/`Transparent`) beserta penjelasan dampaknya, dan status SSL Certificates yang sudah terpasang.

*[Gambar 6. Security Profile - Edit (General Settings & WAF Mode) — Halaman Edit profile production-web dengan form General Settings, pilihan kartu WAF Mode (Blocking vs Transparent) beserta peringatan status aktif, status file SSL Certificates terpasang, tombol Replace Certificate, Cancel, dan Save Changes.]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0004 - Mengedit Security Profile - Pengaturan Umum; FR-DWAF-0005 - Mengedit Security Profile - Sertifikat SSL/TLS | US-004 | Sebagai User Admin, saya mengubah WAF Mode suatu profile antara Blocking dan Transparent, sehingga saya dapat mengontrol tingkat ketegasan enforcement. |

#### 7.1.6 Security Profile - Upload/Replace SSL Certificate

Modal upload SSL Certificate untuk mengganti/menambahkan `SSL Cert`, `SSL Key`, dan `SSL Chain` (maksimum 5 MB per file) pada profile yang sudah ada.

*[Gambar 7. Security Profile - Upload/Replace SSL Certificate — Modal dialog "Upload Certificate" di atas halaman Edit dengan area unggah file untuk SSL Cert, SSL Key, dan SSL Chain (maksimum 5 MB) serta tombol Cancel dan Add.]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0005 - Mengedit Security Profile - Sertifikat SSL/TLS | US-009 | Sebagai User Admin, saya mengunggah file SSL certificate saat membuat profile, sehingga traffic HTTPS dapat diinspeksi. |
| | US-010 | Sebagai User Admin, saya mengunggah file SSL private key yang berpasangan dengan certificate. |
| | US-011 | Sebagai Automation Engine, saya menginstal certificate yang diunggah ke file store F5, sehingga dapat dirujuk oleh SSL profile. |
| | US-012 | Sebagai Automation Engine, saya menginstal private key yang diunggah ke file store F5. |
| | US-013 | Sebagai Automation Engine, saya membuat SSL Client Profile yang merujuk cert/key yang sudah diinstal, dengan SNI diaktifkan, sehingga domain dapat melayani HTTPS. |

#### 7.1.7 Access Control - Whitelist IP (Empty State)

Kondisi awal halaman Whitelist IP ketika belum ada IP yang tercatat pada Access Control Profile (mis. profile baru pada awal learning period 7 hari). IP akan terisi otomatis dari traffic yang mengakses domain/SSL profile; pelanggan tidak menambahkan IP sendiri.

*[Gambar 8. Access Control - Whitelist IP (Empty State) — Tampilan Whitelist IP pada akun Useradmin dengan subjudul "Submit a ticket to request IP allowlisting." dan kartu empty state "Your IP list is empty".]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0009 - Mengelola Whitelist IP; FR-DWAF-0012 - Navigasi Access Control | US-016 | Sebagai Super Admin, saya menambahkan alamat IP terpercaya ke whitelist suatu profile atas permintaan ticket pelanggan setelah learning period berakhir, sehingga IP tersebut dapat mengakses domain profile dan tidak diblok. |

#### 7.1.8 Access Control - Whitelist IP (User Admin, Add & Smart Resolve)

Tampilan Whitelist IP untuk role Super Admin, dilengkapi fitur "Smart Resolve" (input PS Support ID untuk authorize IP secara instan) serta aksi delete per baris. Aksi tambah/hapus dilakukan atas permintaan ticket pelanggan dan langsung diterapkan ke F5 tanpa approval; tidak ada status Pending Approval. Penjelasan learning period 7 hari bagi pelanggan dijelaskan pada User Guide (tidak ada penambahan mockup).

*[Gambar 10. Access Control - Whitelist IP (User Admin, Add & Smart Resolve) — Tampilan Whitelist IP pada akun Superadmin dilengkapi input "Enter IP address" dengan keterangan "Add trusted IP support ID to immediately authorize IP address.", tombol "+ Add", serta tabel IP Address ter-whitelist beserta Added Date dan ikon hapus.]*

| Functional Requirement | User Story ID | User Story |
| :--- | :--- | :--- |
| FR-DWAF-0009 - Mengelola Whitelist IP | US-016 | Sebagai User Admin, saya menambahkan alamat IP terpercaya ke whitelist suatu profile setelah learning period berakhir, sehingga IP tersebut dapat mengakses domain profile dan tidak diblok. |
| | US-018 | Sebagai User Admin, saya menghapus sebuah IP dari whitelist atas permintaan ticket pelanggan. |
| FR-DWAF-0035 - Learning Period & Auto-Whitelist IP | US-060 | Sebagai Automation Engine, saya mencatat otomatis setiap IP yang mengakses domain profile selama 7 hari sejak profile Active ke whitelist profile tersebut. |
| | US-061 | Sebagai Automation Engine, setelah learning period 7 hari berakhir, saya memastikan hanya IP pada whitelist yang dapat mengakses dan IP lain diblok. |

---

### 7.2 Akun perangkat Virtual Server F5

**a. Prod (TBS):**
- **URL:** `https://172.29.172.62/`
- **Username:** `[terlampir di akun.waf.txt]`
- **Password:** `[terlampir di akun.waf.txt]`

**b. Dev:**
- **URL:** `https://103.125.100.126`
- **Username:** `[terlampir di akun.waf.txt]`
- **Password:** `[terlampir di akun.waf.txt]`

---

### 7.3 API List

Berikut adalah daftar lengkap iControl REST API yang digunakan oleh Automation Engine untuk mengorkestrasi operasi F5 BIG-IP. Berikut lampiran `icontrol-rest-api-user-guide-16-1-0.pdf`.

**Tabel 5. iControl REST API — Learning Period & Whitelist IP (Deka WAF)**

| Langkah (WI V2.0 skenario D) | Method & Endpoint | Payload / Catatan & Status Uji |
| :--- | :--- | :--- |
| **1. Aktifkan pencatatan semua request pada VS (saat profile Active)** — WI D.2 | `PATCH /mgmt/tm/ltm/virtual/{vsName}` | Tambahkan Log Profile `"LogProfile-Learning Period"` ke field `securityLogProfiles`. Log Profile ini dibuat sekali oleh NOC, bukan per pelanggan (WI D.1; Application Security, Storage Destination Local, `Request Type = All requests`; path GUI `Security > Event Logs > Logging Profiles`). `PATCH` menimpa seluruh array, sehingga Portal wajib mengirim ulang log profile yang sudah ada (mis. `LogProfile-BotDefense`) beserta yang baru. **Status:** Documented Only - Needs Test. |
| **2. Baca IP yang mengakses, polling berkala pada DC1 & DC2** — WI D.3 | `GET /mgmt/tm/asm/events/requests?$filter=virtualServerName eq '{vsName}'&$select=clientIp&$top={n}&$skip={m}` | Portal membuang duplikat IP lalu menyimpan IP baru ke Whitelist (database Portal). `$filter virtualServerName` dan `$top` Confirmed Live (Test Case US-034); log produksi $\pm 2,47$ juta baris sehingga filter dan pagination wajib. Pembatasan rentang waktu sejak profile Active dan `$select clientIp`: Needs Test. Interval polling ditetapkan saat implementasi. |
| **3. Buat Address List dari Whitelist (hari ke-7)** — WI D.4 | `POST /mgmt/tm/security/firewall/address-list` | `{"name":"AL-{profile}","addresses":[{"name":"198.51.100.77"}, ...]}`. Endpoint Confirmed Live pada perangkat dev (Test Case US-021). |
| **4. Buat Firewall Policy: allow Whitelist + default deny** — WI D.5 | `POST /mgmt/tm/security/firewall/policy`; `POST /mgmt/tm/security/firewall/policy/FW-{profile}/rules` (2 kali) | Policy: `{"name":"FW-{profile}"}`. Rule 1 (dibuat lebih dulu): `{"name":"default_deny","action":"drop","ipProtocol":"any"}`. Rule 2: `{"name":"allow_whitelist","action":"accept","ipProtocol":"any","source":{"addressLists":["/Common/AL-{profile}"]},"placeBefore":"default_deny"}`. Pola rule mengikuti contoh lab F5 (address list + `placeBefore default_deny`); struktur `"source"` Needs Test. |
| **5. Enforce ke VS saat learning period berakhir** — WI D.6 | `PATCH /mgmt/tm/ltm/virtual/{vsName}` | `{"fwEnforcedPolicy":"/Common/FW-{profile}"}` dan lepas `LogProfile-Learning Period` dari `securityLogProfiles` (sertakan log profile lain yang masih dipakai). Diterapkan sinkron ke DC1 & DC2 via Transaction API. Needs Test. |
| **6. Tambah IP setelah learning period (ticket, Super Admin)** — WI D.7 | `PATCH /mgmt/tm/security/firewall/address-list/AL-{profile}` | `{"addresses":[ ...seluruh alamat existing..., {"name":"<IP baru>"} ]}`. `PATCH` menimpa seluruh koleksi (sesuai lab F5), sehingga Portal wajib mengirim daftar lengkap dari database Portal; bila hanya IP baru yang dikirim, IP lain terhapus dan ikut terblok. Belum diuji pada perangkat. |
| **7. Hapus IP (ticket, Super Admin)** — WI D.7 | `PATCH /mgmt/tm/security/firewall/address-list/AL-{profile}` | Kirim daftar lengkap tanpa IP yang dihapus. Entri Address List dikenal berdasarkan alamatnya (tanpa ID terpisah seperti pada ASM `whitelist-ips`). Belum diuji pada perangkat. |
| **8. Smart Resolve: ambil IP dari PS Support ID** — WI D.8 | `GET /mgmt/tm/asm/events/requests?$filter=supportId eq '{supportId}'&$select=clientIp` | Hasil `clientIp` diproses seperti langkah 6. Hanya menemukan request yang diproses ASM (mis. di-block WAF); request yang di-drop AFM tidak memiliki Support ID. Field filter `supportId`: Needs Test. |
| **9. Hapus Security Profile (cleanup)** | `PATCH /mgmt/tm/ltm/virtual/{vsName}`; `DELETE /mgmt/tm/security/firewall/policy/FW-{profile}`; `DELETE /mgmt/tm/security/firewall/address-list/AL-{profile}` | Urutan: lepas `fwEnforcedPolicy` dari VS (nilai `none`) lebih dulu, baru hapus policy dan address list karena objek yang masih direferensikan tidak dapat dihapus. Needs Test. |

---

### 7.4 Tes Case API

Berikut terlampir tes case sesuai mockup yang dibuat pada `Test Case - Kesesuaian Mockup vs F5 API - Deka WAF.xlsx`. Per 05/10/2026 ditambahkan US-062 s.d. US-071 untuk Learning Period dan Whitelist IP; sebagian besar berstatus *Needs Test* (belum diuji pada perangkat), kecuali US-064 (*Confirmed Live* via US-021) dan sebagian US-063.

---

### 7.5 Working Instruction in F5

Berikut adalah working instruction sesuai dengan apa yang disampaikan oleh Mas Fauzan dan sudah divalidasi Mas Iwan: `WI Konsolidasi F5 BIG-IP WAF - Deka WAF V2.0.docx`.

Berikut pemetaan setiap langkah pada `WI Konsolidasi F5 BIG-IP WAF - Deka WAF V2.0.docx` terhadap Service Portal Cloudeka (Deka WAF), sesuai `mockup-waf-new.pdf` dan tabel Functional Requirement pada Bab 5-6: mana yang dieksekusi otomatis oleh Automation Engine (backend) begitu User Admin/Super Admin mengisi field terkait di Service Portal Cloudeka, dan mana yang masih memerlukan aksi manual langsung di perangkat F5.

#### 7.5.1 A. Onboarding Baru (dipicu saat User Admin membuat Security Profile baru mockup 7.1.2 Create Profile Wizard)

| Langkah (WI Konsolidasi V2.0) | Pelaksana & Mekanisme di Service Portal | Referensi FR / Mockup |
| :--- | :--- | :--- |
| **A.1 Membuat Node** | Otomatis oleh Automation Engine, dipicu input field `"Node"` pada Step 2 (Configuration) wizard Create Profile — User Admin hanya mengisi field Node. | FR-DWAF-0002, FR-DWAF-0025 |
| **A.2 Membuat Pool** | Otomatis bersamaan dengan pembuatan Node; Service Port pool member mengikuti field `"Port"` (dropdown `80-HTTP`/`443-HTTPS`) pada Step 2. | FR-DWAF-0002, FR-DWAF-0025, FR-DWAF-0033 |
| **A.3 Import SSL Certificate, Key & Chain** | Otomatis, dipicu upload 3 file (`SSL Cert`, `SSL Key`, `SSL Chain`) pada Step 3 wizard oleh User Admin. Urutan import (`Key` $\rightarrow$ `Certificate` $\rightarrow$ `Chain`) dieksekusi otomatis oleh backend. | FR-DWAF-0002, FR-DWAF-0005 |
| **A.4 Membuat SSL Client Profile** | Otomatis mengikuti langkah A.3; SNI diaktifkan otomatis pada profile. | FR-DWAF-0022 |
| **A.6 Membuat Virtual Server** | Otomatis, dipicu field `"Domain"` dan `"Port"` pada Step 2. Setiap Security Profile menghasilkan tepat 1 Virtual Server (model 1 VS untuk 1 tenant). | FR-DWAF-0002 |
| **A.7 Mengaktifkan Logging** | Otomatis sebagai bagian provisioning standar (diperlukan agar metrik Total Requests/Blocked/Block Rate pada Overview dapat tampil); tidak ada toggle terpisah bagi User Admin di mockup. | Tersirat pada FR-DWAF-0016 |
| **A.8 Membuat Security Policy (WAF Policy)** — pembuatan Policy & assign Virtual Server | Otomatis. | FR-DWAF-0002 |
| **A.8 — Enforcement Mode** | **INPUT User Admin:** wajib dipilih eksplisit (`Blocking`/`Transparent`) pada Step 2, tidak ada default. | FR-DWAF-0002 |
| **A.8 — Application Language** | **OTOMATIS:** backend selalu menetapkan `Application Language = Auto Detect` saat membuat WAF Policy. Tidak ada input User Admin/Super Admin karena field `"Application Language"` pada mockup lama telah digantikan oleh `"Server Technologies"`. | FR-DWAF-0034 |
| **A.8 — Server Technologies** | **INPUT User Admin:** field `"Server Technologies"` (multi-select, kategori Web Server / Language & Framework / Database) pada Step 2 (Configuration), bersifat optional. Nilai dipetakan ke Server Technologies pada WAF Policy (WI A.8) untuk membantu akurasi rule/learning. | FR-DWAF-0034 |
| **A.8 — Learning and Blocking Settings (Automatic Learning)** | **INPUT User Admin:** toggle `"Automatic Learning"`, independen dari Enforcement Mode. | FR-DWAF-0028 |
| **A.8 — Learning and Blocking Settings (Advanced Mode per-violation)** | **INPUT Super Admin** (opsional, disembunyikan default di wizard utama). | FR-DWAF-0029 |
| **A.9 Membuat GSLB Pool** | **BELUM** diotomasi Portal — bagian dari fitur GLB multi-site yang masih `"Diusulkan - Menunggu Keputusan Scope"`. Selama FR-DWAF-0027 belum disetujui, setiap profile diasumsikan single-site dan langkah ini tidak dijalankan Portal; **MANUAL** oleh NOC di F5 bila suatu saat dibutuhkan. | FR-DWAF-0027 (Diusulkan) |
| **A.10 Membuat GSLB Wide IP** | Sama seperti A.9 — **BELUM** diotomasi, **MANUAL** oleh NOC bila dibutuhkan, menunggu FR-DWAF-0027. | FR-DWAF-0027 (Diusulkan) |
| **A.11 Verifikasi Hasil — TLS handshake** | **OTOMATIS**, dijalankan pasca-deployment setelah SSL Client Profile ter-attach. | FR-DWAF-0023 |
| **A.11 Verifikasi Hasil — ping/nslookup ke Wide IP** | Hanya relevan untuk skenario GSLB (A.9/A.10) yang belum aktif; saat ini tidak berlaku. | Terkait FR-DWAF-0027 |

#### 7.5.2 C. Renewal SSL Certificate

| Langkah (WI Konsolidasi V2.0) | Pelaksana & Mekanisme di Service Portal | Referensi FR / Mockup |
| :--- | :--- | :--- |
| **C.1 Import Certificate Baru** | Otomatis, dipicu saat User Admin melakukan Upload/Replace SSL Certificate pada Security Profile existing. | FR-DWAF-0005, mockup 7.1.6 |
| **C.2 Membuat SSL Client Profile Baru** | Otomatis, mengikuti C.1. | FR-DWAF-0022 |
| **C.3 Update Virtual Server** | Otomatis — sistem mengganti referensi profile lama dengan yang baru pada VS terkait. | FR-DWAF-0022 |
| **C.4 Verifikasi** | Otomatis (TLS handshake check pasca-update). | FR-DWAF-0023 |
| **Peran User Admin pada skenario C** | Hanya mengunggah file SSL Certificate/Key/Chain baru melalui Portal; seluruh proses C.1-C.4 di F5 berjalan otomatis di belakang layar. | FR-DWAF-0005 |

#### 7.5.3 D. Whitelist IP (Learning Period 7 Hari & Support ID)

| Langkah (WI Konsolidasi V2.0) | Pelaksana & Mekanisme di Service Portal | Referensi FR / Mockup |
| :--- | :--- | :--- |
| **D.1-D.3 Learning period 7 hari: logging semua request & pencatatan IP** | **SETUP SEKALI oleh NOC (MANUAL):** Logging Profile `"LogProfile-LearningPeriod"` (D.1).<br>**OTOMATIS oleh Portal sejak profile Active:** melampirkan log profile ke VS (D.2) dan polling berkala event log DC1 & DC2 untuk mencatat IP unik ke Whitelist (D.3). Tidak ada input User Admin/Super Admin. | FR-DWAF-0035; API List langkah 1-2 |
| **D.4-D.6 Pemblokiran setelah hari ke-7** | **OTOMATIS pada hari ke-7:** Portal membuat Address List dari Whitelist (D.4), membuat Firewall Policy allow Whitelist + default deny (D.5), meng-enforce policy ke VS dan melepas log profile learning period (D.6). Tidak ada approval. | FR-DWAF-0035; API List langkah 3-5 |
| **D.7 Tambah IP setelah learning period** | **INPUT Super Admin atas ticket pelanggan:** memasukkan IP address langsung, atau memakai `"Smart Resolve"` (D.8). Portal memperbarui Address List secara otomatis dengan daftar lengkap alamat. User Admin hanya read-only dan diarahkan submit ticket. | FR-DWAF-0009; API List langkah 6; mockup 7.1.7-7.1.9 |
| **D.7 Hapus IP** | **INPUT Super Admin atas ticket pelanggan:** menghapus IP dari Whitelist. Portal memperbarui Address List secara otomatis (daftar lengkap tanpa IP tersebut). User Admin hanya open ticket (tidak ada aksi delete pada UI User Admin). | FR-DWAF-0009; API List langkah 7 |
| **D.8 Smart Resolve via Support ID** (menggantikan langkah manual WI lama: cari Support ID di Event Logs, klik Accept, buka Policy, Apply Policy) | **INPUT Super Admin:** PS Support ID.<br>**OTOMATIS:** Portal mencari request pada event log ASM, mengambil client IP, lalu menjalankan D.7. Hanya untuk request yang diproses ASM. | FR-DWAF-0009; API List langkah 8 |

#### 7.5.4 E. Konfigurasi Bot Defense (Bot Protection)

| Langkah (WI Konsolidasi V2.0) | Pelaksana & Mekanisme di Service Portal | Referensi FR / Mockup |
| :--- | :--- | :--- |
| **E.1 Membuat Logging Profile untuk Bot Defense** | **MANUAL**, satu kali oleh tim NOC/Cloudeka — bukan per pelanggan, bukan melalui Portal. | FR-DWAF-0008 |
| **E.2 Membuat Bot Defense Profile** | **MANUAL**, satu kali oleh tim NOC/Cloudeka. Satu objek profile standar (template `Balanced`, mode `Blocking`, nama `bot_sp_blocking_balanced`) dipakai bersama untuk semua pelanggan. | FR-DWAF-0008 |
| **E.3 Mengaitkan Bot Defense Profile ke Virtual Server** | **Otomatis oleh Portal** dipicu saat User Admin mencentang/melepas checkbox `"Bot Protection"` pada Create/Edit Security Profile; Portal attach/detach ke VS milik pelanggan via `PATCH /mgmt/tm/ltm/virtual/{vsName}`. | FR-DWAF-0008 |

---

### 7.6 Server Technologies Description

Field `"Server Technologies"` pada Step 2 (Configuration) wizard Create Security Profile bersifat optional dan multi-select. Daftar berikut bersumber dari dokumen F5 *"Server Technologies Description"* (WAF Declarative Policy, `clouddocs.f5.com/products/waf-declarative-policy/server_technology.html`) sebanyak 73 teknologi, dan yang dicantumkan di sini adalah **57 teknologi** yang termasuk 3 kategori mockup (**Web Server**, **Language & Framework**, **Database**); deskripsi diringkas dan diterjemahkan ke Bahasa Indonesia.

Kolom **Kategori** adalah pengelompokan yang ditetapkan pada SRS ini mengikuti 3 kategori mockup tersebut; dokumen F5 sendiri tidak mengelompokkan teknologi per kategori. Teknologi F5 di luar ketiga kategori (aplikasi/CMS, sistem operasi, perangkat jaringan, dan protokol) tidak dimasukkan. Kolom **Implied Technologies** menunjukkan teknologi yang secara implisit ikut tercakup oleh F5 saat teknologi tersebut dipilih.

| Kategori | Server Technology | Deskripsi | Implied Technologies |
| :--- | :--- | :--- | :--- |
| Web Server | Nginx | Web server yang juga dapat berfungsi sebagai reverse proxy, load balancer, mail proxy, dan HTTP cache. | - |
| Web Server | Apache/NCSA HTTP Server | Apache HTTP Server, perangkat lunak web server yang paling banyak digunakan di dunia. | - |
| Web Server | Apache Tomcat | Java Servlet Container open-source yang dikembangkan Apache Software Foundation. | Java Servlets/JSP |
| Web Server | IIS | Internet Information Services, web server ekstensibel dari Microsoft untuk keluarga Windows NT. | Microsoft Windows |
| Web Server | Jetty | Web server HTTP dan Servlet container berbasis Java. | Java Servlets/JSP |
| Web Server | JBoss | JBoss Enterprise Application Platform (JBoss EAP), application server Java EE untuk membangun dan menjalankan aplikasi Java ber-transaksi tinggi. | Java Servlets/JSP |
| Web Server | BEA Systems WebLogic Server | Oracle WebLogic Server, application server Java EE yang dikembangkan Oracle Corporation. | Java Servlets/JSP |
| Web Server | Macromedia JRun | Application server J2EE yang awalnya dikembangkan (1997) sebagai Java Servlet engine oleh Live Software, kemudian dimiliki Allaire. | Java Servlets/JSP |
| Web Server | Oracle Application Server | Solusi middleware terpadu Oracle: Oracle Containers for J2EE, Web Cache, HTTP Server, Forms, Reports, Portal, dan Discoverer. | - |
| Language & Framework | Python | Bahasa pemrograman interpreted tingkat tinggi untuk tujuan umum dengan penekanan pada keterbacaan kode. | - |
| Language & Framework | Spring Boot | Framework untuk membuat aplikasi dan layanan berbasis Spring yang siap produksi dengan konfigurasi minimal. | - |
| Language & Framework | Handlebars | Template engine untuk membangun semantic template. | - |
| Language & Framework | Mustache | Sistem template web yang sederhana. | - |
| Language & Framework | Prototype | Library JavaScript yang menyederhanakan pemrograman web sisi klien (Ajax dan DOM). | - |
| Language & Framework | Zend | Zend Server, distribusi PHP lengkap dan tersertifikasi yang dikelola Zend Technologies. | - |
| Language & Framework | Underscore.js | Library JavaScript berisi fungsi utilitas untuk tugas pemrograman umum dengan desain fungsional. | - |
| Language & Framework | Ember.js | Framework web JavaScript open-source berbasis pola Model-view-viewmodel untuk single-page application. | - |
| Language & Framework | ZURB Foundation | Framework front-end responsif (grid responsif, komponen UI HTML/CSS, template) yang bersifat open-source. | - |
| Language & Framework | ef.js | Template engine HTML dan framework dasar. | - |
| Language & Framework | Vue.js | Framework JavaScript open-source untuk membangun antarmuka pengguna dan single-page application. | - |
| Language & Framework | UIKit | Framework front-end yang ringan dan modular untuk membangun antarmuka web. | - |
| Language & Framework | RequireJS | Library JavaScript dan file loader untuk mengelola dependensi antar file JavaScript. | - |
| Language & Framework | React | Library JavaScript untuk membangun antarmuka pengguna, dikelola Facebook dan komunitas. | - |
| Language & Framework | MooTools | Framework JavaScript ringan berorientasi objek dengan lisensi MIT. | - |
| Language & Framework | Laravel | Framework web PHP open-source berpola MVC dan berbasis Symfony. | - |
| Language & Framework | GraphQL | Bahasa query dan manipulasi data open-source untuk API beserta runtime-nya. | - |
| Language & Framework | Google Web Toolkit | Kumpulan tool open-source untuk membuat dan memelihara aplikasi front-end JavaScript yang kompleks menggunakan Java. | - |
| Language & Framework | Express.js | Framework aplikasi web untuk Node.js (open-source, lisensi MIT) untuk membangun aplikasi web dan API. | - |
| Language & Framework | CodeIgniter | Framework web PHP open-source untuk pengembangan cepat situs web dinamis. | - |
| Language & Framework | Backbone.js | Library JavaScript dengan antarmuka RESTful JSON yang ringan; bergantung pada Underscore.js dan jQuery. | - |
| Language & Framework | AngularJS | Framework web front-end berbasis JavaScript (dikelola Google dan komunitas) untuk single-page application. | - |
| Language & Framework | JavaScript | Bahasa pemrograman tingkat tinggi, interpreted, dinamis, dan sesuai spesifikasi ECMAScript. | - |
| Language & Framework | JavaServer Faces (JSF) | Spesifikasi Java untuk membangun antarmuka web berbasis komponen. | Java Servlets/JSP |
| Language & Framework | Ruby | Bahasa pemrograman dinamis, reflektif, dan berorientasi objek untuk tujuan umum. | - |
| Language & Framework | Django | Framework web open-source berbasis Python dengan pola model-view-template (MVT). | - |
| Language & Framework | Node.js | Runtime JavaScript open-source lintas platform untuk mengembangkan berbagai tool dan aplikasi. | - |
| Language & Framework | Apache Struts | Framework web open-source untuk mengembangkan aplikasi web Java EE. | Java Servlets/JSP |
| Language & Framework | PHP | Bahasa scripting sisi server yang dirancang untuk pengembangan web dan dapat dipakai untuk tujuan umum. | - |
| Language & Framework | ASP.NET | Framework aplikasi web sisi server open-source untuk menghasilkan halaman web dinamis. | IIS, Microsoft Windows |
| Language & Framework | ASP | Active Server Pages (Classic ASP), script engine sisi server pertama dari Microsoft untuk halaman web dinamis. | IIS, Microsoft Windows |
| Language & Framework | Java Servlets/JSP | Program Java yang memperluas kemampuan server. | - |
| Language & Framework | jQuery | Library JavaScript lintas platform untuk menyederhanakan scripting HTML sisi klien. | - |
| Language & Framework | CGI | Common Gateway Interface, protokol standar agar web server dapat menjalankan program yang menghasilkan halaman web dinamis. | - |
| Language & Framework | SSI (Server Side Includes) | Bahasa scripting sisi server yang sederhana dan hampir hanya dipakai untuk web. | - |
| Language & Framework | Macromedia ColdFusion | Adobe ColdFusion, platform komersial untuk pengembangan aplikasi web secara cepat. | - |
| Language & Framework | XML | Extensible Markup Language, bahasa markup dengan aturan encoding dokumen yang dapat dibaca manusia dan mesin. | - |
| Database | CouchDB | Apache CouchDB, database open-source yang mudah digunakan dan berarsitektur skalabel. | - |
| Database | SQLite | Sistem database relasional dalam bentuk library C yang tertanam pada program (bukan client-server). | - |
| Database | Redis | Penyimpanan struktur data in-memory open-source (key-value terdistribusi) dengan durabilitas opsional. | - |
| Database | MongoDB | Database berorientasi dokumen, open-source, dan lintas platform. | - |
| Database | PostgreSQL | Database objek-relasional (ORDBMS) yang menekankan ekstensibilitas dan kepatuhan terhadap standar. | - |
| Database | IBM DB2 | Rangkaian produk server database dari IBM. | - |
| Database | Sybase/ASE | SAP ASE (Adaptive Server Enterprise), server database relasional yang dahulu bernama Sybase SQL Server. | - |
| Database | MySQL | Sistem manajemen database relasional (RDBMS) open-source. | - |
| Database | Oracle | Oracle Database, sistem manajemen database objek-relasional dari Oracle Corporation. | - |
| Database | Microsoft SQL Server | Sistem manajemen database relasional yang dikembangkan Microsoft. | - |
| Database | Elasticsearch | Mesin pencari (search engine) berbasis Lucene. | - |