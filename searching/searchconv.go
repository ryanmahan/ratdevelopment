package searching

import (
	"fmt"
	"strings"
)

//searchquery: "Longneck, 99967"
//["Longneck" " 99967"]
//["Longneck" "99967"]
//["company_name_index LIKE \"%Longneck%\"" "ser_string_index LIKE \"%99967%\""]
//"company_name_index LIKE \"%Longneck%\" AND ser_string_index LIKE \"%99967%\"";

//SearchQueryToCQL takes a search string and creates a CQL query
func SearchQueryToCQL(query string) string {
	if query == "" {
		return "SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ?"
	}
	queries := strings.Split(query, ",")
	for i, q := range queries {
		queries[i] = strings.Trim(q, " ")
		if isNum(queries[i][0]) { // is a sernum
			queries[i] = fmt.Sprintf("ser_string_index LIKE \"%%%v%%\"", queries[i])
		} else { // is a company name
			queries[i] = fmt.Sprintf("company_name_index LIKE \"%%%v%%\"", queries[i])
		}
	}
	addend := strings.Join(queries, " AND ")
	return fmt.Sprintf("SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ? AND %v", addend)
}

func isNum(r byte) bool {
	return '0' <= r && r <= '9'
}
