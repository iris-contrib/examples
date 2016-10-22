package models

type Error struct {
	Response bool   `json:"response" bson:"response"`
	Error    string `json:"error" bson:"error"`
	ErrorNo  uint8  `json:"errorno" bson:"errorno"`
}

func Err(errIn string) Error {

	switch errIn {
	case "1", "Results":
		return Error{Error: "No Results.", ErrorNo: 1}
		break
	case "2", "Request":
		return Error{Error: "Request Error.", ErrorNo: 2}
		break
	case "3", "Delete":
		return Error{Error: "Can't delete document.", ErrorNo: 3}
		break
	case "4", "Formdata":
		return Error{Error: "Incomplete form data or wrong params.", ErrorNo: 4}
		break
	case "5", "Database":
		return Error{Error: "Database Error.", ErrorNo: 5}
		break
	case "6", "Duplicated":
		return Error{Error: "Duplicated field in Database.", ErrorNo: 6}
		break
	case "7", "Login":
		return Error{Error: "Login Error: wrong user or password.", ErrorNo: 7}
		break
	case "8", "Session":
		return Error{Error: "Session not match.", ErrorNo: 8}
		break

	default:
		return Error{Error: "Error.", ErrorNo: 0}
		break
	}

	return Error{Response: false}

}
