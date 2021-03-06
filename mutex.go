package main

import (
	"math/rand"
	"sync"
)

var globalState int64
var mutex *sync.Mutex
var spinLock SpinLock
var spinLock2 SpinLockBit

func init() {
	mutex = new(sync.Mutex)
	spinLock = NewSpinLock()
	spinLock2 = NewSpinLockBit()
}

func Calculation() {
	globalState += rand.Int63n(1000)
}

func LockedWithMutex() {
	mutex.Lock()
	Calculation()
	mutex.Unlock()
}

func LockedWithDeferMutex() {
	mutex.Lock()
	defer mutex.Unlock()
	Calculation()
}

func LockedWithSpinLock() {
	spinLock.Lock()
	Calculation()
	spinLock.Unlock()
}

func LockedWithDeferSpinLock() {
	spinLock.Lock()
	defer spinLock.Unlock()
	Calculation()
}

func LockedWithSpinLockBit() {
	spinLock2.Lock()
	Calculation()
	spinLock2.Unlock()
}

func LockedWithDeferSpinLockBit() {
	spinLock2.Lock()
	defer spinLock2.Unlock()
	Calculation()
}
