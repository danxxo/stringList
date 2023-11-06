package stringList

import (
	"github.com/pkg/errors"
)

type StringList struct {
	head *node
	len  int
}

func New(s string) *StringList {

	newList := &StringList{head: nil, len: 0}

	if len(s) == 0 {
		return newList
	}
	for _, char := range s {
		newList.Append(char)
	}
	return newList
}

func (s *StringList) Len() int {
	return s.len
}

func (s *StringList) Append(c rune) {
	newNode := &node{char: c}
	s.len++
	if s.head == nil {
		s.head = newNode
		return
	}

	currNode := s.head
	for currNode.next != nil {
		currNode = currNode.next
	}

	currNode.next = newNode
}

//Добавляет символ в конец строки.

func (s *StringList) Prepend(c rune) {
	s.len++
	newNode := &node{char: c, next: s.head}
	s.head = newNode
}

func (s *StringList) Insert(c rune, index int) (err error) {
	if index > s.len || index < 0 {
		return errors.New("Out of range")
	}
	if index == s.len {
		s.Append(c)
		return nil
	}
	if index == 0 {
		s.Prepend(c)
		return nil
	}

	newNode := &node{char: c, next: nil}
	currNode := s.head

	for i := 0; i < index-1; i++ {
		currNode = currNode.next
	}

	newNode.next = currNode.next
	currNode.next = newNode
	s.len++

	return nil

}

func (s *StringList) Remove(index int) (err error) {
	if index < 0 || index > s.len-1 {
		return errors.New("Out of range")
	}

	s.len--

	if index == 0 {
		s.head = s.head.next
		return nil
	}
	currNode := s.head

	for i := 0; i < index-1; i++ {
		currNode = currNode.next
	}
	currNode.next = currNode.next.next
	return
}

//Удаляет символ из строки по индексу.

func (s *StringList) String() (str string) {
	currNode := s.head
	for currNode != nil {
		str += string(currNode.char)
		currNode = currNode.next
	}
	return str
}

/////Возвращает строковое представление списка.

func (s *StringList) At(index int) (rune, error) {
	if index > s.len-1 || index < 0 {
		return '\n', errors.New("Out of range")
	}

	currNode := s.head

	for i := 0; i < index; i++ {
		currNode = currNode.next
	}

	return currNode.char, nil
}

//Получает символ по индексу.

func (s *StringList) Concat(other *StringList) *StringList {

	// Я немного не понял, но буду делать так чтобы ни изначальная ни азер не изменялись

	newList := New(s.String())

	currOtherNode := other.head
	for currOtherNode != nil {
		newList.Append(currOtherNode.char)
		currOtherNode = currOtherNode.next
	}

	return newList
}

//Конкатенирует две строки.

func (s *StringList) Equals(other *StringList) bool {
	return s.len == other.len
}

//Проверяет, равны ли две строки.

func (s *StringList) IndexOf(c rune) int {
	index := 0
	currNode := s.head
	for currNode != nil {
		if currNode.char == c {
			return index
		}
		currNode = currNode.next
		index++
	}
	return -1
}

//Возвращает первый индекс указанного символа в строке.

func (s *StringList) Substring(start, end int) (*StringList, error) {
	if start > s.len-1 || start < 0 || end > s.len-1 || end < 0 {
		return nil, errors.New("out of range")
	}
	if start > end {
		return nil, errors.New("start > end")
	}
	if start == end {
		return New(""), nil
	}

	newList := New("")
	currNode := s.head

	for i := 0; i < start; i++ {
		currNode = currNode.next
	}

	for i := start; i <= end; i++ {
		newList.Append(currNode.char)
		currNode = currNode.next
	}
	return newList, nil
}

//Возвращает подстроку из строки.

// ЧТО ТАКОЕ n?????????????
func (s *StringList) Replace(old, new rune, n int) *StringList {
	if n < 0 {
		return New(s.String())
	}
	newList := &StringList{}
	currNode := s.head
	counter := 0

	for currNode != nil {
		if currNode.char == old && counter < n {
			newList.Append(new)
			counter++
		} else {
			newList.Append(currNode.char)
		}
		currNode = currNode.next
	}
	return newList
}

//Заменяет символы в строке.
