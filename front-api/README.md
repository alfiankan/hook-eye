# Payment Front API



## Api Docs
- ``` POST http://localhost:8080/payment ```
- Json Body ``` {"reference_id":"SA5452022", "total":945000} ```

## Test
- ```curl --request POST --data '{"reference_id":"SA5452022", "total":945000}' http://localhost:8080/payment```