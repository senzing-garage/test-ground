#!/usr/bin/env python3
"""Hello World script that prints a greeting with the current timestamp."""

from datetime import datetime


def main():
    """Print Hello World with current timestamp."""
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    print("Hello, world!")
    print(f"The current time is {timestamp}")


if __name__ == "__main__":
    main()
