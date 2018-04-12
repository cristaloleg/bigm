package bigm

import (
	"fmt"
	"testing"
)

func Test_modInverse(t *testing.T) {
	tests := []struct {
		a    int32
		m    int32
		want int32
	}{
		{3, 11, 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			if got := modInverse(tt.a, tt.m); got != tt.want {
				t.Errorf("modInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		x    string
		y    string
		want string
	}{
		{"34", "44", "78"},
		{"123", "321", "444"},
		{"123123123", "321321321", "444444444"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			xx := FromString(tt.x)
			yy := FromString(tt.y)
			if got := Add(xx, yy).String(); tt.want != got {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		x    string
		y    string
		want string
	}{
		{"44", "34", "10"},
		{"321", "123", "198"},
		{"321321321", "123123123", "198198198"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			xx := FromString(tt.x)
			yy := FromString(tt.y)
			if got := Sub(xx, yy).String(); tt.want != got {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		x    string
		y    string
		want string
	}{
		{"44", "34", "1496"},
		{"321", "123", "39483"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			xx := FromString(tt.x)
			yy := FromString(tt.y)
			if got := Mul(xx, yy).String(); tt.want != got {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
