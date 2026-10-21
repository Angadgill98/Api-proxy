package global

import "database/sql"

//here both side is the remote addr
var Active_listeners=map[string]string{

}


var Db *sql.DB