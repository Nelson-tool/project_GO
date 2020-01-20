package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Block struct {
	Timestamp int
	Data	[]byte
	PrevBlock []byte
	Hash	[]byte
}