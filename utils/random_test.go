package utils

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	n := 5
	var wg sync.WaitGroup
	values := make(chan int64)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			values <- RandomInt(0, rand.Int63n(10000000))
		}(&wg)
	}
	go func() {
		wg.Wait()
		close(values)
	}()
	for v := range values {
		require.NotZero(t, v)
	}
}

func TestRandomString(t *testing.T) {
	n := 5
	values := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			values <- RandomString(50)
		}(&wg)
	}
	go func() {
		wg.Wait()
		close(values)
	}()
	for v := range values {
		require.NotEmpty(t, v)
		require.Len(t, v, 50)
	}
}

func TestStrToBinary(t *testing.T) {
	n := 5
	var wg sync.WaitGroup
	values := make(chan []byte)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			values <- StrToBinary(RandomString(20), 10)
		}(&wg)
	}
	go func() {
		wg.Wait()
		close(values)
	}()
	for v := range values {
		require.NotEmpty(t, v)
		require.IsType(t, []byte{}, v)
	}
}
