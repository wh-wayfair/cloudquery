# Generating resources

Adding resources to a plugins can sometimes be a tedious task: some resources have hundreds of fields and relations, and adding them all can
take a long time. To remedy this issue, we provide code generation utilities as part of our [plugin-sdk](https://github.com/cloudquery/plugin-sdk). Code generation allows you to easily generate the boilerplate code for tables from Go code.

## Examples

The best example would be to check out how this is done in one of our official plugins, such as [gcp](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/codegen) or [aws](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/codegen).
