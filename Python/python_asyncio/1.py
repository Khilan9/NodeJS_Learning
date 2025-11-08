import asyncio

async def simulate_io_process(delay):
    print(f"Fetching {delay}GiB Data")
    await asyncio.sleep(delay)
    print("Data Fetched Successfully")
    return {
        "amount":f"{delay} GiB"
    }

async def couroutine(number=0, delay=1):
    print(f"Starting couroutine {number}")
    task = simulate_io_process(delay)
    fetched_data = await task
    print(f"Data got in couroutine {fetched_data}")
    print(f"End of couroutine {number}")
    print()


routines = [couroutine(1, 3), couroutine(2, 2)]

for routine in routines:
    asyncio.run(routine, debug=True)