package soor


import (
	"regexp"
	"math/rand"
	"time"
	"fmt"
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	file, err := os.OpenFile("/var/log/cportal/cportal.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", file, ":", err)
	}
	Trace = log.New(file,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(file,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(file,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func Validate(p string) bool {

	//phone, err := strconv.Atoi(p)
	//if err != nil {
	//	return false
	//}
	ccode := fmt.Sprintf("%s", p[:3])
	if ccode == "965" {

		phone := []byte(p)
		re := regexp.MustCompile("^965[965][0-9]{7}$")
		matched := re.Match(phone)
		if matched == true {
			return true
		} else {
			return false
		}
	}else{
		return true
	}

}





