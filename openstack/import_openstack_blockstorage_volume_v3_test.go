package openstack

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccBlockStorageV3Volume_importBasic(t *testing.T) {
	resourceName := "openstack_blockstorage_volume_v3.volume_1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckNonAdminOnly(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBlockStorageV3VolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageV3VolumeBasic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
