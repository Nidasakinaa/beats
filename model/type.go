package beats

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID       primitive.ObjectID `bson:"mhs_id,omitempty" json:"mhs_id,omitempty"`
	NPM      string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Kelas    string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Angkatan string             `bson:"angkatan,omitempty" json:"angkatan,omitempty"`
}

type Alamat struct {
	Jalan   string `bson:"jalan,omitempty" json:"jalan,omitempty"`
	Kota    string `bson:"kota,omitempty" json:"kota,omitempty"`
	KodePos string `bson:"kodepos,omitempty" json:"kodepos,omitempty"`
}

type Biodata struct {
	ID           primitive.ObjectID `bson:"bio_id,omitempty" json:"bio_id,omitempty"`
	NamaLengkap  string             `bson:"nama,omitempty" json:"nama,omitempty"`
	JenisKelamin string             `bson:"jk,omitempty" json:"jk,omitempty"`
	Agama        string             `bson:"agama,omitempty" json:"agama,omitempty"`
	Prodi        string             `bson:"prodi,omitempty" json:"prodi,omitempty"`
	DataDiri     Mahasiswa          `bson:"datadiri,omitempty" json:"datadiri,omitempty"`
	Alamat       Alamat             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}
