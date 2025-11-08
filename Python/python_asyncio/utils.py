"""
Utility functions defined for small reuseable blocks of repetative code
"""

from datetime import datetime
import asyncio

def async_timer(func):
    """
    Decorator function that prints execution time for async functions

    Args:
        func (callable): async function
    """
    async def wrapper(*args, **kwargs):
        start_time = datetime.now()
        print(f"Starting execution for {func.__name__}")
        print("=================================")
        await func(*args, **kwargs)
        end_time = datetime.now()
        delta = end_time - start_time
        print("=================================")
        print(f"It took {delta.total_seconds()} seconds to execute {func.__name__}!")
        return 0
    return wrapper

def timer(func):
    """
    Decorator function that prints execution time for functions

    Args:
        func (callable): function
    """
    def wrapper(*args, **kwargs):
        start_time = datetime.now()
        print(f"Starting execution for {func.__name__}")
        print("=================================")
        func(*args, **kwargs)
        end_time = datetime.now()
        delta = end_time - start_time
        print("=================================")
        print(f"It took {delta.total_seconds()} seconds to execute {func.__name__}!")
        return 0
    return wrapper

async def fetch_data(_id:int, delay:int) -> dict[str,str]:
    """
    Demo method used for simulating I/O operations

    Args:
        _id (int): id of the coroutine
        delay (int): delay for which we want to simulate I/O operation

    Returns:
        dict[str,str]: return some data
    """
    print(f"Starting fetching data for {_id}")
    await asyncio.sleep(delay)
    print(f"Data Fetched for {_id}")
    return {
        "data":"Some data",
        "id":_id
    }
