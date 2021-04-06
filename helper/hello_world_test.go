package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)



func TestMain(m *testing.M){
	//before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")

	//test main hanya jalan di satu package saja jika ingin menjalankan di package lain harus buat testmain nya lagi
	//ini bisa untuk database untuk di before nya
}
func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Irsan")
	assert.Equal(t, "Hello Irsan", result, "Result must be , Hello Irsan")
	fmt.Println("TestHelloWorld with Assert Done")

	//testify menyeduakan dua package untuk assertion , yaitu assert dan require
	//saat kita menggunakan assert jikan pengecekan gagal , maka assert akan memanggil Fail()
	//sedangkan require jika pengecekan gagal akan memanggil FailNow() sehingga ekesekusi unit test tidak di lanjutkan

}


func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Irsan")
	require.Equal(t, "Hello Irsan", result, "Result must be , Hello Irsan")
	//fmt tidak di panggil karena ini failnow()
	fmt.Println("TestHelloWorld with require Done")

	//testify menyeduakan dua package untuk assertion , yaitu assert dan require
	//saat kita menggunakan assert jikan pengecekan gagal , maka assert akan memanggil Fail()
	//sedangkan require jika pengecekan gagal akan memanggil FailNow() sehingga ekesekusi unit test tidak di lanjutkan

}



func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Irsan")
	require.Equal(t, "Hello Irsan", result , "Result must be 'Hello Irsan")
}

func TestHelloWorldTaufik(t *testing.T){
	result := HelloWorld("irsan")
	if result != "Hello irsan" {
		//error
		//t.Fail()
		t.Error("Result must be Hello irsan")
	}
	//go test -v itu fireboost ini akan menjalankan semua function
	//jika ingin spesifik menjalankan function tertentu di unit test menggunakan go test -v -run TestNamaFunction
	//golang by default ngerunning test unit nya by prefix jadi kalau ada yang mirip di akan di jalankan juga otomatis seperti contoh
	//TestHelloWorld dan TestHelloWorldIrsan ini akan di jalankan bersamaan karena prefix nya mirip sesuai awalan nya
	//go test tidak bisa di jalankan di luar
	//folder package/project nya jika ingin menjalankan di luar dan menjalankan semua unit test nya
	//gunakan command go test ./...
	fmt.Println("TestHelloWorld Done")
}

func TestTableHelloWorldSubTest(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name:      "Taufik",
			request:   "Taufik",
			expected: "Hello Taufik",
		},
		{
			name:      "Irsan",
			request:   "Irsan",
			expected: "Hello Irsan",
		},
	}

	//sub test pakai iterasi tak perlu copas tiap function nya
	for _, test := range tests {
		t.Run(test.name, func (t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkSub(b *testing.B){
	b.Run("Irsan", func(b *testing.B) {
		for i:= 0 ; i < b.N ; i++ {
			HelloWorld("Irsan")
		}
	})

	b.Run("Taufik", func(b *testing.B) {
		for i:= 0 ; i < b.N ; i++ {
			HelloWorld("Taufik")
		}
	})
}


func BenchmarkTableSub(b *testing.B){
	benchmarks := []struct{
		name string
		request string
		expected string
	}{
		{
			name:      "Taufik",
			request:   "Taufik",

		},
		{
			name:      "Taufik Irsan",
			request:   "Taufik Irsan",

		},
		{
			name:      "Sasageyo Irsan",
			request:   "Sasageyo Irsan",

		},
	}

	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B) {
			for i:=0 ; i< b.N ; i++{
				HelloWorld(benchmark.request)
			}
		})
	}
}
//Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0 ; i<b.N ; i++{
		HelloWorld("Taufik")
	}
	//untuk menjalankan seluhur benchmark di module bisa menggunakan command go test -v -benc=
	//jika hanya ingin menjalankan benchmark tanpa unit bisa menggunakan command go test -v -run=NotMathUnitTest -bench=
	//jika hanya ingin menjalankan benchmark tertentu bisa menggunakan command go test -v -run=NotMathUnitTest -bench=NamaFuncBenchmark
	//jika ingin menjalankan benchmark di root module dan ingin semua module di jalankan bisa menggunakan command go test-v -bench=. ./...
}

//Benchmark
func BenchmarkHelloWorldIrsan(b *testing.B) {
	for i := 0 ; i<b.N ; i++{
		HelloWorld("IrsanEdogawa")
	}
	//untuk menjalankan seluhur benchmark di module bisa menggunakan command go test -v -benc=
	//jika hanya ingin menjalankan benchmark tanpa unit bisa menggunakan command go test -v -run=NotMathUnitTest -bench=
	//jika hanya ingin menjalankan benchmark tertentu bisa menggunakan command go test -v -run=NotMathUnitTest -bench=NamaFuncBenchmark
	//jika ingin menjalankan benchmark di root module dan ingin semua module di jalankan bisa menggunakan command go test-v -bench=../...
}
func TestSubTest(t *testing.T){
	t.Run("Taufik", func(test *testing.T){
		result := HelloWorld("Taufik")
		require.Equal(t, "Hello Taufik", result , "Result must be 'Hello Taufik")
	})

	//jika hanya ingin menjalankan Subtest nya saja bisa menggunakan command go test -v -run TestNamaFunction/NamaSubTest (disarankan karena lebih spesifik)
	//jika ingin menjalankan semua sub test di semua function kita bisa menggunakan command go test -run /NamaSubTest
	t.Run("Irsan",func(*testing.T){
		result := HelloWorld("Irsan")
		require.Equal(t, "Hello Irsan", result , "Result must be 'Hello Irsan")
	})

}

func TestHelloWorldIrsan(t *testing.T){
	result := HelloWorld("Taufik")
	if result != "Hello Taufik" {
		//error
		//t.FailNow()
		t.Fatal("Result must be Taufik")
	}
	//terdapat dua function untuk menggagalkan unit test yaitu Fail() dan failNow()
	//beda dari kedua function tersebut , kalau fail() tetep melanjutkan eksekusi unit test namun di akhir ketika selesai untit test di anggap gagal
	//kalau untuk Failnow akan gagal unit test saat ini itu juga tanpa melanjutkan eksekusi unit test
	fmt.Println("TestHelloWorldIrsan Done")

	//ada juga function Error() dalam menangkap error unit test , error seperti melakukan log(print)error,
	//setelah melakukan log error dia akan secara otomatis memanggil function Fail sehingka mengakibatkan unit test di anggap gagal
	//namu karena hanya memanggil fail artinya eksekusi unit test akan tetap berjalan sampai selesai

	//ada juga function namanya Fatal() mirip dengan error hanya saja setelah melakukan log
	//error dia akan memanggil FailNow yang mana akan membuat eksekusi unit test berhenti
}
