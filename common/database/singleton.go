package database

import (
	"sync"
)

/*Singleton strunct type*/
type Singleton struct{}

var instance *Singleton
var once sync.Once

/*GetInstance function singleton*/
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
