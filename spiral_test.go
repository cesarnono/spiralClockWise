package main

import (
	"fmt"
	"testing"
)

func TestSpiralParameterNoNumeric(t *testing.T) {
	got := createSpiral("rt")
	want := make([][]int, 0)
	fmt.Println("want", want)
	fmt.Println("got", got)
	if len(got) != len(want) {
		t.Errorf("createSpiral('rt') == %d, want %d", got, want)
	}
}

func TestSpiralParameterLeastThanOne(t *testing.T) {
	got := createSpiral("0")
	want := make([]int, 0)
	if len(got) != len(want) {
		t.Errorf("createSpiral('rt') == %d, want %d", got, want)
	}
}

func TestSpiralParameterEqualOne(t *testing.T) {
	got := createSpiral("1")
	want := make([][]int, 1)
	want[0] = append(want[0], 1)
	if len(got) != len(want) {
		t.Errorf("TestSpiralParameterEqualOne::size == %d, want %d", got, want)
	}
	if got[0][0] != want[0][0] {
		t.Errorf("TestSpiralParameterEqualOne::value == %d, want %d", got, want)
	}
}

func TestSpiralParameterEqualTwo(t *testing.T) {
	got := createSpiral("2")
	want := make([][]int, 2)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			want[i] = append(want[i], 0)
		}
	}
	want[0][0] = 1
	want[0][1] = 2
	want[1][0] = 4
	want[1][1] = 3
	if len(got) != len(want) {
		t.Errorf("TestSpiralParameterEqualOne::size == %d, want %d", got, want)
	}
	if got[0][0] != want[0][0] || got[0][1] != want[0][1] || got[1][0] != want[1][0] || got[1][1] != want[1][1] {
		t.Errorf("TestSpiralParameterEqualOne::value == %d, want %d", got, want)
	}
}

func TestSpiralParameterEqualThree(t *testing.T) {
	got := createSpiral("3")
	fmt.Println("got: ", got)
	want := make([][]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			want[i] = append(want[i], 0)
		}
	}
	want[0][0] = 1
	want[0][1] = 2
	want[0][2] = 3
	want[1][2] = 4
	want[2][2] = 5
	want[2][1] = 6
	want[2][0] = 7
	want[1][0] = 8
	want[1][1] = 9
	if len(got) != len(want) {
		t.Errorf("got::size == %d, want %d", len(got), len(want))
	}
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(want); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("got::value == %d, want %d", got, want)
				i = len(got)
				break
			}
		}
	}
}
