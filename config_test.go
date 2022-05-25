package googleplay

import (
	"os"
	"testing"
	"time"
)

func TestCheckinArmeabi(t *testing.T) {
	err := checkin(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckinArm64(t *testing.T) {
	err := checkin(2)
	if err != nil {
		t.Fatal(err)
	}
}

func checkin(id int64) error {
	platform := Platforms[id]
	config, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	device, err := Phone.Checkin(platform)
	if err != nil {
		return err
	}
	platform += ".json"
	if err := device.Create(config, "googleplay", platform); err != nil {
		return err
	}
	time.Sleep(Sleep)
	return nil
}

func TestCheckinX86(t *testing.T) {
	err := checkin(0)
	if err != nil {
		t.Fatal(err)
	}
}
