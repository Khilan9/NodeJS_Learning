"""
Module Teaching synchronisation in asyncio
"""
import asyncio

LOCK = asyncio.Lock()
GLOBAL_DB = {
    "users":["shivam","antra", "khilan"],
    "ids":{6:"shivam", 9:"antra", 4:"khilan"}
}

async def update_user(user):
    global GLOBAL_DB, LOCK
    async with LOCK:
        try:
            print("Updating user", user)
            GLOBAL_DB["users"][GLOBAL_DB["users"].index(user["name"])] = user["uname"]
            GLOBAL_DB["ids"][user["iid"]] = user["uname"]
            await asyncio.sleep(5)
            return "UPDATE: User updated"
        except Exception as e:
            print(f"{type(e).__name__}: {e}")
            return "UPDATE: Unable to update!!"

async def read_user(iid):
    global GLOBAL_DB, LOCK
    async with LOCK:
        try:
            print("Reading user", iid)
            await asyncio.sleep(1)
            return "READ: "+GLOBAL_DB["ids"][iid]
        except Exception as e:
            print(f"{type(e).__name__}: {e}")
            return "READ: No user found!!"

async def delete_user(iid):
    global GLOBAL_DB, LOCK
    async with LOCK:
        try:
            print("Deleting user: ", iid)
            name = GLOBAL_DB["ids"][iid]
            await asyncio.sleep(2)
            GLOBAL_DB["users"].remove(name)
            del GLOBAL_DB["ids"][iid]
            return "DELETE: User deleted"
        except Exception as e:
            print(f"{type(e).__name__}: {e}")
            print("User not found!!")
            return "DELETE: No user Found!!"


async def main():
    tasks = []
    async with asyncio.TaskGroup() as tg:
        task = tg.create_task(update_user({"name":"khilan", "uname":"khilan1", "iid":4}))
        tasks.append(task)
        task = tg.create_task(delete_user(4))
        tasks.append(task)
        task = tg.create_task(read_user(6))
        tasks.append(task)

    res = [task.result() for task in tasks]
    print("------------------------")
    for r in res:
        print(r)

    print(GLOBAL_DB)

asyncio.run(main())