package store

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"urlconv/global"
	"urlconv/logger"
)

type URLStore struct{
	mutex   sync.RWMutex
	urls map[string]string
	file *os.File
	saveCh chan record
}

type record struct{
    Key, Url string
}

func New(fileName string) (*URLStore, error) {
	var s URLStore
	var err error
	s.urls = make(map[string]string)
	s.saveCh = make(chan record)
	s.file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if err != nil {
		panic(err.Error())
	}

	err = s.load()

	for i := 0;i < global.NCPU ;i++  {
		go func() {
			s.save()
		}()
	}
	return &s, err
}

func (s *URLStore)Get(key string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if url, ok := s.urls[key]; ok {
		return url, nil
	}
	return "", errors.New("The Key have not relative Url!")
	
}

func (s *URLStore)Set(url string) (string, error){
	s.mutex.Lock()
	defer s.mutex.Unlock()
	key := genKey(len(s.urls))
	if _, ok := s.urls[key]; ok {
	return "",errors.New("The Key Url relatived has existed!")
	}
	s.urls[key] = url
	go func() {
		s.saveCh <- record{Key: key, Url:url}
	}()
	return key, nil
}

func (s *URLStore)save() {
	for {
		r := <-s.saveCh
		e := json.NewEncoder(s.file)
		err := e.Encode(r)
		if err != nil {
			logger.RunLogger.Println(err.Error())
		}
	}
}

func (s *URLStore)load() error {
	d := json.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record 
		if err = d.Decode(&r); err == nil {
			s.urls[r.Key] = r.Url
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
