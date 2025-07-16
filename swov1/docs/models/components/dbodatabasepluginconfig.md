# DboDatabasePluginConfig


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `PluginName`                                                                     | *string*                                                                         | :heavy_check_mark:                                                               | Name of plugin observing database server                                         |
| `ConfigOptions`                                                                  | [][components.CommonKeyValuePair](../../models/components/commonkeyvaluepair.md) | :heavy_check_mark:                                                               | Configuration of plugin observing database server                                |
| `DbConnOptions`                                                                  | [][components.CommonKeyValuePair](../../models/components/commonkeyvaluepair.md) | :heavy_check_mark:                                                               | Database connection options of plugin observing database server                  |