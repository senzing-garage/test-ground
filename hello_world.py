#!/usr/bin/env python3
"""Hello World script with greeting, timestamp and IBM stock price."""

from datetime import datetime

import yfinance as yf


def get_ibm_stock_price():
    """Fetch the current IBM stock price."""
    ibm = yf.Ticker("IBM")
    price = ibm.info.get("regularMarketPrice")
    return price


def main():
    """Print Hello World with current timestamp and IBM stock price."""
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    ibm_price = get_ibm_stock_price()
    print("Hello, world!")
    print(f"The current time is {timestamp}")
    print(f"The current IBM stock price is: ${ibm_price:.2f}")


if __name__ == "__main__":
    main()
