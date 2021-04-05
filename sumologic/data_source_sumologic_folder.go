package sumologic

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSumologicFolder() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSumologicFolderRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceSumologicFolderRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	var folder *Folder
	var err error

	if fid, ok := d.GetOk("id"); ok {
		id := fid.(string)
		folder, err = c.GetFolder(id)
		if err != nil {
			return fmt.Errorf("folder with id %s not found: %v", id, err)
		}
	} else {
		if cname, ok := d.GetOk("name"); ok {
			name := cname.(string)
			// TODO: Define a TimeoutGetFolder?
			folder, err = c.GetGlobalFolder(name, d.Timeout(schema.TimeoutDelete))
			if err != nil {
				return fmt.Errorf("folder with name %s not found: %v", name, err)
			}
			if folder == nil {
				return fmt.Errorf("folder with name %s not found", name)
			}
		} else {
			return errors.New("please specify either id or name")
		}
	}

	if err != nil {
		return err
	}

	d.SetId(folder.ID)
	d.Set("name", folder.Name)
	d.Set("description", folder.Description)

	return nil
}
