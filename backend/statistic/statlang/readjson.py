import jsonlines

with open("iterm.jl","r+", encoding="utf8") as f:
    for item in jsonlines.Reader(f):
        print(item)