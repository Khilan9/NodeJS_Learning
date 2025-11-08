"""
Module teaching how to create task and run them simultaneously
"""
import asyncio
from utils import async_timer, fetch_data

@async_timer
async def task_tutorial():
    """
    task_tutorial method
    """
    task1 = asyncio.create_task(fetch_data(1, 3))
    task3 = asyncio.create_task(fetch_data(3, 1))

    result1 = await task1
    print("Print 1")
    task2 = asyncio.create_task(fetch_data(2, 2))
    result2 = await task2
    print("Print 2")
    result3 = await task3
    print("Print 3")

    print(result1)
    print(result2)
    print(result3)

@async_timer
async def task_tutorial_temp():
    """
    task_tutorial method
    """
    result = await fetch_data(2, 1)

    print(result)

@async_timer
async def task_tutorial_2():
    """
    task_tutorial method
    """
    task1 = asyncio.create_task(fetch_data(2, 1))
    task2 = asyncio.create_task(fetch_data(1, 2))

    result1 = await task1
    result2 = await task2

    task3 = asyncio.create_task(fetch_data(3, 3))
    result3 = await task3

    print(result1)
    print(result2)
    print(result3)

@async_timer
async def normal_await_tutorial():
    """
    normal_await_tutorial method
    """
    task1 = fetch_data(2, 1)
    task2 = fetch_data(1, 2)
    task3 = fetch_data(3, 3)

    result1 = await task1
    result2 = await task2
    result3 = await task3

    print(result1)
    print(result2)
    print(result3)

@async_timer
async def gather_asyncs():
    """
    Tutorial for asycio.gather()
    """
    results = await asyncio.gather(fetch_data(1, 1), fetch_data(2, 2), fetch_data(3, 3))

    for r in results:
        print(r)

# asyncio.run(task_tutorial_temp())

# asyncio.run(task_tutorial())

# asyncio.run(task_tutorial_2())

# asyncio.run(normal_await_tutorial())

asyncio.run(gather_asyncs())