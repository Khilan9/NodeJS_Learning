
def filterme(services, miss):

    cache = set()
    new_details_macros = []
    for svc in services:
        for curr in miss['macros']['details']:
            if curr['name'] in cache:
                continue
            for objects in curr['used_in_itsi_objects']:
                svc_kpi_keys = objects["_key"]
                if svc in svc_kpi_keys:
                    new_details_macros.append(curr.copy())
                    cache.add(curr['name'])

    new_details_ss = []
    for svc in services:
        for curr in miss['savedsearches']['details']:
            if curr['name'] in cache:
                continue
            for objects in curr['used_in_itsi_objects']:
                svc_kpi_keys = objects["_key"]
                if svc in svc_kpi_keys:
                    new_details_ss.append(curr.copy())
                    cache.add(curr['name'])

    temp = {
        'macros':{
            'count': len(new_details_macros),
            'details':new_details_macros
        },
        'savedsearches':{
            'count':len(new_details_ss),
            'details':new_details_ss
        }
    }

    return temp