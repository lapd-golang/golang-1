package model

import (
	"testing"
	"log"
)


// args: device mode speed databits parity stopbits 
func TestGetServerArgs(t*testing.T){
	cases := []struct {
		conf ConfigDf1;
		slice []string;
		err error
	}{
		{ConfigDf1{
			DeviceName: "df1",
			Endpoint: "/dev/ttyUSB0",
			AccessOptions: AccessOptionDf1{
				BaudRate: 19200,
				DataBits: 8,
				StopBits: 1,
				Parity: 0,
				Mode: "full",
			},
		}, []string{"/dev/ttyUSB0", "full", "19200", "8", "0", "1"}, nil},
	}

	for _, c := range cases {
		sli, reserr := GetServerArgs(&c.conf)
		log.Println(sli, reserr)
		if reserr != nil {
			t.Errorf("in:(%q), out:(%q), want:(%q)", c.conf, sli, c.slice)
		}
	}
}

