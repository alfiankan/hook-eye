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

## Demo Running
<img width="1277" alt="Screen Shot 2022-01-28 at 00 11 16" src="https://user-images.githubusercontent.com/40946917/151408915-52bd3d4e-3072-4cf0-a80d-28022f2cf6df.png">
