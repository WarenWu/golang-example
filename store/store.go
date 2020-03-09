package store

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"golang-example/global"
	"golang-example/logger"
)

type URLStore struct{
	mutex   sync.RWMutex
	urls map[string]string
	file *os.File
	saveCh chan record
}

type record struct{
    key, url string
}

func New(fileName string) (*URLStore, error) {
	var s URLStore
	var err error
	s.urls = make(map[string]string)

	s.file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if err != nil {
		panic(err.Error())
	}
	err = s.load()

	for i := 0;i < global.NCPU ;i++  {
		go func(s *URLStore) {
			s.save()
		}(&s)
	}
	return &s, err
}

func (s *URLStore)Get(key string) (string, error) {
	s.mutex.Lock()
	if url, ok := s.urls[key]; ok {
        return url,nil
	}
	s.mutex.Unlock()
	return "", errors.New("The key have not relative url!")
	
}

func (s *URLStore)Set(url string) (string, error){
	s.mutex.Lock()
	key := genKey(len(s.urls))
	if _, ok := s.urls[key]; ok {
	return "",errors.New("The key url relatived has existed!")
	}
	s.urls[key] = url
	s.mutex.Unlock()
	go func() {
		s.saveCh <- record{key:key, url:url}
	}()
	return key, nil
}

func (s *URLStore)save() {
	for {
		r := <-s.saveCh
		e := json.NewEncoder(s.file)
		err := e.Encode(r)
		if err != nil {
			logger.FLogger.Println(err.Error())
		}
	}
}

func (s *URLStore)load() error {
	d := json.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record 
		if err = d.Decode(&r); err == nil {
			s.urls[r.key] = r.url
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20)
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}
