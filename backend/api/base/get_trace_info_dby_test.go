package base

import (
	"fmt"
	"testing"

	"brilliance/client_e2e_test/blockchain/common/config"
)

func TestBaseGetDby_GetInfo(t *testing.T) {
	config.InitConfig("../../../config/config.yaml")

}

func TestChangeIp(t *testing.T) {
	fmt.Println(ChangeIp("http://10.0.90.53/group1/M00/00/0C/CgBaNVttCj2AEyKzAAAbcao60uM745.jpg"))
}
