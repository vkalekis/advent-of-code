from PIL import Image, ImageDraw, ImageColor
import json

def adjust_color(x, min_value, max_value):
    start_color = "yellow"
    end_color = "red"

    start_rgba = ImageColor.getcolor(start_color, "RGBA")
    end_rgba = ImageColor.getcolor(end_color, "RGBA")

    interpolated_color = [
        int(start + (end - start) * (x - min_value) / (max_value - min_value))
        for start, end in zip(start_rgba, end_rgba)
    ]

    return tuple(interpolated_color)

width, height = 200, 200
wr, hr = 9, 9


class Drawable:
    def __init__(self, draw, coords, c="blue"):
        self.draw = draw
        self.coords = coords
        self.c = c

    def draw1(self, startx, starty, c=None):
        if c is None:
            c = self.c
        for coord in self.coords:
            if len(coord) != 2:
                continue
            x,y = coord[0], coord[1]
            self.draw.rectangle(
                (
                    startx + x * wr / 3,
                    starty + y * hr / 3,
                    startx + (x + 1) * wr / 3,
                    starty + (y + 1) * hr / 3,
                ),
                fill=c,
                width=2,
            )


image = Image.new("RGB", (width, height), "white")


draw = ImageDraw.Draw(image)

# d = Drawable(draw, [0], [1])


with open("data/2023/input10_example") as f:
    lines = [line.strip() for line in f.readlines()]
    


vert = Drawable(draw,[(1,0),(1,1),(1,2)])
hor = Drawable(draw,[(0,1),(1,1),(2,1)])
ne = Drawable(draw,[(1,0),(1,1),(2,1)])
nw = Drawable(draw,[(1,0),(1,1),(0,1)])
sw = Drawable(draw,[(1,2),(1,1),(0,1)])
se = Drawable(draw, [(1,2),(1,1),(2,1)])
gr = Drawable(draw,[(1,1)], c="black")



j = None
try:
    with open("./output.json") as f:
        j = json.load(f)
except Exception as ex:
    pass

print(j)

for y,line in enumerate(lines):
    for x,ch in enumerate(line):
        startx, starty = x*wr, y*hr

        c = None
        if j is not None:
            coord = f"{y}-{x}"
            if coord in j.keys() and j[coord]!=0:
                # c = "orange"
                c = adjust_color(j[coord], 0, 6828)

        if ch == ".":
            gr.draw1(startx,starty,c)
        if ch == "|":
            vert.draw1(startx,starty,c)
        if ch == "-":
            hor.draw1(startx,starty,c)
        if ch == "L":
            ne.draw1(startx,starty,c)
        if ch == "J":
            nw.draw1(startx,starty,c)
        if ch == "7":
            sw.draw1(startx,starty,c)
        if ch == "F":
            se.draw1(startx,starty,c)
        if ch == "S":
            nw.draw1(startx,starty,"red")

# image.show()
image.save("/tmp/test.png")
