# EDcommerce API

This is the API for the e-commerce built by EDteam

## How to configure and execute

1. Build the binary file

We have a `Makefile` that allows us to execute the formatting, linter, test and build the binary file.
```shell
make
```

2. Configure the `.env` file

2.1 Copy the example `.env.example` file
```shell
cp cmd/.env.example .env
```

2.2 Set up your values for your execution

3. Execute your binary file
```shell
./ecommerce
```

## What is a webhook?

For process and validate the payments via PayPal we need to process the PayPal's webhook tool. The docs are in [Webhook Documentation](Webhook.md)

## How to configure your PayPal

Go to [PayPal Documentation](PayPal.md) documentation.

## Hexagonal Architecture explanation

Go to [Hexagonal Architecture](Hexagonal.md) documentation.
