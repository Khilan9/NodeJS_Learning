"""
Module Teaching Futures in asyncio
"""

from utils import async_timer
import asyncio

SLEEP = 3

@async_timer
async def function_setting_future_result(future1, future2, value):
    """
    Method used for demoing function used for setting future value

    Args:
        future (obj): asyncio future object
        value (int): value to be set to future
    """
    global SLEEP
    print(f"Doint I/O task for {SLEEP} seconds")
    await asyncio.sleep(SLEEP)

    # Set future result
    future1.set_result(value)
    print(f"Setting future1 result value to: {value}")

    # %%%%%%%%%%%%%%%
    print(f"Doint I/O task for {SLEEP} seconds")
    await asyncio.sleep(SLEEP)

    # Set future result
    future2.set_result(value)
    print(f"Setting future2 result value to: {value}")

async def main():
    """
    Main Function
    """
    # Creating future object
    loop = asyncio.get_running_loop()
    future1 = loop.create_future()
    future2 = loop.create_future()

    asyncio.create_task(function_setting_future_result(future1, future2, 10))

    # Wait for result
    res = await future1
    print(f"Result for future1: {res}")
    res = await future2
    print(f"Result for future2: {res}")



asyncio.run(main())
