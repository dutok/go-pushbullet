package main

import (
    "os"
    "fmt"
    "encoding/json"
    "strconv"
)

type Push struct {
	Active                  bool    `json:"active"`
	Body                    string  `json:"body"`
	Created                 float64 `json:"created"`
	Dismissed               bool    `json:"dismissed"`
	Iden                    string  `json:"iden"`
	Modified                float64 `json:"modified"`
	ReceiverEmail           string  `json:"receiver_email"`
	ReceiverEmailNormalized string  `json:"receiver_email_normalized"`
	ReceiverIden            string  `json:"receiver_iden"`
	SenderEmail             string  `json:"sender_email"`
	SenderEmailNormalized   string  `json:"sender_email_normalized"`
	SenderIden              string  `json:"sender_iden"`
	Title                   string  `json:"title"`
	Type                    string  `json:"type"`
}

type Pushes struct {
	Pushes []Push `json:"pushes"`
}


func main() {
    key = os.Getenv("pbkey")   
    push, err := newNote("Test", "Test body", "channel_tag", "go", "")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(push)
    }
    pushes, err := getPushes(1, true, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(pushes)
    }
}

func getPushes(after int, active bool, cursor int) (Pushes, error) {
    pushes := Pushes{}
    url := "pushes?modified_after=" + strconv.Itoa(after)
    if active { url += "&active=true" }
    if cursor != 0 { url += "&cursor=" + strconv.Itoa(cursor) }
    respbytes, err := request("GET", url, "")
    if err != nil {
        return pushes, err
    }
    err = json.Unmarshal(respbytes, &pushes)
    if err != nil {
        fmt.Println(err)
    }
    return pushes, nil
}

func newPush(pushtype, title, url, filename, filetype, fileurl, body, targettype, target, sourceiden string) (Push, error) {
    push := Push{}
    jsonStr := ""
    if pushtype == "note" {
        jsonStr = `{"type": "note", "title": "`+ title +`", "body": "`+ body +`", "`+ targettype +`": "`+ target +`", "source_device_iden": "` + sourceiden + `"}`
    } else if pushtype == "link" {
        jsonStr = `{"type": "link", "title": "`+ title +`", "body": "`+ body +`", "url": "`+ url +`", "`+ targettype +`": "`+ target +`", "source_device_iden": "` + sourceiden + `"}`
    } else {
        jsonStr = `{"type": "file", "file_name": "`+ filename +`", "file_type": "`+ filetype +`", "file_url": "`+ fileurl +`", "body": "`+ body +`", "`+ targettype +`": "`+ target +`", "source_device_iden": "` + sourceiden + `"}`
    }
    respbytes, err := request("POST", "pushes", jsonStr)
    if err != nil {
        return push, err
    }
    err = json.Unmarshal(respbytes, &push)
    if err != nil {
        fmt.Println(err)
    }
    return push, nil
}

func newNote(title, body, targettype, target, sourceiden string) (Push, error) {
    push, err := newPush("note", title, "", "", "", "", body, targettype, target, sourceiden)
    if err != nil {
        return push, err
    }
    return push, nil
}

func newLink(title, body, url, targettype, target, sourceiden string) (Push, error) {
    push, err := newPush("link", title, url, "", "", "", body, targettype, target, sourceiden)
    if err != nil {
        return push, err
    }
    return push, nil
}

func newFile(title, filename, filetype, fileurl, body, targettype, target, sourceiden string) (Push, error) {
    push, err := newPush("file", title, "", filename, filetype, fileurl, body, targettype, target, sourceiden)
    if err != nil {
        return push, err
    }
    return push, nil
}