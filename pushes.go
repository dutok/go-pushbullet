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
    /*push, err := pushList("Test list", []string{"apple", "peach", "pear"}, "", "", "")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(push)
    }
    push, err = pushLink("Test link", "Test link body", "http://google.com", "", "", "")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(push)
    }
    push, err = pushNote("Test note", "Test note body", "", "", "")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(push)
    }*/
    pushes, err := getPushes(0, true, 0)
    if err != nil {
        fmt.Println(err)
    }
    pushes.delete()
}

func (push *Push) delete() (*Push, error) {
    url := "pushes"
    url += "/" + push.Iden
    respbytes, err := request("DELETE", url, "")
    if err != nil {
        return push, err
    }
    if string(respbytes) == "{}" {
        *push = Push{}
    }
    return push, nil
}

func (pushes *Pushes) delete() (*Pushes, error) {
    respbytes, err := request("DELETE", "pushes", "")
    if err != nil {
        return pushes, err
    }
    if string(respbytes) == "{}" {
        *pushes = Pushes{}
    }
    return pushes, nil
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

func newPush(jsonBytes []byte) (Push, error) {
    push := Push{}
    respbytes, err := request("POST", "pushes", string(jsonBytes))
    if err != nil {
        return push, err
    }
    err = json.Unmarshal(respbytes, &push)
    if err != nil {
        fmt.Println(err)
    }
    return push, nil
}

func pushNote(title, body, targettype, targetvalue, sourceiden string) (Push, error) {
    target := loadTarget(targettype, targetvalue)
    req := PushReq{
        Type: "note",
        Target: target,
        SourceIden: sourceiden,
    }
    note := Note{
        PushReq: req,
        Title: title,
        Body: body,
    }
    json, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
	}
    push, err := newPush(json)
    if err != nil {
        return push, err
    }
    return push, nil
}

func pushLink(title, body, url, targettype, targetvalue, sourceiden string) (Push, error) {
    target := loadTarget(targettype, targetvalue)
    req := PushReq{
        Type: "link",
        Target: target,
        SourceIden: sourceiden,
    }
    link := Link{
        PushReq: req,
        Title: title,
        Body: body,
        Url: url,
    }
    json, err := json.Marshal(link)
	if err != nil {
		fmt.Println(err)
	}
    push, err := newPush(json)
    if err != nil {
        return push, err
    }
    return push, nil
}

func pushFile(filename, filetype, fileurl, body, targettype, targetvalue, sourceiden string) (Push, error) {
    target := loadTarget(targettype, targetvalue)
    req := PushReq{
        Type: "file",
        Target: target,
        SourceIden: sourceiden,
    }
    file := File{
        PushReq: req,
        FileName: filename,
        FileType: filetype,
        FileUrl: fileurl,
        Body: body,
    }
    json, err := json.Marshal(file)
	if err != nil {
		fmt.Println(err)
	}
    push, err := newPush(json)
    if err != nil {
        return push, err
    }
    return push, nil
}

func pushList(title string, items []string, targettype, targetvalue, sourceiden string) (Push, error) {
    target := loadTarget(targettype, targetvalue)
    req := PushReq{
        Type: "list",
        Target: target,
        SourceIden: sourceiden,
    }
    list := List{
        PushReq: req,
        Title: title,
        Items: items,
    }
    json, err := json.Marshal(list)
    fmt.Println(string(json))
	if err != nil {
		fmt.Println(err)
	}
    push, err := newPush(json)
    if err != nil {
        return push, err
    }
    return push, nil
}