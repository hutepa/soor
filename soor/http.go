package soor

import (
//	"io"
	"net/http"
	"log"
	"fmt"
	"net"
	"os/exec"
	"strings"
//	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mostlygeek/arp"
//
	//        "github.com/gorilla/mux"


)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func home(res http.ResponseWriter, req *http.Request) {
	//io.WriteString(res, "default Soor Server!\n")
	//http.ServeFile(res, req, "index.html")
	if req.Method == "GET" {
		err := strings.Contains(req.UserAgent(),"CaptiveNetworkSupport")
		if err == true {
			http.Redirect(res,req, "http://10.102.0.5:8020/login",302)
		}else{
			http.Redirect(res,req, "http://10.102.0.5:8020/login",302)
		}
	}else{
		fmt.Fprintf(res, "You are not authorized to process this request")
	}
}

func login(res http.ResponseWriter, req *http.Request) {
	InitLogger()
	Trace.Println("new request")
	db, err := sql.Open("mysql", "portal:gopherT!n@@tcp(localhost:3306)/cportal?parseTime=true&loc=Asia%2FKuwait")
	if err != nil {
		Error.Printf("%v\n",err)
		//fmt.Fprintf(res, "error %q\n",err)
	}
	defer db.Close()

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		//fmt.Fprintf(res, "userip: %q is not IP:port", req.RemoteAddr)
		Error.Printf("%v\n",err)
	}
	//fmt.Fprintf(res, "hi %q\n",ip)
	mac := arp.Search(ip)
	//fmt.Fprintf(res, "hi %q\n",mac)
	var count int
	err = db.QueryRow("select count(*) as count from daily where mac=?",mac).Scan(&count)
	CheckErr(err)

	if req.Method == "GET" {
		//t, err := template.ParseFiles("index.gtpl")
		//if err != nil {
		//	fmt.Printf("%v\n",err)
		//}
		//t.Execute(res, nil)
		//if count > 0 {
		//	http.ServeFile(res, req, "logout.html")
		//}else {
		//	http.ServeFile(res, req, "index.html")
		//}
		http.ServeFile(res, req, "indexx.html")
	} else if req.Method == "POST" {
		req.ParseForm()
		// logic part of log in
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
		//fmt.Fprintf(res, "hi\n")

		if req.Form["tsantsa"] != nil {
			//fmt.Fprintf(res, "%v\n",count)
			if count > 0 {
				http.ServeFile(res, req, "logout.html")
				//fmt.Fprintf(res, "%v\n",req.Form["tsantsa"])
				return
			}
			http.ServeFile(res, req, "indexxx.html")

		}else if req.Form["phone"] != nil {
			//fmt.Fprintf(res, "%v\n",req.Form["phone"])
			//return
			fone := strings.Join(req.Form["phone"]," ")
			phone := fmt.Sprintf("%s", fone[1:])
			valid := Validate(phone)


			if valid == true {
				var dbphone *string
				var expired *string
				var verified *string
				err = db.QueryRow("select phone as dbphone,expired as expired,verified as verified from sessions where phone=? order by dtm_created desc limit 1;",phone).Scan(&dbphone,&expired,&verified)
				if err != nil {
					//log.Fatal(err)
					//fmt.Println(err)
					Error.Printf("%v\n",err)
				}
				if expired != nil || dbphone == nil {

					ipSplit := strings.Split(ip, ".")
					vlan := ipSplit[1]
					pin := Random(1000, 9999)
					pre_sess_stmt, err := db.Prepare("INSERT INTO sessions (phone,pin,ip,mac,vlan) values(?,?,?,?,?)")
					CheckErr(err)

					_, err = pre_sess_stmt.Exec(phone, pin, ip, mac, vlan)
					CheckErr(err)
					//_, err = SendSMS(phone, fmt.Sprintf("%d", pin))
					go SendSMS(phone, fmt.Sprintf("%d", pin))
					if err != nil {
						fmt.Fprintf(res, "Error: \n%v\n", err)
					} else {
						//fmt.Fprintf(res, "Response: \n%v\n", smsresp)
						http.ServeFile(res, req, "verify.html")
					}
					//http.ServeFile(res, req, "verify.html")
				}else if expired == nil && verified == nil {
					//fmt.Fprintf(res, "pincode has already sent to you, kindly verify", "")
					http.ServeFile(res, req, "verify2.html")
				}
			}else if valid == false {
				fmt.Fprintf(res, "The phone number: %v you entered is not valid mobile number\n", phone)
			}
			//fmt.Fprintf(res, "phone: %v\n", valid)
			//return
		} else if req.Form["pincode"] != nil {
			var dbpin *string
			pin := strings.Join(req.Form["pincode"]," ")
			//var expired sql.NullString
			//var verified sql.NullString
			var verified *string
			var expired *string
			var dbmac *string
			err = db.QueryRow("select pin as dbpin, mac as dbmac,verified as verified,expired as expired from sessions where pin=? order by dtm_created desc limit 1;",pin).Scan(&dbpin,&dbmac,&verified,&expired)
			//CheckErr(err)
			if err != nil {
				//log.Fatal(err)
				fmt.Println(err)
			}
			if *dbmac != mac {

				http.ServeFile(res, req, "wrongmac.html")
				return
			}

			if dbpin == nil {
				fmt.Fprintf(res, "The pin code: %v you entered is not valid\n", pin)
				return
			}else if expired == nil && verified == nil {

				//TODO
				//if verified != nil && expired != nil {
				//	//fmt.Println(id, *value)
				//	fmt.Fprintf(res, "p %v\n v %v\n e %v\n",pin,*verified,*expired)
				//} else {
				//	//fmt.Println(id, value)
				//	fmt.Fprintf(res, "p %v\n v %v\n e %v\n",pin,verified,expired)
				//}
				//fmt.Fprintf(res, "p %v\n v %v\n e %v\n",pin,verified,expired)
				//}
				//return
				//fmt.Fprintf(res, "wwww %v\n",count)
				if count == 0 {
					//fmt.Fprintf(res, "on the track\n")
					h_in_stmt, err := db.Prepare("INSERT INTO hourly (ip,mac) values(?,?)")
					CheckErr(err)

					_, err = h_in_stmt.Exec(ip, mac)
					CheckErr(err)

					d_in_stmt, err := db.Prepare("INSERT INTO daily (ip,mac) values(?,?)")
					CheckErr(err)

					_, err = d_in_stmt.Exec(ip, mac)
					CheckErr(err)

					_, err = exec.Command("iptables", "-I", "internet", "1", "-t", "mangle", "-m", "mac", "--mac-source", mac, "-j", "RETURN").Output()
					CheckErr(err)

					_, err = exec.Command("/usr/bin/rmtrack", ip).Output()
					CheckErr(err)
					req.Method = "GET"
					http.Redirect(res, req, "http://www.google.com/", 302)
					//fmt.Fprintf(res, "HTTP/1.1 302 Encryption Required\nLocation: http://google.com/")
					//fmt.Fprintf(res, "Connection: close\n")
					//fmt.Fprintf(res, "Cache-control: private\n")
					//fmt.Fprintf(res, "\n")
					sess_update_stmt, err := db.Prepare("update sessions set expired=1 , verified=1 where pin=?")
					CheckErr(err)

					_, err = sess_update_stmt.Exec(dbpin)
					CheckErr(err)

				} else {
					//fmt.Println("Existing")
					//fmt.Fprintf(res, "existing\n")
					//req.Method = "POST"
					//http.Redirect(res, req, "http://10.102.0.5:8020/logout",302)
					http.ServeFile(res, req, "logout.html")
					//fmt.Fprintf(res, "Trace \n%v\n%v\n",count,mac)
					//http.ServeFile(res, req, "logout.html")
				}

			}
		}

	}
}

