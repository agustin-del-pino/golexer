package golexer

const EOF = 0x00

type ByteRange interface {
	IsInRange(byte) bool
}

type BytePoints interface {
	HasPoint(byte) bool
}

type SingleByteRange struct {
	from byte
	to   byte
}

func (r SingleByteRange) IsInRange(b byte) bool {
	return b >= r.from && b <= r.to
}

type CompoundByteRange struct {
	ranges []ByteRange
}

func (r *CompoundByteRange) IsInRange(b byte) bool {
	for _, v := range r.ranges {
		if !v.IsInRange(b) {
			return false
		}
	}
	return true
}

type BPoints struct {
	points []byte
}

func (p *BPoints) HasPoint(b byte) bool {
	for _, v := range p.points {
		if b == v {
			return true
		}
	}
	return false
}

type defaultByteRange struct{}

func (d defaultByteRange) IsInRange(b byte) bool {
	return false
}

type defaultBytePoint struct{}

func (d defaultBytePoint) HasPoint(b byte) bool {
	return false
}

func NewSingleByteRange(f byte, t byte) ByteRange {
	return SingleByteRange{
		from: f,
		to:   t,
	}
}

func NewBPoints(p ...byte) BytePoints {
	return &BPoints{
		points: p,
	}
}
