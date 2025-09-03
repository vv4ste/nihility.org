package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	//"github.com/niuhuan/nhentai-go"
)

type translation struct {
	From	string
	Title	string
	NhId	int
	NhLikes int
	EhId	string
	Date	string
}

type staff struct {
	Avatar	string
	Nick	string
	Roles	[]string
}

type dto struct {
	Staff			[]staff
	Translations	[]translation
}

var translations = []translation {
	{
		From: "[Chirimen Naoyuki (Naoyuki)] (Princess Connect! Re:Dive)",
		Title: "Kyaru-chan is wild in the mating season!?",
		NhId: 546037,
		EhId: "3171969/245a553fd3",
		Date: "2024-12-24",
	},
	{
		From: "[Kadutikiya (Kaduki)] (Genshin Impact)",
		Title: "To the inexperienced 6000 year old you",
		NhId: 530496,
		EhId: "3059227/88342cd22b",
		Date: "2024-09-15",
	},
};

var vstaff = []staff {
	{
		Nick: ".asapgiri",
		Avatar: "asa.png",
		Roles: []string{"translator", "proofreader"},
	},
	{
		Nick: "vv4ste",
		Avatar: "wxxstd.png",
		Roles: []string{"editor", "typesetter"},
	},
	{
		Nick: "rd.szili",
		Avatar: "rd.szili.png",
		Roles: []string{"translator", "proofreader"},
	},
}



func Unexpected(w http.ResponseWriter, r *http.Request) {
    fil, typ := read_artifact(r.URL.Path, w.Header())

    if "text" == typ {
        Render(w, fil, nil)
    } else {
        io.WriteString(w, fil)
    }
}

func collect_translations() []translation {
	// var client = nhentai.Client{}
	//
	// client.Transport = &http.Transport{
	// 	TLSHandshakeTimeout:   time.Second * 10,
	// 	ExpectContinueTimeout: time.Second * 10,
	// 	ResponseHeaderTimeout: time.Second * 10,
	// 	IdleConnTimeout:       time.Second * 10,
	// }

	for i := 0; i < len(translations); i++ {
		// info, err := client.ComicInfo(translations[i].NhId)
		// fmt.Println(info)
		// fmt.Println(err)
		translations[i].NhLikes = 4000 //info.NumFavorites
	}

	return translations
}

func Root(w http.ResponseWriter, r *http.Request) {
    if "/" == r.URL.Path {
		// TODO: Read nh stuff...

		dto_tr := dto{
			Staff: vstaff,
			Translations: collect_translations(),
		}

        fil, _ := read_artifact("index.html", w.Header())
        Render(w, fil, dto_tr)
    } else {
        Unexpected(w, r)
    }
}

func main() {
    InitConfig()

    http.HandleFunc("GET /",                    Root)
    http.HandleFunc("GET /index",               Root)
    http.HandleFunc("GET /index.html",          Root)

    args := os.Args[1:]
    if 0 < len(args) {
        Config.Http.Port = args[0];
    }

    http.ListenAndServe(strings.Join([]string{":", Config.Http.Port}, ""), nil)
}