func LoginServer() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/root/dev/src/soor/html"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func logout(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		//t, err := template.ParseFiles("index.gtpl")
		//if err != nil {
		//	fmt.Printf("%v\n",err)
		//}
		//t.Execute(res, nil)
		fmt.Fprintf(res, "Restricted!!!")
		//http.ServeFile(res, req, "logout.html")
	} else {
		http.ServeFile(res, req, "logout.html")
	}
}

func Vlan102Redirector(){

	
	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://10.102.0.5:8020/login", 307)
	//rh := http.RedirectHandler("http://10.102.0.1/", 307)
	mux.Handle("/", rh)

//	log.Println("Listening...")
	go http.ListenAndServe("10.102.0.5:8018", mux)

	go http.ListenAndServeTLS("10.102.0.5:8021", "server.pem", "server.key", mux)

}

func Vlan103Redirector(){

	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://10.103.0.5:8020/login", 307)
	//rh := http.RedirectHandler("http://10.103.0.1/", 307)
	mux.Handle("/", rh)

	//	log.Println("Listening...")
	go http.ListenAndServe("10.103.0.5:8018", mux)

	go http.ListenAndServeTLS("10.103.0.5:8021", "server.pem", "server.key", mux)

}

func Vlan104Redirector(){

	mux := http.NewServeMux()
	//rh := http.RedirectHandler("http://10.104.0.5:8020/login", 307)
	rh := http.RedirectHandler("http://10.104.0.1/", 307)
	mux.Handle("/", rh)

	//	log.Println("Listening...")
	go http.ListenAndServe("10.104.0.5:8018", mux)

	go http.ListenAndServeTLS("10.104.0.5:8021", "server.pem", "server.key", mux)

}


