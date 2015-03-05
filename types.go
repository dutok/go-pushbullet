package main

type Target struct {
    DeviceIden string `json:"device_iden,omitempty"`
    Email string `json:"email,omitempty"`
    ChannelTag string `json:"channel_tag,omitempty"`
    ClientIden string `json:"client_iden,omitempty"`
}

type PushReq struct {
    Type string  `json:"type"`
    Target
    SourceIden string  `json:"source_device_iden"`
}

type Note struct {
    PushReq
    Title string  `json:"title"`
    Body  string  `json:"body"`
}

type Link struct {
    PushReq
    Title string  `json:"title"`
    Body  string  `json:"body"`
    Url   string  `json:"url"`
}

type File struct {
    PushReq
    FileName  string  `json:"file_name"`
    FileType  string  `json:"file_type"`
    FileUrl   string  `json:"file_url"`
    Body      string  `json:"body"`
}

type List struct {
    PushReq
    Title   string  `json:"title"`
    Items   []string  `json:"items"`
}

func loadTarget(targettype, targetvalue string) (Target) {
    target := Target{}
    switch targettype {
		case "device_iden":
			target = Target {targetvalue, "", "", ""}
		case "email":
			target = Target {"", targetvalue, "", ""}
		case "channel_tag":
			target = Target {"", "", targetvalue, ""}
		case "client_iden":
			target = Target {"", "", "", targetvalue}
	}
	return target
}