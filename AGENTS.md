<!-- BEGIN:pengetahuan-ringkas -->
<!-- pengetahuan-ringkas:v1 -->
## Knowledge Graph (graphify)

Repo ini punya knowledge graph di `graphify-out/`. Prosedur lengkap: skill `.claude/skills/pengetahuan-ringkas/SKILL.md`.

1. **Untuk pertanyaan arsitektur atau hubungan antarberkas, baca graf lebih dulu.** Buka `graphify-out/GRAPH_REPORT.md`, lalu `graphify query "<pertanyaan>"`, sebelum membuka banyak berkas satu per satu. Hemat token.
2. **Graf adalah peta, bukan sumber kebenaran.** Verifikasi setiap klaim ke `source_file` yang dikutip; edge `INFERRED`/`AMBIGUOUS` bukan fakta.
3. **Bila `graphify-out/STATUS.md` ada, graf baru berisi kode saja.** Jalankan `/graphify . --update` sebelum mengandalkannya untuk pertanyaan dokumen, hapus `STATUS.md`, lalu commit hasilnya.
4. **Perbarui graf setelah menambah, menghapus, atau mengubah substansial berkas**, dalam commit yang sama (`graphify update .` untuk kode, `/graphify . --update` untuk dokumen).
5. **Jangan edit `graphify-out/` dengan tangan** dan jangan memasang `graphify hook install`.
<!-- END:pengetahuan-ringkas -->
