package main

import (
           "bytes"
           "context"
           "fmt"
           "net/http"
           "os"
           "os/signal"
           "syscall"

           -"github.com/mattn//go-sqlite3"
           "github.com/mdp/qrterminal/v3"
           "go.mau.fi/whatsmeow"
           waProto "go.mau.fi/whatsmeowbinary/protp"
           "go.mau.fi/whatsmeow/store/sqlstore"
           "go.mau.fi/whatsmeow/types"
           "go.mau.fi/whatsmeow/types/events"
           waLog "go.mau.fi/whatsmeow/util/log"
           "google.golang.org/protobuf/proto"
)