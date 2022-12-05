-- SECTION 1
select * from {{ ref('avoid_root_usage') }}
UNION ALL
select * from {{ ref('mfa_enabled_for_console_access') }}
UNION ALL
select * from {{ ref('unused_creds_disabled') }}
-- TODO: ...
UNION ALL
-- SECTION 2
select * from {{ ref('avoid_root_usage') }}
-- TODO: ...