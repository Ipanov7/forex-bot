# Forex Bot

A basic Telegram bot to fetch currency rates and notify you for any potential gain/loss

## Getting Started

* You will need to setup a Telegram bot. You can start from [here](https://core.telegram.org/bots)
* I use APILayer's [Exchange Rates Data API](https://apilayer.com/marketplace/exchangerates_data-api) because it's free and simple, but anything goes as long as you are willing to adapt the code a bit.

### Executing program

This bot was made to run on AWS Lambda functions, so it's necessary to create a zip archive (sigh) to deploy the code 
1. To generate an archive that can be uploaded on AWS,
just run the `zip.sh` script.

1. Remember to setup the input data:
    ```json
    {
        "from": "USD",
        "to":"EUR",
        "avg_rate":1.00000
    }
    ```
1. Choose an appropriate schedule to run the bot, like `rate(1 hour)`
1. ???
1. Profit!

## Author

Loris Occhipinti
* ✉️Contact me at: loris@lorisocchipinti.com
* ⭐Website: https://blog.lorisocchipinti.com

