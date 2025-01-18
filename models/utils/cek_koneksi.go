package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func CekKoneksi() {
	//cek koneksi localhost:5774

}

// Fungsi umum untuk mengonversi array byte menjadi slice dari struct
func BytesToStructSlice(data []byte, result interface{}) error {
	// Pastikan result adalah pointer ke slice
	sliceValue := reflect.ValueOf(result)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result harus berupa pointer ke slice")
	}

	// Ambil tipe elemen dari slice
	elementType := sliceValue.Elem().Type().Elem()
	elementSize := binary.Size(reflect.New(elementType).Elem().Interface())
	if elementSize == -1 {
		return fmt.Errorf("ukuran elemen struct tidak valid")
	}

	// Pastikan panjang data sesuai dengan ukuran struct
	if len(data)%elementSize != 0 {
		return fmt.Errorf("panjang data tidak sesuai dengan ukuran struct")
	}

	// Hitung jumlah elemen dan alokasi slice
	numElements := len(data) / elementSize
	newSlice := reflect.MakeSlice(sliceValue.Elem().Type(), numElements, numElements)

	// Membaca data ke dalam slice
	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, newSlice.Interface())
	if err != nil {
		return fmt.Errorf("gagal membaca data: %v", err)
	}

	// Menetapkan slice hasil konversi ke result
	sliceValue.Elem().Set(newSlice)
	return nil
}
