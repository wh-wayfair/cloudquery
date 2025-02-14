# Table: stripe_payment_links

This table shows data for Stripe Payment Links.

https://stripe.com/docs/api/payment_links

The primary key for this table is **id**.

## Relations

The following tables depend on stripe_payment_links:
  - [stripe_payment_link_line_items](stripe_payment_link_line_items)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|active|Bool|
|after_completion|JSON|
|allow_promotion_codes|Bool|
|application_fee_amount|Int|
|application_fee_percent|Float|
|automatic_tax|JSON|
|billing_address_collection|String|
|consent_collection|JSON|
|currency|String|
|customer_creation|String|
|custom_fields|JSON|
|custom_text|JSON|
|invoice_creation|JSON|
|line_items|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|on_behalf_of|JSON|
|payment_intent_data|JSON|
|payment_method_collection|String|
|payment_method_types|StringArray|
|phone_number_collection|JSON|
|shipping_address_collection|JSON|
|shipping_options|JSON|
|submit_type|String|
|subscription_data|JSON|
|tax_id_collection|JSON|
|transfer_data|JSON|
|url|String|