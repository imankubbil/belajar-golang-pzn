package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")
	m.Run() // eksekusi semua unit test
	fmt.Println("Setelah Unit Test")
}

func TestHelloWorldIman(t *testing.T) {
	result := HelloWorld("Iman")
	if result != "Hello World Iman" {
		t.Fail()
	}

	fmt.Println("TestHelloWorldIman Done")
}

func TestHelloWorldKubbil(t *testing.T) {
	result := HelloWorld("Kubbil")
	if result != "Hello World Kubbil" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldKubbil Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Iman")
	assert.Equal(t, "Hello World Iman", result, "Result must be 'Hello World Iman'")
	fmt.Println("Running with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Iman")
	require.Equal(t, "Hello World Iman", result, "Result must be 'Hello World Iman'")
	fmt.Println("Running with require done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows OS")
	}

	result := HelloWorld("Iman")
	assert.Equal(t, "Hello World Iman", result, "Result must be 'Hello World Iman'")
}

func TestSubTest(t *testing.T) {
	t.Run("Iman", func(t *testing.T) {
		result := HelloWorld("Iman")
		require.Equal(t, "Hello World Iman", result)
	})
	t.Run("Kubbil", func(t *testing.T) {
		result := HelloWorld("Kubbil")
		require.Equal(t, "Hello World Kubbil", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Iman)",
			request:  "Iman",
			expected: "Hello World Iman",
		},
		{
			name:     "HelloWorld(Kubbil)",
			request:  "Kubbil",
			expected: "Hello World Kubbil",
		},
		{
			name:     "HelloWorld(Lekha)",
			request:  "Lekha",
			expected: "Hello World Lekha",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Iman")
	}
}

func BenchmarkHelloWorldKubbil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kubbil")
	}
}

func BenchmarkSubBenchMark(b *testing.B) {
	b.Run("Iman", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Iman")
		}
	})
	b.Run("Kubbil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kubbil")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Iman)",
			request: "Iman",
		},
		{
			name:    "HelloWorld(Kubbil)",
			request: "Kubbil",
		},
		{
			name:    "HelloWorld(Lekha)",
			request: "Lekha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}
