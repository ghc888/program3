package main

import (
	"strings"
	"fmt"

)

func IsSqlSep(r rune) bool {
	return r == ' ' || r == ',' ||
		r == '\t' || r == '/' ||
		r == '\n' || r == '\r'
}
var (
	TK_ID_INSERT   = 1
	TK_ID_UPDATE   = 2
	TK_ID_DELETE   = 3
	TK_ID_REPLACE  = 4
	TK_ID_SET      = 5
	TK_ID_BEGIN    = 6
	TK_ID_COMMIT   = 7
	TK_ID_ROLLBACK = 8
	TK_ID_ADMIN    = 9
	TK_ID_USE      = 10

	TK_ID_SELECT      = 11
	TK_ID_START       = 12
	TK_ID_TRANSACTION = 13
	TK_ID_SHOW        = 14
	TK_ID_TRUNCATE    = 15


PARSE_TOKEN_MAP = map[string]int{
"insert":      TK_ID_INSERT,
"update":      TK_ID_UPDATE,
"delete":      TK_ID_DELETE,
"replace":     TK_ID_REPLACE,
"set":         TK_ID_SET,
"begin":       TK_ID_BEGIN,
"commit":      TK_ID_COMMIT,
"rollback":    TK_ID_ROLLBACK,
"admin":       TK_ID_ADMIN,
"select":      TK_ID_SELECT,
"use":         TK_ID_USE,
"start":       TK_ID_START,
"transaction": TK_ID_TRANSACTION,
"show":        TK_ID_SHOW,
"truncate":    TK_ID_TRUNCATE,
}

	TK_STR_SELECT = "select"
	TK_STR_FROM   = "from"
	TK_STR_INTO   = "into"
	TK_STR_SET    = "set"


)

func main()  {
	sql:="SELECT imei,mid from device where imei='a' and uptime between '2017-09-06 00:00:00' and '2017-09-06 23:59:59' order by id limit 0,10"
	tokens := strings.FieldsFunc(sql, IsSqlSep)

	tokensLen := len(tokens)

	tokenId:=PARSE_TOKEN_MAP[strings.ToLower(tokens[0])]
	fmt.Println("tokensLen:",tokensLen)
	fmt.Println("tokenId:",tokenId)
	fmt.Println(tokens)




}