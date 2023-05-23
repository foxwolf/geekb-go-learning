package main

import (
	"fmt"
	"geekb-go-learn/src/chirpstack/demo1"
	"geekb-go-learn/src/chirpstack/demo2"
	"time"

	// "geekb-go-learn/src/chirpstack/test"
	"github.com/go-zeromq/zmq4/security/null"
	log "github.com/sirupsen/logrus"
)

// var version string // set by the compiler

type UnsupportedConfigError string

// Error returns the formatted configuration error.
func (str UnsupportedConfigError) Error() string {
	return fmt.Sprintf("Unsupported Config Type %q", string(str))
}

func T1() error {
	return UnsupportedConfigError("Errfxxae")
}

func main() {
	err := T1()
	switch err.(type) {
	case UnsupportedConfigError:
		fmt.Printf("%s\n", err)
	}

	demo2.StartDemo2()
	demo1.StartDemo1()

	log.SetFormatter(&log.TextFormatter{TimestampFormat: nil})
	log.WithField("aebc", "geekb").Info("signal received")
	// test.Execute(version)
	// test.Ta()
	// test.Init()
}
