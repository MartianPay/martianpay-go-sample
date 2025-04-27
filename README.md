# go-sample

This is a sample Go application that demonstrates how to use the MartianPay API.


## Test

测试前修改const apiKey = "your_api_key_here"
还有每个用例中的email和ID

```sh
go test -count=1 -v -run TestListPaymentIntents sdk/*.go
go test -count=1 -v -run TestGetPaymentIntent sdk/*.go
go test -count=1 -v -run TestListCustomers sdk/*.go
go test -count=1 -v -run TestGetCustomer sdk/*.go
```

## Webhook event receiver test

首先运行main.go，
```sh
go run main.go
```
然后用下面的curl发送event测试
```sh
curl --location 'http://localhost:8080/v1/webhook_test' \
--header 'Content-Type: application/json' \
--header 'Martian-Pay-Signature: t=1745580845,v1=c16a830eae640a659025a5f4bf91866ddd4be0c85619a82defad3c10af42ec89' \
--header 'User-Agent: MartianPay/1.0' \
--data '{"id":"evt_056qw9fr9PndljykFUqSIf6t","object":"event","api_version":"2025-01-22","created":1745580845,"data":{"previous_attributes":null,"object":{"id":"pi_nKSxQrU2Pjh9KGzyIRJcWpGJ","object":"payment_intent","amount":{"asset_id":"USD","amount":"1"},"payment_details":{"amount_captured":{"asset_id":"USD","amount":"1"},"amount_refunded":{"asset_id":"USD","amount":"0"},"tx_fee":{"asset_id":"USD","amount":"0"},"tax_fee":{"asset_id":"USD","amount":"0"},"frozen_amount":{"asset_id":"USD","amount":"0"},"net_amount":{"asset_id":"USD","amount":"0"},"gas_fee":{},"network_fee":{"asset_id":"USD","amount":"0"}},"canceled_at":0,"cancellation_reason":"","client_secret":"pi_nKSxQrU2Pjh9KGzyIRJcWpGJ_secret_S77UnKrHOv10a6x1SSjbKwwI6","created":1745580777,"updated":1745580781,"currency":"USD","customer":{"id":"cus_7Xpxk2n22WAEBDaBVlG5dnRf","object":"customer","total_expense":300,"total_payment":320,"total_refund":20,"currency":"USD","created":1740731398,"name":"","email":"yihuazhai@163.com","description":"","phone":""},"description":"test","livemode":false,"metadata":{"merchant_id":"accu_M7PTgveSgMtTtPHbjFgEtAlD","merchant_name":"ABC Company 1","payment_link":{"active":true,"created_at":1744181243,"id":"test_NivLsKfDfKuNPdULFVPvrhVx","product_items":[{"product":{"active":true,"created_at":1744181220,"description":"","id":"prod_7pwJj8FLyBsehx","metadata":null,"name":"test","picture_url":"","price":{"amount":"1","asset_id":"USD"},"tax_code":"","updated_at":1744181220},"quantity":1}],"total_price":{"amount":"1","asset_id":"USD"},"updated_at":1744181243}},"payment_link_details":{"merchant_id":"accu_M7PTgveSgMtTtPHbjFgEtAlD","merchant_name":"ABC Company 1","payment_link":{"id":"test_NivLsKfDfKuNPdULFVPvrhVx","product_items":[{"product":{"id":"prod_7pwJj8FLyBsehx","name":"test","price":{"asset_id":"USD","amount":"1"},"description":"","picture_url":"","tax_code":"","metadata":null,"active":true,"updated_at":1744181220,"created_at":1744181220},"quantity":1}],"total_price":{"asset_id":"USD","amount":"1"},"active":true,"updated_at":1744181243,"created_at":1744181243}},"merchant_order_id":"test_NivLsKfDfKuNPdULFVPvrhVx","payment_method_type":"crypto","charges":[{"id":"ch_UQBFpYEomiYmkAF0wfRIQQ7g","object":"charge","amount":{"asset_id":"USDC-Ethereum-TEST","amount":"1"},"payment_details":{"amount_captured":{"asset_id":"USDC-Ethereum-TEST","amount":"1"},"amount_refunded":{"asset_id":"USDC-Ethereum-TEST","amount":"0"},"tx_fee":{"asset_id":"USDC-Ethereum-TEST","amount":"0"},"tax_fee":{"asset_id":"USDC-Ethereum-TEST","amount":"0"},"frozen_amount":{"asset_id":"USDC-Ethereum-TEST","amount":"0"},"net_amount":{"asset_id":"USDC-Ethereum-TEST","amount":"0"},"gas_fee":{},"network_fee":{"asset_id":"USDC-Ethereum-TEST","amount":"0"}},"exchange_rate":"1.0","calculated_statement_descriptor":"","captured":false,"created":0,"customer":"","description":"","disputed":false,"failure_code":"","failure_message":"","fraud_details":null,"livemode":false,"metadata":null,"paid":false,"payment_intent":"pi_nKSxQrU2Pjh9KGzyIRJcWpGJ","payment_method_type":"crypto","payment_method_options":{"crypto":{"amount":"1000000","token":"USDC","asset_id":"USDC-Ethereum-TEST","network":"Ethereum Sepolia","decimals":6,"exchange_rate":"1.0","deposit_address":"0xa4547D6644a46ed4F395D36d4680800eF5c53bf6","expired_at":1745582581}},"transactions":[{"tx_id":"de3ce20e-1407-47e5-a406-ded38947c486","source_address":"0x36279Ac046498bF0cb742622cCe22F3cE3c2AfD9","destination_address":"0xa4547D6644a46ed4F395D36d4680800eF5c53bf6","tx_hash":"0xa2b3ed89adab6fa946c3be99d65a43f718e12a87eb92f548b3d1ded241fb8e12","amount":"1000000","decimals":6,"asset_id":"USDC-Ethereum-TEST","token":"USDC","network":"Ethereum Sepolia","type":"deposit","created_at":1745580844,"status":"confirmed","aml_status":"approved","aml_info":"","charge_id":"ch_UQBFpYEomiYmkAF0wfRIQQ7g","refund_id":"","fee_info":"network_fee:\"0.001791756437565262\"","fee_currency":"ETH_TEST5"}],"receipt_email":"","receipt_url":"","refunded":false,"refunds":[],"review":null}],"receipt_email":"yihuazhai@163.com","status":"Success","payment_intent_status":"Confirmed","one_time_payment":false}},"livemode":false,"pending_webhooks":0,"type":"payment_intent.succeeded"}'
```