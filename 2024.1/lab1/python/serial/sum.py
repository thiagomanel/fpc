import os
import sys

def do_sum(path):
    _sum = 0
    with open(path, 'rb',buffering=0) as f:
        byte = f.read(1)
        while byte:
            _sum += int.from_bytes(byte, byteorder='big', signed=False)
            byte = f.read(1)
        return _sum

if __name__ == "__main__":
    paths = sys.argv[1:]
    for path in paths:
    #many error could be raised error. we don't care
        _sum = do_sum(path)
        print(path + " : " + str(_sum))
