package ssmmu

import "testing"

func setup() (mmu *SSMMU) {
	mmu = NewSSMMU("udp", "127.0.0.1:2403")
	return
}

func TestAddPort(t *testing.T) {
	mmu := setup()
	succ, err := mmu.Add(8592, "m")
	if err != nil {
		t.Error(err)
	}
	if !succ {
		t.Fatal("add port should be succ")
	}
}

func TestRemovePort(t *testing.T) {
	mmu := setup()
	succ, err := mmu.Remove(8592)
	if err != nil {
		t.Error(err)
	}
	if !succ {
		t.Fatal("remove port should be succ")
	}
}

func TestPing(t *testing.T) {
	mmu := setup()
	succ, duration, err := mmu.Ping()
	if err != nil {
		t.Error(err)
	}
	if !succ {
		t.Fatal("shoud recv pong")
	}
	t.Logf("ping time: %.3fms", duration.Seconds()*1000)
}
