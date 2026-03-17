package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!!" {
		t.Fatal("Test fail")
	}
}

func TestHW2(t *testing.T) {
	if helloworld() != "Hello World!!" {
		t.Fatal("Test fail")
	}
}

func TestNihao(t *testing.T) {
	if nihao() != "你好" {
		t.Fatal("Test fail")
	}
}
