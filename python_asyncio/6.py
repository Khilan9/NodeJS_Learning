"""
Module Teaching event.set(), event.wait() semaphore
"""
import asyncio

async def waiter(event):
    print("Waiting for event to be set in waiter")
    await event.wait()
    print("Event is set, Resuming task in waiter")


async def setter(event):
    print("Waiting for some event to happen in setter")
    await asyncio.sleep(3)
    print("Event will be set to true in setter")
    event.set()
    print("Event has been set in setter")

async def main():
    event = asyncio.Event()
    waiter1 = asyncio.create_task(waiter(event))
    setter1 = asyncio.create_task(setter(event))

    await waiter1
    await setter1

asyncio.run(main())