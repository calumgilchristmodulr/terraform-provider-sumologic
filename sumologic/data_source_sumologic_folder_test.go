package sumologic

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourcSumologicFolder_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceAccSumologicFolderConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceFolderCheck("data.sumologic_folder.by_name", "sumologic_folder.test"),
					testAccDataSourceFolderCheck("data.sumologic_folder.by_id", "sumologic_folder.test"),
				),
			},
		},
	})
}

func testAccDataSourceFolderCheck(name, reference string) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(name, "id"),
		resource.TestCheckResourceAttrPair(name, "id", reference, "id"),
		resource.TestCheckResourceAttrPair(name, "name", reference, "name"),
		resource.TestCheckResourceAttrPair(name, "description", reference, "description"),
		resource.TestCheckResourceAttrPair(name, "category", reference, "category"),
		resource.TestCheckResourceAttrPair(name, "timezone", reference, "timezone"),
	)
}

var testDataSourceAccSumologicFolderConfig = `
resource "sumologic_folder" "test" {
  name = "MyFolder"
  description = "MyFolderDesc"
}

data "sumologic_collector" "by_name" {
  name = "${sumologic_folder.test.name}32"
}

data "sumologic_collector" "by_id" {
  id = "${sumologic_folder.test.id}"
}
`
