package lib

import (
	"container/list"
	"errors"
	"homework/lib/model"
)

type MusicManager struct {
	musics model.MusicEntry
}

type MusicManagerList struct {
	listMusics list.List
}

func NewMusuicManager() *MusicManagerList {
	return &MusicManagerList{*list.New()}
}

func (m *MusicManagerList) Len() int {
	return m.listMusics.Len()
}

func (m *MusicManagerList) Get(index int) (music *model.MusicEntry, err error) {
	if index < 0 || index >= m.listMusics.Len() {
		return nil, errors.New("Index out of range.")
	}

	elementNumber := 0
	for e := m.listMusics.Front(); e != nil; e = e.Next() {
		if elementNumber == index {
			return e.Value.(*model.MusicEntry), nil
		}
		elementNumber++
	}
	return nil, errors.New("Index out of range.")
}

func (m *MusicManagerList) Perv(index int) (music *model.MusicEntry, err error) {

	if index < 0 || index >= m.listMusics.Len() {
		return nil, errors.New("Index out of range.")
	}

	elementNumber := 0
	for e := m.listMusics.Front(); e != nil; e = e.Prev() {
		if elementNumber == index {
			//return e.Value.(*model.MusicEntry), nil
		}
		elementNumber++
	}
	return nil, errors.New("Index out of range.")
}

func (m *MusicManagerList) Find(name string) *model.MusicEntry {
	if m.listMusics.Len() == 0 {
		return nil
	}
	for e := m.listMusics.Front(); e != nil; e = e.Next() {
		if e.Value.(*model.MusicEntry).Name == name {
			return e.Value.(*model.MusicEntry)
		}
	}
	return nil
}

func (m *MusicManagerList) Add(music *model.MusicEntry) {
	println("Music was added: " + music.Name)
	m.listMusics.PushFront(music)
}

func (m *MusicManager) Remove(index int) *model.MusicEntry {
	return nil
}

func (m *MusicManager) RemoveByName(name string) *model.MusicEntry {
	return nil
}
