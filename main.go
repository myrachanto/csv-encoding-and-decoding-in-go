package main

import (
	"encoding/csv"
	"encoding/xml"
	"encoding/json"
	"strconv"
	
	"fmt"
	"io"
	"log"
	"os"
)
type(
	User struct {
		Id int `json:"id" xml:"id,attr"`
		Firstname string `json:"firstname" xml:"name>first"`
		Lastname string `json:"lastname" xml:"name>last"`
		Username string `json:"username,omitempty" xml:"secret>username"`
		Password string `json:"password,omitempty" xml:"password>secret"`
		Email string `json:"email" xml:"email"`
	}
	UserDb struct {
		XMLName xml.Name `json:"-" xml:"users"`
		Type string `json:"type,omitempty" xml:"type"`
		Users []User `json:"users,omitempty" xml:"user"`

	}
)
func main(){
	db,err := readJsonFile("../user.db.json")
	Check(err)
	header := []string{`1`, `anthony`, `miracho`, `myrachanto`, `password`, `myrachanto@gmail.com`}
	//user1 := []string{`2`, `anthony`, `myracho`, `myrachanto`, `password`, `info@gmail.com`}
	//user2 := []string{`3`, `tony`, `miracho`, `myrachanto`, `password`, `chantos@gmail.com`}
	//user3 := []string{`4`, `antonio`, `myracho`, `myrachanto`, `password`, `myrachanto@gmail.com`}
	f, err := os.Create("csv_data1.csv")
	Check(err)

	defer f.Close()
	
	w := csv.NewWriter(f)
	//w.Write(user1)
	/*data := [][]string{
		header,user1,user2,user3,
	}*/
	w.Write(header)
	for _,user := range db.Users{
		ss := user.EncodeAsStrings()
		w.Write(ss)
  
	}
	w.Flush()
	Check(err)
	//from csv
	g, err := os.Open("csv_data1.csv")
	Check(err)

	defer g.Close()
	r := csv.NewReader(g)
	for{
		csvRecord, err := r.Read()
		if err != nil {
			Process(csvRecord)
		} else if io.EOF == err {
			break
		}else {
			log.Fatalln(err)
		}
	}

	
}
func readJsonFile(s string)(db *UserDb, err error) {
	f, err := os.Open(s)
	Check(err)
	defer f.Close()
	var d = json.NewDecoder(f)
	db = new(UserDb)
	d.Decode(db)
	return
}
func (user User) EncodeAsStrings() (ss []string){
	ss = make([]string, 6)
	ss[0] =  strconv.Itoa(user.Id)
	ss[1] = user.Firstname
	ss[2] = user.Lastname
	ss[3] = user.Username
	ss[4] = user.Password
	ss[5] = user.Email
 return
}

func CsvDecoder (){

}
func Process(ss []string){
	u := &User{}
	u.FromCSv(ss)
	fmt.Println(u.Firstname, u.Lastname, u.Username, u.Password, u.Email)

}
func (user *User) FromCSv(ss []string){
	if user == nil {
		return
	}
	if   nil == ss {
		return
	}
	user.Id,_ = strconv.Atoi(ss[0])
	user.Firstname = ss[1]
	user.Lastname = ss[2]
	user.Username = ss[3]
	user.Password = ss[4]
	user.Email = ss[5]
}
func Check(err error){
	if err != nil {
		log.Fatalln(err)
	}
}