package books

import (
 )
//Book struct,have the information of a book
type Book struct{
	Title string
	Author string
	Pages int
}
//CategoryByLength method,split novel which have pages more tan 300,or short story which have pages less than 300
func(b Book)CategoryByLength()(title string){
	if b.Pages>300 {title = "NOVEL"}
	if b.Pages<300 {title = "SHORT STORY"}
	return  title
}
//AuthorLastName  method,pick the last name by choice letters after space
func(b Book)AuthorLastName()(name string){
	for i,x:= range b.Author{
		if string(x)==" "{
			name = b.Author[(i+1):]
		}
	}
	return  name
}
////NewBookFromJSON method,Unmarshal json and fill the book infomation
//func NewBookFromJSON(info string) (myBook Book,err error) {
//
//	err = json.Unmarshal([]byte(info), &myBook)
//	if (err != nil) {
//		log.Printf("Error -> %v", err.Error())
//	}
//	return myBook,err
//}
