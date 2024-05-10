package beats_test

import (
	"fmt"
	"testing"

	"github.com/Nidasakinaa/beats/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertBiodata(t *testing.T) {
	mahasiswa := map[string]interface{}{
		"mhs_id":   primitive.NewObjectID(),
		"npm":      "Sela",
		"kelas":    "perempuan",
		"angkatan": "2022",
	}
	alamat := map[string]interface{}{
		"jalan":   "Kala",
		"kota":    "Bangka",
		"kodepos": "1234",
	}
	insertedID := beats.InsertBiodata(
		"Sela laa",
		"Perempuan",
		"Islam",
		"Logistik",
		mahasiswa,
		alamat,
	)
	fmt.Println("Inserted Biodata ID:", insertedID)
}

func TestGetAllBiodata(t *testing.T) {
	biodataa, err := beats.GetAllBiodata()
	if err != nil {
		t.Errorf("Error getting all Biodata: %v", err)
		return
	}

	fmt.Println("All Biodata:")
	for _, biodata := range biodataa {
		fmt.Println(biodata)
	}
}
