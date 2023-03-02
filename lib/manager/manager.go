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
	//return &MusicManager{make([]model.MusicEntry, 0)}
	//return &MusicManager{make model.MusicEntry, 0)}
	return &MusicManagerList{*list.New()}
}

func (m *MusicManagerList) Len() int {
	return m.listMusics.Len()
}

func (m *MusicManagerList) Get(index int) (music *model.MusicEntry, err error) {
	//index = index - 1
	if index < 0 || index >= m.listMusics.Len() {
		return nil, errors.New("Index out of range.")
	}

	//for e := m.listMusics.Front(); e != nil; e = e.Next() {
	//	if index == 0 {
	//		return e.Value.(*model.MusicEntry), nil
	//	}
	//	index = index - 1
	//}
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
	//index = index - 1
	if index < 0 || index >= m.listMusics.Len() {
		return nil, errors.New("Index out of range.")
	}

	//for e := m.listMusics.Front(); e != nil; e = e.Next() {
	//	if index == 0 {
	//		return e.Value.(*model.MusicEntry), nil
	//	}
	//	index = index - 1
	//}

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

	//for _, music := range m.musics {
	//	if music.Name == name {
	//		return &music
	//	}
	//}

	return nil
}

func (m *MusicManagerList) Add(music *model.MusicEntry) {
	//m.musics = append(m.musics, *music)
	println("Music was added: " + music.Name)
	m.listMusics.PushFront(music)
}

func (m *MusicManager) Remove(index int) *model.MusicEntry {
	//if index < 0 || index >= len(m.musics) {
	//	return nil
	//}
	//
	//removedMusic := &m.musics[index]
	//
	//// remove from m.musics
	//if index < len(m.musics)-1 {
	//	m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	//} else if index == 0 {
	//	m.musics = []model.MusicEntry{}
	//} else {
	//	m.musics = m.musics[:index-1]
	//}
	//
	//return removedMusic
	return nil
}

func (m *MusicManager) RemoveByName(name string) *model.MusicEntry {
	//if len(m.musics) == 0 {
	//	return nil
	//}
	//
	//for i, music := range m.musics {
	//	if music.Name == name {
	//		return m.Remove(i)
	//
	//	}
	//}

	return nil
}
