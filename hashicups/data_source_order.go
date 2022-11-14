package hashicups

import (
	"context"
	"strconv"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrderRead,
		Schema: map[string]*schema.Schema{ // Schemas are a map with string keys for object properties and schema.Schema for
			"id": &schema.Schema{ // values.
				Type:     schema.TypeInt, // Schema type allows you to specify a value type, wheter the property is required
				Required: true,           // and if it should be provided or fetched from the resource api (computed)
			},
			"items": &schema.Schema{
				Type:     schema.TypeList, // Schema lists can hold multiple elements of a given resource type
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_teaser": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_price": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_image": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// This "meta" call is convention based and I don't have any idea how I'm supposed to know what anything is
	c := m.(*hc.Client)
	// Diagnostics is a slice of err to be collected and reported back to the user
	var diags diag.Diagnostics

	orderID := strconv.Itoa(d.Get("id").(int))

	order, err := c.GetOrder(orderID)
	if err != nil {
		return diag.FromErr(err)
	}

	orderItems := flattenOrderItemsData(&order.Items)
	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orderID)

	return diags
}

func flattenOrderItemsData(orderItems *[]hc.OrderItem) []interface{} {
	if orderItems != nil {
		ois := make([]interface{}, len(*orderItems), len(*orderItems))

		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})

			oi["coffee_id"] = orderItem.Coffee.ID
			oi["coffee_name"] = orderItem.Coffee.Name
			oi["coffee_teaser"] = orderItem.Coffee.Teaser
			oi["coffee_description"] = orderItem.Coffee.Description
			oi["coffee_price"] = orderItem.Coffee.Price
			oi["coffee_image"] = orderItem.Coffee.Image
			oi["quantity"] = orderItem.Quantity

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}
