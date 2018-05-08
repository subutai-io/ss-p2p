package ptp

import (
	"testing"
	"net"
	"os"
	"reflect"
)

func TestGetDeviceBase(t *testing.T) {
	get := GetDeviceBase()
	if get != "vptp" {
		t.Error("Error. Return wrong value")
	}
}

func TestGetConfigurationTool(t *testing.T) {
	get := GetConfigurationTool()
	wait := []string{"/sbin/ip", "/bin/ip", "/usr/bin/ip"}
	for _, w := range wait {
		if get == w {
			return
		}
	}
	t.Error("Error: ", get)
}

func TestNewTAP(t *testing.T) {
	get1, err := newTAP("tool", "", "01:02:03:04:05:06", "255.255.255.0", 1)
	if get1 != nil {
		t.Error(err)
	}
	get2, err2 := newTAP("tool", "192.168.1.1", "-", "255.255.255.0", 1)
	if get2 != nil {
		t.Error(err2)
	}
}
/*
Generated Test_newTAP
Generated TestTAPLinux_GetName
Generated TestTAPLinux_GetHardwareAddress
Generated TestTAPLinux_GetIP
Generated TestTAPLinux_GetMask
Generated TestTAPLinux_GetBasename
Generated TestTAPLinux_SetName
Generated TestTAPLinux_SetHardwareAddress
Generated TestTAPLinux_SetIP
Generated TestTAPLinux_SetMask
Generated TestTAPLinux_Init
Generated TestTAPLinux_Open
Generated TestTAPLinux_Close
Generated TestTAPLinux_Configure
Generated TestTAPLinux_ReadPacket
Generated TestTAPLinux_WritePacket
Generated TestTAPLinux_Run
Generated TestTAPLinux_createInterface
Generated TestTAPLinux_setMTU
Generated TestTAPLinux_linkUp
Generated TestTAPLinux_linkDown
Generated TestTAPLinux_setIP
Generated TestTAPLinux_setMac
Generated TestTAPLinux_IsConfigured
Generated TestTAPLinux_MarkConfigured
Generated TestFilterInterface

package ptp

import (
	"net"
	"os"
	"reflect"
	"testing"
)
*/

/*
func TestGetDeviceBase(t *testing.T) {
	get := GetDeviceBase()
	if get != "vptp" {
		t.Error("Error. Return wrong value")
	}
}

func TestGetConfigurationTool(t *testing.T) {
	get := GetConfigurationTool()
	wait := []string{"/sbin/ip", "/bin/ip", "/usr/bin/ip"}
	for _, w := range wait {
		if get == w {
			return
		}
	}
	t.Error("Error: ", get)
}

func TestNewTAP(t *testing.T) {
	get1, err := newTAP("tool", "", "01:02:03:04:05:06", "255.255.255.0", 1)
	if get1 != nil {
		t.Error(err)
	}
	get2, err2 := newTAP("tool", "192.168.1.1", "-", "255.255.255.0", 1)
	if get2 != nil {
		t.Error(err2)
	}
}

func Test_newTAP(t *testing.T) {
	type args struct {
		tool string
		ip   string
		mac  string
		mask string
		mtu  int
	}
	tests := []struct {
		name    string
		args    args
		want    *TAPLinux
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTAP(tt.args.tool, tt.args.ip, tt.args.mac, tt.args.mask, tt.args.mtu)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTAP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTAP() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func TestTAPLinux_GetName(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if got := x.GetName(); got != tt.want {
				t.Errorf("TAPLinux.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_GetHardwareAddress(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
		want   net.HardwareAddr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if got := x.GetHardwareAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TAPLinux.GetHardwareAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_GetIP(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
		want   net.IP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if got := x.GetIP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TAPLinux.GetIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_GetMask(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
		want   net.IPMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if got := x.GetMask(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TAPLinux.GetMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_GetBasename(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if got := x.GetBasename(); got != tt.want {
				t.Errorf("TAPLinux.GetBasename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_SetName(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			x.SetName(tt.args.name)
		})
	}
}

func TestTAPLinux_SetHardwareAddress(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	type args struct {
		mac net.HardwareAddr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			x.SetHardwareAddress(tt.args.mac)
		})
	}
}

func TestTAPLinux_SetIP(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			x.SetIP(tt.args.ip)
		})
	}
}

func TestTAPLinux_SetMask(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	type args struct {
		mask net.IPMask
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			x.SetMask(tt.args.mask)
		})
	}
}

func TestTAPLinux_Init(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.Init(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_Open(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.Open(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_Close(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.Close(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_Configure(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.Configure(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.Configure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_ReadPacket(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Packet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			got, err := x.ReadPacket()
			if (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.ReadPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TAPLinux.ReadPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_WritePacket(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	type args struct {
		packet *Packet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.WritePacket(tt.args.packet); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.WritePacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_Run(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			x.Run()
		})
	}
}

func TestTAPLinux_createInterface(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.createInterface(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.createInterface() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_setMTU(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.setMTU(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.setMTU() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_linkUp(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.linkUp(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.linkUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_linkDown(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.linkDown(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.linkDown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_setIP(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.setIP(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.setIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_setMac(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if err := x.setMac(); (err != nil) != tt.wantErr {
				t.Errorf("TAPLinux.setMac() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTAPLinux_IsConfigured(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			if got := x.IsConfigured(); got != tt.want {
				t.Errorf("TAPLinux.IsConfigured() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTAPLinux_MarkConfigured(t *testing.T) {
	type fields struct {
		IP         net.IP
		Mask       net.IPMask
		Mac        net.HardwareAddr
		Name       string
		Tool       string
		MTU        int
		file       *os.File
		Configured bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &TAPLinux{
				IP:         tt.fields.IP,
				Mask:       tt.fields.Mask,
				Mac:        tt.fields.Mac,
				Name:       tt.fields.Name,
				Tool:       tt.fields.Tool,
				MTU:        tt.fields.MTU,
				file:       tt.fields.file,
				Configured: tt.fields.Configured,
			}
			x.MarkConfigured()
		})
	}
}

func TestFilterInterface(t *testing.T) {
	type args struct {
		infName string
		infIP   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInterface(tt.args.infName, tt.args.infIP); got != tt.want {
				t.Errorf("FilterInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
