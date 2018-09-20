package poly

// Rect is rectangle
type Rect struct {
	Min, Max Point
}

// Ring returns an exterior ring
func (rect Rect) Ring() Ring {
	return Ring{
		rect.Min,
		{rect.Max.X, rect.Min.Y},
		rect.Max,
		{rect.Min.X, rect.Max.Y},
		rect.Min,
	}
}

// Polygon returns a polygon for the rect
func (rect Rect) Polygon() Polygon {
	return Polygon{rect.Ring(), nil}
}

// IntersectsRect detects if two bboxes intersect.
func (rect Rect) IntersectsRect(other Rect) bool {
	if rect.Min.Y > other.Max.Y || rect.Max.Y < other.Min.Y {
		return false
	}
	if rect.Min.X > other.Max.X || rect.Max.X < other.Min.X {
		return false
	}
	return true
}

// InsideRect detects rect is inside of another rect
func (rect Rect) InsideRect(other Rect) bool {
	if rect.Min.X < other.Min.X || rect.Max.X > other.Max.X {
		return false
	}
	if rect.Min.Y < other.Min.Y || rect.Max.Y > other.Max.Y {
		return false
	}
	return true
}

// InsidePolygon detects if a rect intersects another polygon
func (rect Rect) InsidePolygon(polygon Polygon) bool {
	return rect.Polygon().Exterior.InsidePolygon(polygon)
}

// IntersectsPolygon detects if a rect intersects another polygon
func (rect Rect) IntersectsPolygon(polygon Polygon) bool {
	return rect.Polygon().Exterior.IntersectsPolygon(polygon)
}