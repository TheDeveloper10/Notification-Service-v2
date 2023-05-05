# Service Configuration

The file "service.example.json" contains an example configuration. Clone that file in "service.json" and put inside it the actual configuration of the service.

## Configuration Fields

| field name | description | allowed values |
|------------|-------------|----------------|
| service | Contains configuration for the entire service |
| service.apis | List of APIs the service will run (e.g. http_rest_v1 - REST API Version 1) | http_rest_v1; rabbitmq_v1
| service.notification_types | List of allowed notification types (if one is not allowed notifications of this type won't be sent) | mail; sms; push
| service.allowed_languages | List of allowed language fields | *
| service.auth | Contains configuration for the auth
| service.auth.master_client_id | Client ID of the master | any 16 character string
| service.auth.master_client_secret | Client Secret of the master | any 128 character string
| service.auth.token_expiry_time | Seconds after which the issued token will expire | any positive number
| service.auth.max_active_clients | Maximum number of issued tokens that are active at the same time | any positive number
| service.cache | Contains configuration for the internal caches
| service.cache.templates_cache_limit | Maximum number of items in the templates cache | any positive number
| service.cache.templates_cache_entry_expiry | Time (in seconds) after which entries in the templates cache will expire | any positive number
| service.cache.templates_cache_cleanup_time | Delay (in seconds) between cleanups of filled templates cache | any positive number
| http | Configuration for the http server
| http.address | Address on which the server is going to work (e.g. ":8080")
| rabbitmq | Configuration of the RabbitMQ client
| rabbitmq.username | Username of the client
| rabbitmq.password | Password of the client
| rabbitmq.host | Host of the client
| database | Configuration of the Database client
| database.driver | Used driver (by default it only works with the mysql driver) | "mysql"
| database.username | Username of the client
| database.password | Password of the client
| database.host | Host of the client
| database.name | Name of the database
| database.pool_size | Maximum number of connections in the pool | any positive integer
| mail | Configuration of the Mail client
| mail.from_email | Email from which mails will be sent
| mail.from_password | Password of the client
| mail.host | Host of the smtp server
| mail.port | Port of the smtp server
| sms | Configuration of the Mail client (Twilio configuration)
| sms.from_phone_number | Phone number which will be issuing the notifications
| sms.account_sid | Account SID
| sms.authentication_token | Authentication Token