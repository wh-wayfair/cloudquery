package client

import "github.com/cloudquery/plugin-sdk/cqtypes"

func (c *Client) TransformBool(v *cqtypes.Bool) interface{} {
	return v.String()
}

func (c *Client) TransformBytea(v *cqtypes.Bytea) interface{} {
	return v.String()
}

func (c *Client) TransformFloat8(v *cqtypes.Float8) interface{} {
	return v.String()
}

func (c *Client) TransformInt8(v *cqtypes.Int8) interface{} {
	return v.String()
}

func (c *Client) TransformInt8Array(v *cqtypes.Int8Array) interface{} {
	return v.String()
}

func (c *Client) TransformJSON(v *cqtypes.JSON) interface{} {
	return v.String()
}

func (c *Client) TransformText(v *cqtypes.Text) interface{} {
	return v.String()
}

func (c *Client) TransformTextArray(v *cqtypes.TextArray) interface{} {
	return v.String()
}

func (c *Client) TransformTimestamptz(v *cqtypes.Timestamptz) interface{} {
	return v.String()
}

func (c *Client) TransformUUID(v *cqtypes.UUID) interface{} {
	return v.String()
}

func (c *Client) TransformUUIDArray(v *cqtypes.UUIDArray) interface{} {
	return v.String()
}

func (c *Client) TransformCIDR(v *cqtypes.CIDR) interface{} {
	return v.String()
}

func (c *Client) TransformCIDRArray(v *cqtypes.CIDRArray) interface{} {
	return v.String()
}

func (c *Client) TransformInet(v *cqtypes.Inet) interface{} {
	return v.String()
}

func (c *Client) TransformInetArray(v *cqtypes.InetArray) interface{} {
	return v.String()
}

func (c *Client) TransformMacaddr(v *cqtypes.Macaddr) interface{} {
	return v.String()
}

func (c *Client) TransformMacaddrArray(v *cqtypes.MacaddrArray) interface{} {
	return v.String()
}
