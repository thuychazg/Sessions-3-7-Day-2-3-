package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIPScanner_Scan(t *testing.T){

	scanner := &IPScanner{}

	result, err := scanner.Scan("127.0.0.1")

	assert.NoError(t, err)
	assert.NotNil(t, result)

	data := result.(map[string]interface{})

	assert.Equal(
		t,
		"127.0.0.1",
		data["ip_address"],
	)
}



func TestPortScanner_Scan(t *testing.T){

	scanner := &PortScanner{}

	result, err := scanner.Scan("127.0.0.1")

	assert.NoError(t, err)
	assert.NotNil(t,result)

}


func TestPortScanner_PublicIP(t *testing.T){

	scanner := &PortScanner{}

	_, err := scanner.Scan("8.8.8.8")

	assert.Error(t,err)

}
