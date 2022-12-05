select
--     :'execution_time' as execution_time,
--     :'framework' as framework,
--     :'check_id' as check_id,
    'Ensure MFA is enabled for all IAM users that have a console password (Scored)' as title,
    split_part(arn, ':', 5) as account_id,
    arn as resource_id,
    case when
                     password_status IN ('TRUE', 'true') and not mfa_active
             then 'fail'
         else 'pass'
        end as status
from {{ source('cloudquery', 'aws_iam_credential_reports') }}
