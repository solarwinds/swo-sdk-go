// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DboDatabasePluginConfigResponse struct {
	// Configuration of plugin(s) observing database server
	PluginConfig []DboDatabasePluginConfig `json:"pluginConfig"`
}

func (o *DboDatabasePluginConfigResponse) GetPluginConfig() []DboDatabasePluginConfig {
	if o == nil {
		return []DboDatabasePluginConfig{}
	}
	return o.PluginConfig
}
