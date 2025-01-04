import hashlib


secretKey = "iwrupvqb"

def GetLowestValue(secretKey):
    number = 1
    while True:
        combined = f"{secretKey}{number}"
        md5_hash = hashlib.md5(combined.encode()).hexdigest()

        if md5_hash.startswith("000000"):
            return number
    
        number += 1

print(GetLowestValue(secretKey))
