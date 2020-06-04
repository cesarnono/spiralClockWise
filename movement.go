package main

type movement interface {
	execute([][]int, int, int, int) (int, int, int)
	setNext(movement)
}

type movementRight struct {
	next movement
}

func (m *movementRight) execute(s [][]int, counter int, row int, col int) (int, int, int) {
	for col < len(s) {
		if s[row][col] > 0 && s[row][col] != counter {
			break
		}
		s[row][col] = counter
		counter++
		col++
	}
	if counter > (len(s) * len(s)) {
		return counter, col, row
	}
	counter--
	col--
	counter, col, row = m.next.execute(s, counter, row, col)
	return counter, col, row
}
func (m *movementRight) setNext(next movement) {
	m.next = next
}

type movementDown struct {
	next movement
}

func (m *movementDown) execute(s [][]int, counter int, row int, col int) (int, int, int) {
	for row < len(s) {
		if s[row][col] > 0 && s[row][col] != counter {
			break
		}
		s[row][col] = counter
		counter++
		row++
	}
	if counter > (len(s) * len(s)) {
		return counter, col, row
	}
	row--
	counter--
	counter, col, row = m.next.execute(s, counter, row, col)
	return counter, col, row
}
func (m *movementDown) setNext(next movement) {
	m.next = next
}

type movementLeft struct {
	next movement
}

func (m *movementLeft) execute(s [][]int, counter int, row int, col int) (int, int, int) {
	for col >= 0 {
		if s[row][col] > 0 && s[row][col] != counter {
			break
		}
		s[row][col] = counter
		counter++
		col--
	}
	if counter > (len(s) * len(s)) {
		return counter, col, row
	}
	col++
	counter--
	counter, col, row = m.next.execute(s, counter, row, col)
	return counter, col, row
}
func (m *movementLeft) setNext(next movement) {
	m.next = next
}

type movementUp struct {
	next movement
}

func (m *movementUp) execute(s [][]int, counter int, row int, col int) (int, int, int) {
	for row >= 0 {
		if s[row][col] > 0 && s[row][col] != counter {
			break
		}
		s[row][col] = counter
		counter++
		row--
	}
	if counter > (len(s) * len(s)) {
		return counter, col, row
	}
	row++
	counter--
	counter, col, row = m.next.execute(s, counter, row, col)
	return counter, col, row
}
func (m *movementUp) setNext(next movement) {
	m.next = next
}
