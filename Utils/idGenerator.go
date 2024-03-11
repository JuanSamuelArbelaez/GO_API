package Utils

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/SQL"
	"strconv"
)

func GenerateId(id *string) error {
	i, e := SQL.CountProducts()
	if e != nil {
		return e
	}
	*id = formatId(i + 1)
	return nil
}

func formatId(number int) string {
	hexStr := strconv.FormatInt(int64(number), 16)
	formattedNumber := fmt.Sprintf("%04s", hexStr)
	return formattedNumber
}
