package main

import (
	"fmt"
	"os"
	"time"
)

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func main() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
	// secondHandPt := secondHand(os.Stdout, t)
	// io.WriteString(os.Stdout, svgStart)
	// io.WriteString(os.Stdout, bezel)
	// io.WriteString(os.Stdout, secondHandTag(secondHandPt))
	// io.WriteString(os.Stdout, svgEnd)
}

func secondHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
