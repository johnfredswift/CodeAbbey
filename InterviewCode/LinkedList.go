package InterviewCode

import "fmt"

func LinkedList() {
	listElement1 := &listElement{1,"1", 1.0, nil}
	elementList := listElement1

	listElement2 := &listElement{2,"2", 2.0, nil}
	listElement3 := &listElement{3,"3", 3.0, nil}

	elementList = addEndElement(listElement2, elementList)
	elementList = addEndElement(listElement3, elementList)
	printList(elementList)
}

type listElement struct {
	value1 int
	value2 string
	value3 float64
	nextElement *listElement
}
func addEndElement(newElement, currentList *listElement) *listElement {
	if currentList == nil { return currentList }
	for e := currentList; e != nil; e = e.nextElement{
		if e.nextElement == nil {
			e.nextElement = newElement
			return currentList
		}
	}
	return currentList
}
func printList(elementList *listElement){
	for e := elementList; e != nil; e = e.nextElement{ fmt.Println(e)}
}
func deleteElement(p *listElement){

}
