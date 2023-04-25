Resume Materi "CI and CD"
- CI (continous integration) adalah proses untuk mengintegrasi berbagai code dari sumber yang berbeda untuk dilakukan build dan test secara otomatis apakah software berjalan sesuai algoritma code semestinya
- CD (continous delivery) adalah proses setelah CI, dimana akan mempersiapkan perubahan code yang lulus testing dan integrasi ke tahap deploy staging secara manual
- CD (continous deployment) adalah tahap terakhir, dimana proses deploy production secara otomatis dari pipeline
- CI/CD pipeline menjadi penghubung antara tim developer dengan tim operasional dengan melewati 3 tahap tersebut
- cycle untuk continous delivery adalah unit test, platform test, deliver to staging, application acceptance tests, deploy to production, post deploy tests. semua dilakukan secara otomatis kecuali tahap deploy to production yang dilakukan secara manual
- cycle untuk continous deployment adalah unit test, platform test, deliver to staging, application acceptance tests, deploy to production, post deploy tests. semua dilakukan secara otomatis
- contoh tools untuk CI/CD adalah git actions, git lab, jenkins, aws codebuild, azure devops, dkk