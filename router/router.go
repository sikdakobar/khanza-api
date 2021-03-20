package router

import (
	"net/http"
	"simpus/app/harga"
	"simpus/app/keuangan"
	"simpus/app/obat"
	"simpus/app/pegawai"
	"simpus/app/pelayanan/pendaftaran/antrian"
	"simpus/app/pelayanan/pendaftaran/pasien"
	"simpus/app/settings/config"

	"github.com/gorilla/mux"
)

// Init is
func Init() {
	r := mux.NewRouter()

	// Pelayanan > Pendaftaran > Pasien
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien", pasien.Index).Methods("GET")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien", pasien.Store).Methods("POST")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien/{id}", pasien.Show).Methods("GET")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien/{id}", pasien.Update).Methods("PUT")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien/{id}", pasien.Destroy).Methods("DELETE")
	r.HandleFunc("/api/pelayanan/pendaftaran/alamatpasien/{id}", pasien.AlamatStore).Methods("POST")
	r.HandleFunc("/api/pelayanan/pendaftaran/alamatpasien/{id}/{index}", pasien.AlamatUpdate).Methods("PUT")

	// Pelayanan > Rekam Medis
	r.HandleFunc("/api/pelayanan/pasien/{id}/tindakan", pasien.RekamMedisStore).Methods("POST")
	r.HandleFunc("/api/pelayanan/pasien/{id}/tindakan/{index}", pasien.RekamMedisUpdate).Methods("PUT")

	// Pelayanan > Pendaftaran > Antrian
	r.HandleFunc("/api/pelayanan/pendaftaran/antrian", antrian.Index).Methods("GET")
	r.HandleFunc("/api/pelayanan/pendaftaran/antrian", antrian.Store).Methods("POST")
	r.HandleFunc("/api/pelayanan/pendaftaran/antrian/list", antrian.ListAntrian).Methods("GET")

	// Pegawai
	r.HandleFunc("/api/pegawai", pegawai.Index).Methods("GET")
	r.HandleFunc("/api/pegawai", pegawai.Store).Methods("POST")
	r.HandleFunc("/api/pegawai/{id}", pegawai.Show).Methods("GET")
	r.HandleFunc("/api/pegawai/{id}", pegawai.Update).Methods("PUT")
	r.HandleFunc("/api/pegawai/{id}", pegawai.Destroy).Methods("DELETE")
	r.HandleFunc("/api/pegawai/{id}", pegawai.TMTUpdate).Methods("PATCH")

	// Harga
	r.HandleFunc("/api/harga", harga.Index).Methods("GET")
	r.HandleFunc("/api/harga", harga.Store).Methods("POST")
	r.HandleFunc("/api/harga/{id}", harga.Update).Methods("PUT")
	r.HandleFunc("/api/harga/{id}", harga.Destroy).Methods("DELETE")

	// Obat
	r.HandleFunc("/api/obat", obat.Index).Methods("GET")
	r.HandleFunc("/api/obat", obat.Store).Methods("POST")
	r.HandleFunc("/api/obat/{id}", obat.Update).Methods("PUT")
	r.HandleFunc("/api/obat/{id}", obat.Destroy).Methods("DELETE")
	r.HandleFunc("/api/obat/{id}/batch", obat.BatchObatStore).Methods("POST")
	r.HandleFunc("/api/obat/{id}/{index}/batch", obat.BatchObatUpdate).Methods("PUT")
	r.HandleFunc("/api/obat/{id}/{index}/batch", obat.BatchObatDestroy).Methods("DELETE")

	// Keuangan > COA
	r.HandleFunc("/api/coa", keuangan.IndexCOA).Methods("GET")
	r.HandleFunc("/api/coa", keuangan.StoreCOA).Methods("POST")
	r.HandleFunc("/api/coa/{id}", keuangan.ShowCOA).Methods("GET")
	r.HandleFunc("/api/coa/{id}", keuangan.UpdateCOA).Methods("PUT")
	r.HandleFunc("/api/coa/{id}", keuangan.DestroyCOA).Methods("DELETE")
	r.HandleFunc("/api/coa/{id}/subaccount", keuangan.StoreSubAccountCOA).Methods("POST")
	r.HandleFunc("/api/coa/{id}/subaccount/{index}", keuangan.DestroySubAccountCOA).Methods("DELETE")

	r.HandleFunc("/api/db/users", config.GetUser).Methods("GET")
	r.HandleFunc("/api/db", config.IndexDB).Methods("GET")
	r.HandleFunc("/api/db/collections", config.IndexCollection).Methods("GET")
	// r.HandleFunc("/api/db/collection", config.Store).Methods("POST")

	http.ListenAndServe(":1234", r)
}
