# Hook-eye

### Just my learning playground about event driven kafka, distributed app and like this repo name Webhook


## Infrastructure
- Front Api
    - handing end user / app API

- Payment Service
    - processing payment

- Webhook
    - send notification/hook to 3rd party app

```

    end user/app                                    3rd party app/hook
        ^                                                   ^
        |                                                   |
        |                                                   |
    |===========|       |=================|         |=================|
    | Front API |       | payment service |         | Webhook Service |
    |===========|       |=================|         |=================|
        ||                      ||                          ||
        ||                      ||                          ||
        ||                      ||                          ||
    |===================================================================|
    |                           EVENT BUS                               |
    |===================================================================|
                        ||
                        ||
                        ||
            |===========================|
            |       Websocket           |
            |    Notification Service   |
            |===========================|
                        |
                        |
                        |
                        |
                    Web/Mobile App





```