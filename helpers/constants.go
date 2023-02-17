package helpers

// User Types

var USER_TYPE_NORMAL string = "NORMAL"
var USER_TYPE_ADMIN string = "ADMIN"
var USER_TYPE_MODERATOR string = "MODERATOR"

var TICKET_STATUS_OPEN string = "OPEN"
var TICKET_STATUS_CLOSED string = "CLOSED"
var TICKET_STATUS_WONT_DO string = "WONT_DO"

func AllUserTypes() []string {
	return []string{USER_TYPE_ADMIN, USER_TYPE_MODERATOR, USER_TYPE_NORMAL}
}

func IsValidUserType(query string) bool {
	for _, usertype := range AllUserTypes() {
		if usertype == query {
			return true
		}
	}
	return false
}

func TicketStatuses() []string {
	return []string{TICKET_STATUS_CLOSED, TICKET_STATUS_WONT_DO}
}

func IsValidTicketStatus(query string) bool {
	for _, usertype := range TicketStatuses() {
		if usertype == query {
			return true
		}
	}
	return false
}
