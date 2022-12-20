package paginator

import (
	"testing"

	_ "github.com/MortezaRastegarRad/themes/sword"
	"github.com/MortezaRastegarRad/go-admin/modules/config"
	"github.com/MortezaRastegarRad/go-admin/plugins/admin/modules/parameter"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
