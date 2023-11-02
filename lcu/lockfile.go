package lcu

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Lockfile struct {
	ProcessName  string
	ProcessId    int
	Port         int
	Password     string
	Protocol     string
	LockfilePath string
	LastModified time.Time
	LockfileSize int64
	// Adding the LockfileSize field to the struct
}

func (lf *Lockfile) Wss() string {
	return fmt.Sprintf("wss://127.0.0.1:%d", lf.Port)
}
func (lf *Lockfile) Url() string {
	return fmt.Sprintf("%s://127.0.0.1:%d", lf.Protocol, lf.Port)
}
func (lf *Lockfile) TokenEncoded() string {
	return base64.RawStdEncoding.EncodeToString([]byte("riot:" + lf.Password))
}
func (lf *Lockfile) Update() error {
	info, err := os.Stat(lf.LockfilePath)
	if err != nil {
		return err
	}

	if info.Size() != lf.LockfileSize {
		return fmt.Errorf("Lockfile size has changed unexpectedly")
	}

	if info.ModTime().After(lf.LastModified) {
		content, err := os.ReadFile(lf.LockfilePath)
		if err != nil {
			return err
		}

		data := strings.Split(string(content), ":")

		processID, err := strconv.Atoi(data[1])
		if err != nil {
			return err
		}

		port, err := strconv.Atoi(data[2])
		if err != nil {
			return err
		}

		lf.ProcessName = data[0]
		lf.ProcessId = processID
		lf.Port = port
		lf.Password = data[3]
		lf.Protocol = data[4]
		lf.LastModified = info.ModTime()

	}
	return nil
}
func NewLockfile(lockfilePath string) (*Lockfile, error) {
	content, err := os.ReadFile(lockfilePath)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(content), ":")

	processID, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(data[2])
	if err != nil {
		return nil, err
	}

	lockfile := &Lockfile{
		ProcessName:  data[0],
		ProcessId:    processID,
		Port:         port,
		Password:     data[3],
		Protocol:     data[4],
		LockfilePath: lockfilePath,
		LastModified: time.Now(),
		LockfileSize: int64(len(content)),
	}

	lockfile.StartWatcher(time.Second)
	return lockfile, nil

}
func (lf *Lockfile) StartWatcher(interval time.Duration) {
	go func() {
		for range time.Tick(interval) {
			if err := lf.Update(); err != nil {
				fmt.Println("Error:", err)
			}
		}
	}()
}
