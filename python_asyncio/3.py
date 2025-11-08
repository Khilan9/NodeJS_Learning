"""
Module used to learn asyncio.TaskGroups
"""

import asyncio
from utils import fetch_data, async_timer

@async_timer
async def main():
    """
    Example for asyncio.TaskGroups
    """
    tasks = list()

    async with asyncio.TaskGroup() as tg:
        for i, delay in enumerate([3,2,1], start=1):
            task = tg.create_task(fetch_data(i, delay))
            tasks.append(task)

    results = [task.result() for task in tasks]
    print("***********\nPrinting results")
    for r in results:
        print(r)
        print("%%%%%%%%%%")

asyncio.run(main())