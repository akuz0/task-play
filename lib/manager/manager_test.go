package lib

import (
	"homework/lib/model"
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusuicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty")
	}

	m0 := &model.MusicEntry{
		Id:       "1",
		Name:     "My Heart Will Go On",
		Duration: "60",
	}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}

	if m.Id != m0.Id ||
		m.Name != m0.Name ||
		m.Duration != m0.Duration {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusiceManager.Get() failed.", err)
	}
}
