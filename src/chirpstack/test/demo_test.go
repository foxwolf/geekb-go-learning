package test

import (
	"fmt"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("seee")
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	cobra.OnInitialize()
}

func TestDemo(t *testing.T) {

}
